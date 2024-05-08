package outbound

import (
	"context"
	"net"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/dialer"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	shadowsocks "github.com/sagernet/sing-shadowsocks2"
	"github.com/sagernet/sing/common"
	"github.com/sagernet/sing/common/bufio"
	E "github.com/sagernet/sing/common/exceptions"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"github.com/sagernet/sing/common/uot"
)

var _ adapter.Outbound = (*Shadowsocks)(nil)

type Shadowsocks struct {
	myOutboundAdapter
	dialer     N.Dialer
	method     shadowsocks.Method
	serverAddr M.Socksaddr
	uotClient  *uot.Client
}

func NewShadowsocks(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.ShadowsocksOutboundOptions) (*Shadowsocks, error) {
	method, err := shadowsocks.CreateMethod(ctx, options.Method, shadowsocks.MethodOptions{
		Password: options.Password,
	})
	if err != nil {
		return nil, err
	}
	outboundDialer, err := dialer.New(router, options.DialerOptions)
	if err != nil {
		return nil, err
	}
	outbound := &Shadowsocks{
		myOutboundAdapter: myOutboundAdapter{
			protocol:     C.TypeShadowsocks,
			network:      options.Network.Build(),
			router:       router,
			logger:       logger,
			tag:          tag,
			dependencies: withDialerDependency(options.DialerOptions),
		},
		dialer:     outboundDialer,
		method:     method,
		serverAddr: options.ServerOptions.Build(),
	}
	uotOptions := common.PtrValueOrDefault(options.UDPOverTCP)
	if uotOptions.Enabled {
		outbound.uotClient = &uot.Client{
			Dialer:  (*shadowsocksDialer)(outbound),
			Version: uotOptions.Version,
		}
	}
	return outbound, nil
}

func (h *Shadowsocks) DialContext(ctx context.Context, network string, destination M.Socksaddr) (net.Conn, error) {
	ctx, metadata := adapter.AppendContext(ctx)
	metadata.Outbound = h.tag
	metadata.Destination = destination
	switch N.NetworkName(network) {
	case N.NetworkTCP:
		h.logger.InfoContext(ctx, "outbound connection to TCP = ", destination)
	case N.NetworkUDP:
		if h.uotClient != nil {
			h.logger.InfoContext(ctx, "outbound UoT connect packet connection to ", destination)
			return h.uotClient.DialContext(ctx, network, destination)
		} else {
			h.logger.InfoContext(ctx, "outbound packet connection to UDP = ", destination)
		}
	}
	return (*shadowsocksDialer)(h).DialContext(ctx, network, destination)
}

func (h *Shadowsocks) ListenPacket(ctx context.Context, destination M.Socksaddr) (net.PacketConn, error) {
	ctx, metadata := adapter.AppendContext(ctx)
	metadata.Outbound = h.tag
	metadata.Destination = destination
	if h.uotClient != nil {
		h.logger.InfoContext(ctx, "outbound UoT packet connection to ", destination)
		return h.uotClient.ListenPacket(ctx, destination)
	} else {
		h.logger.InfoContext(ctx, "outbound packet connection to Listen1 = ", destination)
	}
	h.logger.InfoContext(ctx, "outbound packet connection to  Listen2 =", destination)
	return (*shadowsocksDialer)(h).ListenPacket(ctx, destination)
}

func (h *Shadowsocks) NewConnection(ctx context.Context, conn net.Conn, metadata adapter.InboundContext) error {
	return NewConnection(ctx, h, conn, metadata)
}

func (h *Shadowsocks) NewPacketConnection(ctx context.Context, conn N.PacketConn, metadata adapter.InboundContext) error {
	return NewPacketConnection(ctx, h, conn, metadata)
}

func (h *Shadowsocks) InterfaceUpdated() {

}

func (h *Shadowsocks) Close() error {
	return nil
}

var _ N.Dialer = (*shadowsocksDialer)(nil)

type shadowsocksDialer Shadowsocks

func (h *shadowsocksDialer) DialContext(ctx context.Context, network string, destination M.Socksaddr) (net.Conn, error) {
	ctx, metadata := adapter.AppendContext(ctx)
	metadata.Outbound = h.tag
	metadata.Destination = destination
	switch N.NetworkName(network) {
	case N.NetworkTCP:
		var outConn net.Conn
		var err error
		outConn, err = h.dialer.DialContext(ctx, N.NetworkTCP, h.serverAddr)
		if err != nil {
			return nil, err
		}
		return h.method.DialEarlyConn(outConn, destination), nil
	case N.NetworkUDP:
		outConn, err := h.dialer.DialContext(ctx, N.NetworkUDP, h.serverAddr)
		if err != nil {
			return nil, err
		}
		return bufio.NewBindPacketConn(h.method.DialPacketConn(outConn), destination), nil
	default:
		return nil, E.Extend(N.ErrUnknownNetwork, network)
	}
}

func (h *shadowsocksDialer) ListenPacket(ctx context.Context, destination M.Socksaddr) (net.PacketConn, error) {
	ctx, metadata := adapter.AppendContext(ctx)
	metadata.Outbound = h.tag
	metadata.Destination = destination
	outConn, err := h.dialer.DialContext(ctx, N.NetworkUDP, h.serverAddr)
	if err != nil {
		return nil, err
	}
	return h.method.DialPacketConn(outConn), nil
}

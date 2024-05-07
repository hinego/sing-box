package inbound

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/experimental/libbox/platform"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func New(ctx context.Context, router adapter.Router, logger log.ContextLogger, options option.Inbound, platformInterface platform.Interface) (adapter.Inbound, error) {
	if options.Type == "" {
		return nil, E.New("missing inbound type")
	}
	switch options.Type {
	case C.TypeTun:
		return NewTun(ctx, router, logger, options.Tag, options.TunOptions, platformInterface)
	case C.TypeSOCKS:
		return NewSocks(ctx, router, logger, options.Tag, options.SocksOptions), nil
	case C.TypeHTTP:
		return NewHTTP(ctx, router, logger, options.Tag, options.HTTPOptions)
	case C.TypeMixed:
		return NewMixed(ctx, router, logger, options.Tag, options.MixedOptions), nil
	default:
		return nil, E.New("unknown inbound type: ", options.Type)
	}
}

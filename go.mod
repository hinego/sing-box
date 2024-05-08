module github.com/sagernet/sing-box

go 1.20

replace github.com/sagernet/sing-shadowsocks2 v0.2.0 => /opt/temp/singbox_dependon/sing-shadowsocks2

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/insomniacslk/dhcp v0.0.0-20231206064809-8c70d406f6d2
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/miekg/dns v1.1.59
	github.com/oschwald/maxminddb-golang v1.12.0
	github.com/sagernet/bbolt v0.0.0-20231014093535-ea5cb2fe9f0a
	github.com/sagernet/gomobile v0.1.3
	github.com/sagernet/sing v0.4.0-beta.18
	github.com/sagernet/sing-dns v0.2.0-beta.16
	github.com/sagernet/sing-shadowsocks2 v0.2.0
	github.com/sagernet/sing-tun v0.2.7-beta.1
	github.com/sagernet/tfo-go v0.0.0-20231209031829-7b5343ac1dc6
	github.com/sagernet/wireguard-go v0.0.0-20231215174105-89dec3b2f3e8
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.9.0
	go4.org/netipx v0.0.0-20231129151722-fdeea329fbba
	golang.org/x/crypto v0.22.0
	golang.org/x/sys v0.19.0
	howett.net/plist v1.0.1
)

//replace github.com/sagernet/sing => ../sing

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sagernet/gvisor v0.0.0-20240315080113-799fb6b6d311 // indirect
	github.com/sagernet/netlink v0.0.0-20220905062125-8043b4a9aa97 // indirect
	github.com/sagernet/quic-go v0.42.0-beta.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/u-root/uio v0.0.0-20230220225925-ffce2a382923 // indirect
	github.com/vishvananda/netns v0.0.0-20211101163701-50045581ed74 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.2.1 // indirect
)

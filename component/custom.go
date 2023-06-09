//go:build custom || full
// +build custom full

package build

import (
	_ "github.com/nzai/trojan-go/proxy/custom"
	_ "github.com/nzai/trojan-go/tunnel/adapter"
	_ "github.com/nzai/trojan-go/tunnel/dokodemo"
	_ "github.com/nzai/trojan-go/tunnel/freedom"
	_ "github.com/nzai/trojan-go/tunnel/http"
	_ "github.com/nzai/trojan-go/tunnel/mux"
	_ "github.com/nzai/trojan-go/tunnel/router"
	_ "github.com/nzai/trojan-go/tunnel/shadowsocks"
	_ "github.com/nzai/trojan-go/tunnel/simplesocks"
	_ "github.com/nzai/trojan-go/tunnel/socks"
	_ "github.com/nzai/trojan-go/tunnel/tls"
	_ "github.com/nzai/trojan-go/tunnel/tproxy"
	_ "github.com/nzai/trojan-go/tunnel/transport"
	_ "github.com/nzai/trojan-go/tunnel/trojan"
	_ "github.com/nzai/trojan-go/tunnel/websocket"
)

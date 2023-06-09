package server

import (
	"github.com/nzai/trojan-go/config"
	"github.com/nzai/trojan-go/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return new(client.Config)
	})
}

package main

import (
	"flag"
	"os"

	"github.com/kardianos/service"
	_ "github.com/nzai/trojan-go/component"
	"github.com/nzai/trojan-go/log"
	"github.com/nzai/trojan-go/proxy"
	_ "github.com/nzai/trojan-go/proxy/client"
)

type program struct {
	Proxy *proxy.Proxy
}

func (p *program) Start(s service.Service) error {
	go p.Proxy.Run()

	return nil
}

func (p *program) Stop(s service.Service) error {
	return p.Proxy.Close()
}

var (
	configPath = flag.String("c", "", "Trojan-Go config filename (.yaml/.yml/.json)")
)

func main() {
	flag.Parse()

	buffer, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("failed to read config.json due to %v", err)
	}

	proxy, err := proxy.NewProxyFromConfigData(buffer, true)
	if err != nil {
		log.Fatal(err)
	}

	prg := &program{Proxy: proxy}

	svcConfig := &service.Config{
		Name:        "tg",
		DisplayName: "Trojan-go Service",
		Description: "proxy service",
	}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

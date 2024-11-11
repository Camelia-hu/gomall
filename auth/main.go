package main

import (
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Println("auth service register err : ", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")

	svc := authservice.NewServer(new(AuthServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(
			&registry.Info{
				ServiceName: "auth",
				Weight:      1,
			},
		),
		server.WithServiceAddr(addr),
	)

	err = svc.Run()
	if err != nil {
		log.Println("auth service run err : ", err)
	}
}

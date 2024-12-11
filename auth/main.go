package main

import (
	"context"
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth/authservice"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/trace"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	conf.ViperInit()
	dao.MysqlInit()
	dao.RedisInit()
	p := trace.TraceInit("auth")
	defer p.Shutdown(context.Background())
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
		server.WithSuite(tracing.NewServerSuite()),
	)

	err = svc.Run()
	if err != nil {
		log.Println("auth service run err : ", err)
	}
}

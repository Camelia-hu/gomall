package main

import (
	"context"
	"github.com/Camelia-hu/gomall/cart/kitex_gen/cart/cartservice"
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
	p := trace.TraceInit("cart")
	defer p.Shutdown(context.Background())
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatalln("cart register err : ", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8084")
	if err != nil {
		log.Fatalln("resolve addr err : ", err)
	}
	svc := cartservice.NewServer(new(CartServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(
			&registry.Info{
				ServiceName: "cart",
				Weight:      1,
			},
		),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)
	err = svc.Run()
	if err != nil {
		log.Fatalln("err。", err)
	}
}

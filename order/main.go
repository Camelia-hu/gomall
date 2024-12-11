package main

import (
	"context"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/order/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	conf.ViperInit()
	dao.MysqlInit()
	dao.RedisInit()
	p := provider.NewOpenTelemetryProvider(
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
		provider.WithServiceName("order"),
	)
	defer p.Shutdown(context.Background())
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Println("order service register err :", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8085")

	svc := orderservice.NewServer(new(OrderServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(
			&registry.Info{
				ServiceName: "order",
				Weight:      1,
			},
		),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)

	err = svc.Run()
	if err != nil {
		log.Println("order service run err : ", err)
	}
}

package main

import (
	"context"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/Camelia-hu/gomall/product/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	conf.ViperInit()
	dao.MysqlInit()
	dao.RedisInit()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("product"),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Println("product service register err :", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8083")

	svc := productcatalogservice.NewServer(new(ProductCatalogServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(
			&registry.Info{
				ServiceName: "product",
				Weight:      1,
			},
		),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)

	err = svc.Run()
	if err != nil {
		log.Println("product service run err : ", err)
	}
}

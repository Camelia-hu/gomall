package main

import (
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/order/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

func main() {
	conf.ViperInit()
	dao.MysqlInit()
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
	)

	err = svc.Run()
	if err != nil {
		log.Println("order service run err : ", err)
	}
}

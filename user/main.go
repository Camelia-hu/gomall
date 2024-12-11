package main

import (
	"context"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/dao"
	"github.com/Camelia-hu/gomall/trace"
	"github.com/Camelia-hu/gomall/user/kitex_gen/user/userservice"
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
	p := trace.TraceInit("user")
	defer p.Shutdown(context.Background())
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal("user service register err : ", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")

	svc := userservice.NewServer(new(UserServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "user",
			Weight:      1,
		}),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
	)
	err = svc.Run()
	if err != nil {
		log.Println("user service run err : ", err)
	}
}

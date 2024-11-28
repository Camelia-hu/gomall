package main

import (
	checkout "github.com/Camelia-hu/gomall/checkout/kitex_gen/checkout/checkoutservice"
	"log"
)

func main() {
	svr := checkout.NewServer(new(CheckoutServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
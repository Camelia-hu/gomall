package dao

import (
	"github.com/Camelia-hu/gomall/cart/kitex_gen/cart"
	"github.com/Camelia-hu/gomall/conf"
	"github.com/Camelia-hu/gomall/module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func MysqlInit() {
	dsn := conf.Conf.GetString("data.mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database init err : ", err)
	}
	DB = db

	err = DB.AutoMigrate(&module.User{})
	if err != nil {
		log.Println(err)
	}
	err = DB.AutoMigrate(&module.Product{})
	if err != nil {
		log.Println(err)
	}
	err = DB.AutoMigrate(&cart.CartItem{})
}

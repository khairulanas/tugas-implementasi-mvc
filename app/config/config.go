package config

import (
	"fmt"
	"tugasmvc/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("root:@/digitalent_bank?charset=utf8&parseTime=True&loc=Local")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(new(model.Account), new(model.Transaction))
	return db
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("root:pintar123@/digitalent_bank?charset=utf8&parseTime=True&loc=Local")), &gorm.Config{})
	if err != nil {
		panic("failede to connect to database" + err.Error())
	}
	DB.AutoMigrate(new(model.Account), new(model.Transaction))

}

package db

import (
	"fmt"
	"log"
	"ops_tool/config"
	"ops_tool/module"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQL *gorm.DB

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.MySQL.Username, config.Config.MySQL.Password, config.Config.MySQL.Host, config.Config.MySQL.Port, config.Config.MySQL.DB)
	log.Println(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		MySQL = db
		db.AutoMigrate(&module.DNSServer{})
	}

}

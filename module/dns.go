package module

import "gorm.io/gorm"

type DNSServer struct {
	gorm.Model
	ID       int    `gorm:"column:id;primaryKey unique"`
	ServerIP string `gorm:"column:server_ip;type:varchar(25)"`
}

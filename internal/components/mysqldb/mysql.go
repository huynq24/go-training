package mysqldb

import (
	"golang-training/internal/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config common.Config) *gorm.DB {
	user := config.Mysql.User
	password := config.Mysql.Password
	ip := config.Mysql.Ip
	port := config.Mysql.Port
	dbName := config.Mysql.DbName

	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

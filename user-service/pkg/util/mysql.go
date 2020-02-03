package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDBConnection(cfg *Config) (*gorm.DB, error) {
	// Get database details from environment variables
	/*
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		DBName := os.Getenv("DB_NAME")
		password := os.Getenv("DB_PASSWORD")
	*/

	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.MysqlUser, cfg.Mysql.MysqlPwd, cfg.Mysql.MysqlIp, cfg.Mysql.MysqlPort, cfg.Mysql.DbName))
}

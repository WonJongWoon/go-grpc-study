package todo

import (
	"fmt"
	"go-grpc-study/internal/config"
	"go-grpc-study/internal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _logger = logger.NewSugar("mysql")

func generateDSN(user string, password string, databaseName string) string {
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, databaseName)
}

func NewDB(appConfig *config.AppConfig) *gorm.DB {
	if appConfig.DB.User == "" {
		_logger.Panicf("'database.user' in profile can not be empty")
	}

	if appConfig.DB.DBName == "" {
		_logger.Panicf("'database.database-name' in profile can not be empty")
	}

	db, err := gorm.Open(
		mysql.Open(generateDSN(
			appConfig.DB.User,
			appConfig.DB.Password,
			appConfig.DB.DBName,
		)),
		&gorm.Config{},
	)
	if err != nil {
		_logger.Panicf("failed to open mysql connection")
	}
	return db
}

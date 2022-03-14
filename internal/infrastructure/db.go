package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

const (
	EnvDBHostKey     string = "PATIKA_DB_HOST"
	EnvDBPortKey     string = "PATIKA_DB_PORT"
	EnvDBUserKey     string = "PATIKA_DB_USER"
	EnvDBPasswordKey string = "PATIKA_DB_PASSWORD"
	EnvDBNameKey     string = "PATIKA_DB_NAME"
)

func NewPostgresDB() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv(EnvDBHostKey),
		os.Getenv(EnvDBUserKey),
		os.Getenv(EnvDBPasswordKey),
		os.Getenv(EnvDBNameKey),
		os.Getenv(EnvDBPortKey))
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "bakyazi_",
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}
	return db
}

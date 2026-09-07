package db

import (
	"fmt"
	"log"
	"time"

	"github.com/soheilsleep/golang-clean-web-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	dbClient, err := gorm.Open(postgres.Open(cnn))
	if err != nil {
		return err
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)
	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	log.Println("Connected to PostgreSQL successfully")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}
func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}

package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("envファイルの読み込みに失敗しました。 \n", err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("PSQL_CONTAINER_HOST"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_SSLMODE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	return db
}

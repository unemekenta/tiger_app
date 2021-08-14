package models

import (
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("envファイルの読み込みに失敗しました。 \n", err)
	}
	dsn := fmt.Sprintf("host=localhost port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("PSQL_PORT"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PASS"))

	log.Print(dsn)
	// log.Print("PostgreSQL DBに接続しています...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	return db
}

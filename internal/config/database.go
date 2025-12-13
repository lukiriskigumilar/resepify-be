package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	DB = db
	log.Println("database connected")
}

func TestConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("failed to get sql Db form GROm", err)
		return
	}

	//test ping
	err = sqlDB.Ping()
	if err != nil {
		log.Println("database pin failed")
	} else {
		log.Println("database ping success")
	}

	var result int
	row := sqlDB.QueryRow("SELECT 1")
	err = row.Scan(&result)

	if err != nil {
		log.Println("❌ Test query failed:", err)
	} else {
		log.Println("🔥 Test query success! Result:", result)
	}
}

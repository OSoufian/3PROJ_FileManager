package main

import (
	"fmt"
	"log"
	"os"
	"video/internal/domain"
	"video/internal/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	/* env vars */

	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load(".env")
	}

	postgresHost := os.Getenv("PostgresHost")
	postgresUser := os.Getenv("PostgresUser")
	postgresPassword := os.Getenv("PostgresPassword")
	postgresDatabase := os.Getenv("PostgresDatabase")
	postgresPort := os.Getenv("PostgresPort")

	appListen := os.Getenv("AppListen")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort)
	domain.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	// Assuming you're using the "database/sql" package

	if err != nil {
        log.Fatal(err)
    }
    // Assuming you're using the "database/sql" package
    tx := domain.Db.Begin()
    tx.AutoMigrate(&domain.Videos{})

    if tx.Error != nil {
        // Handle error
        tx.Rollback()
    }

    if tx.Commit().Error != nil {
        // Handle error
        tx.Rollback()
	}

	// Start Fiber app
	http.Http().Listen(appListen)
}

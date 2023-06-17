package main

import (
	"fmt"
	// "log"
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
	domain.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
		PrepareStmt:       true,
	})

	// Assuming you're using the "database/sql" package
    domain.Db.AutoMigrate(&domain.Videos{})


	// Start Fiber app
	http.Http().Listen(appListen)
}

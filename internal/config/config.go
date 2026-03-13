package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("root@tcp(localhost:3306)/nutrimama?parseTime=true")

    fmt.Printf("Connecting to database '%s' at %s:%s as user '%s'\n",
        os.Getenv("DB_NAME"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
    )
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
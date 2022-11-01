package bootstrap

import (
	"fmt"
	"os"

	"github.com/feacx/study/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PATH"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database.Connect(mysql.Open(dsn))
}

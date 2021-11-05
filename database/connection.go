package database

import (
	"fmt"
	"os"

	"github.com/fuadsuleyman/go-auth/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
// auth_db_1
func Connect() {
	x := gorm.Open

	// connection, err := x(postgres.Open("host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Baku"), &gorm.Config{})
	connection, err := x(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
	viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), viper.GetString("db.dbname"), os.Getenv("DB_PASSWORD"), viper.GetString("db.sslmode"), viper.GetString("db.TimeZone"))), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}

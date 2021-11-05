package database

import (
	"github.com/fuadsuleyman/go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// "postgres:fuadfuad@/authdb"
// const DNS := "host=localhost user=postgres password=fuadfuad dbname=authdb port=5437 sslmode=disable TimeZone=Asia/Baku"
// docker run --name=auth2-db -e POSTGRES_PASSWORD='fuadfuad' -p 5437:5432 -d --rm postgres

var DB *gorm.DB
// auth_db_1
func Connect() {
	x := gorm.Open
	connection, err := x(postgres.Open("host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Baku"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}

package models

type User struct {
	Id uint
	Username string `gorm:"unique"`
	Usertype string
	Password []byte
}
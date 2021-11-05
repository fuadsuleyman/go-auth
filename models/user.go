package models

type User struct {
	Id uint
	Username string `gorm:"unique" validate:"required" json:"username"`
	Usertype string `validate:"required" json:"usertype"`
	Password []byte `validate:"required" json:"password"`
}
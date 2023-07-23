package model

type User struct {
	Id       int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username string `gorm:"type:varchar(40);not null;unique"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"type:varchar(512);not null"`
}

type Users []User

package models

type User struct {
	ID       uint   `gorm:"primaryKey;column:id"`
	Name     string `gorm:"type:varchar(255);not null;column:name"`
	Email    string `gorm:"type:varchar(255);unique;not null;column:email"`
	Password string `gorm:"type:varchar(255);not null;column:password"`
}

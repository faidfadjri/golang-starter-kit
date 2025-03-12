package models

type User struct {
	ID              uint   `gorm:"primaryKey;column:id"`
	Fullname        string `gorm:"type:varchar(255);not null;column:fullname"`
	Username        string `gorm:"type:varchar(255);not null;column:username"`
	Email           string `gorm:"type:varchar(255);unique;not null;column:email"`
	EmailVerifiedAt string `gorm:"type:timestamp;unique;null;column:email_verified_at"`
	Password        string `gorm:"type:varchar(255);not null;column:password"`
	Phone           string `gorm:"type:varchar(255);null;column:phone"`
	Address         string `gorm:"type:varchar(255);null;column:address"`
	Avatar          string `gorm:"type:varchar(255);null;column:avatar"`
	RoleId          string `gorm:"type:bigint(20);null;column:role_id"`
	VerifiedAt      string `gorm:"type:timestamp;null;column:verified_at"`
	RememberToken   string `gorm:"type:varchar(100);null;column:remember_token"`
	CreatedAt       string `gorm:"type:timestamp;column:created_at"`
	UpdatedAt       string `gorm:"type:timestamp;column:updated_at"`
}

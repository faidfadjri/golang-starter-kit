package users

type User struct {
	ID              uint64   `gorm:"primaryKey;column:id"`
	Fullname        string   `gorm:"type:varchar(255);not null;column:fullname"`
	Username        string   `gorm:"type:varchar(255);not null;column:username"`
	Email           string   `gorm:"type:varchar(255);unique;not null;column:email"`
	EmailVerifiedAt *string  `gorm:"type:timestamp;null;column:email_verified_at"`
	Password        string   `gorm:"type:varchar(255);not null;column:password"`
	Phone           *string  `gorm:"type:varchar(255);null;column:phone"`
	Address         *string  `gorm:"type:varchar(255);null;column:address"`
	Avatar          *string  `gorm:"type:varchar(255);null;column:avatar"`
	RoleID          uint64   `gorm:"column:role_id"`
	Role            UserRole `gorm:"foreignKey:RoleID;references:ID"`
	VerifiedAt      *string  `gorm:"type:timestamp;null;column:verified_at"`
	RememberToken   *string  `gorm:"type:varchar(100);null;column:remember_token"`
	CreatedAt       *string  `gorm:"type:timestamp;column:created_at"`
	UpdatedAt       *string  `gorm:"type:timestamp;column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

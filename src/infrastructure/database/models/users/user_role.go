package users

type UserRole struct {
	ID        uint64  `gorm:"primaryKey;column:id"`
	Name      string  `gorm:"type:varchar(255);not null;column:name"`
	CreatedAt *string `gorm:"type:timestamp;null;column:created_at"`
	UpdatedAt *string `gorm:"type:timestamp;null;column:updated_at"`
}

func (UserRole) TableName() string {
	return "master_role"
}

package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserRegisterPayload struct {
	Fullname string  `json:"fullname" validate:"required"`
	Username *string `json:"username,omitempty"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
	Phone    string  `json:"phone" validate:"required"`
	Address  string  `json:"address"`
	Avatar   *string `json:"avatar,omitempty"`
	RoleId   *int    `json:"role_id,omitempty"`
}

type UserCredentials struct {
	EmailOrUsername string `json:"email_or_username" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

func (UserRegisterPayload) TableName() string {
	return "users"
}
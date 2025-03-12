package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserRegisterPayload struct {
	Fullname string  `json:"fullname"`
	Username *string `json:"username,omitempty"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
	Avatar   string  `json:"avatar"`
	RoleId   *int    `json:"role_id,omitempty"`
}

type UserLoginPayload struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}

func (UserRegisterPayload) TableName() string {
	return "users"
}
package users

type UserResponse struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}

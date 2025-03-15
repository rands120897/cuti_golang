package web

type UserRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password"`
	User_role string `json:"user_role"`
}

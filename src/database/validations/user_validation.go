package validations

type UserCreatePayload struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdatePayload struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type UserUpdatePasswordPayload struct {
	OldPassword     string `json:"oldPassword"`
	ConfirmPassword string `json:"confirmPassword"`
	Password        string `json:"password"`
}

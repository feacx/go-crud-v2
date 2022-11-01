package requests

type CreateUserRequest struct {
	Name  string `valid:"name" json:"name"`
	Email string `valid:"email" json:"email"`
}

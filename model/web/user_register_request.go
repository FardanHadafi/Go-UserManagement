package web

type UserRegisterRequest struct {
	Email string `validate:"required"`
	Password string `validate:"required"`
	Name string `validate:"required,min=3,max=255"`
}
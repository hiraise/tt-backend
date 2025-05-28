package request

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type VerifyRequest struct {
	Token string `validate:"required,uuid"`
}

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

package request

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@$*&#"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Age      int8   `json:"age" binding:"required,min=18,max=120"`
}

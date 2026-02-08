package dto

type UserRegistrationDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
	Bio      string `json:"bio" binding:"max=255"`
	Image    string `json:"image" binding:"omitempty,url"`
}

type UserResponseDTO struct {
	ID       uint   `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

type UserUpdateDTO struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty,min=3"`
	Bio      string `json:"bio" binding:"omitempty,max=255"`
	Image    string `json:"image" binding:"omitempty,url"`
}

type UserLoginDTO struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
}

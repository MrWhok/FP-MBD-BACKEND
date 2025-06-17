package model

type UserRegisterModel struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	NoTelp   string `json:"no_telp" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

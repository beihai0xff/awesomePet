package models

type RequestUser struct {
	Uid         uint64 `json:"uid"`
	UserName    string `json:"username"`
	Sex         uint64 `json:"sex"`
	Description string `json:"description"`
	Email       string `json:"email"`
	City        uint64 `json:"city"`
	Street      string `json:"street"`
	Password    string `json:"password"`
}

type PasswordReset struct {
	Uid         uint64 `json:"uid"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

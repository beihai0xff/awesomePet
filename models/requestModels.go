package models

type RequestUser struct {
	Uid      uint64 `json:"uid"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RequestReset struct {
	Uid         uint64 `json:"uid"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

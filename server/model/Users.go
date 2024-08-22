package model

type UserModel struct {
	Model
	User_name   string `json:"user_name"`
	Nick_name   *string `json:"nick_name"`
	Password    string `json:"password"`
	Verified  	bool `json:"verified" gorm:"default:false"`
	Login_token *string `json:"login_token"`
}

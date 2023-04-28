package models

type User struct {
	ID       uint   `form:"id"`
	Email    string `form:"email"`
	Username string `form:"username"`
	Fullname string `form:"full_name"`
	//Password string `form:"password"`
	// Salt       string `form:"password_salt"`
	CreateTime uint `form:"create_time"`
	UpdateTime uint `form:"update_time"`
}

type Register struct {
	Email         string `form:"email"`
	Username      string `form:"username"`
	Fullname      string `form:"full_name"`
	Password      string `form:"password"`
	PasswordMatch string `form:"password_match"`
}

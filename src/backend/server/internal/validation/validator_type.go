package validation

type UserRegisterRequestValidator struct {
	Username     string `form:"username" binding:"required,username"`
	Email        string `form:"email" binding:"required,email"`
	Telephone    string `form:"telephone" binding:"required,telephone"`
	Organization string `form:"organization" `
	Address      string `form:"address" `
	Postcode     string `form:"postcode" `
	Password     string `form:"password" binding:"required,password"`
}

type UserLoginRequestValidator struct {
	Telephone string `form:"telephone" binding:"required,telephone"`
	Password  string `form:"password" binding:"required"`
}

type SubmissionCreateValidator struct {
	Name                string
	FirstAuthor         int
	CorrespondingAuthor int
	SecondAuthor        int
	ThirdAuthor         int
	ForthAuthor         int
	Topic               int
	FilePath            string
	KeyWord             []string
}

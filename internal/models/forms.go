package models

type RegisterFormFields struct {
	Username string
	Email    string
	Password string
}

type RegisterForm struct {
	Values RegisterFormFields
	Errors RegisterFormFields
}

func (f RegisterForm) HasErrors() bool {
	return f.Errors.Username != "" || f.Errors.Email != "" || f.Errors.Password != ""
}

type LoginFormFields struct {
	Email    string
	Password string
}

type LoginForm struct {
	Values LoginFormFields
	Errors LoginFormFields
}

func (f LoginForm) HasErrors() bool {
	return f.Errors.Email != "" || f.Errors.Password != ""
}

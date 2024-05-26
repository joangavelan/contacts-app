package models

type SignUpFormFields struct { 
	Username string
	Email string
	Password string
}

type SignUpForm struct {
	Values SignUpFormFields
	Errors SignUpFormFields
}
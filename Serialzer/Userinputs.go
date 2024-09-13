package serialzer



type UpdatePassword struct{
	Oldpassword string `validate:"required"`
	Newpassword string  `validate:"required"`
}
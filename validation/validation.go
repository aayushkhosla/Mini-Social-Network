package validation

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

func StrongPassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    re := regexp.MustCompile(`(.*[a-z]+)(.*[A-Z]+)(.*[0-9]+)(.*[!@#$%]+)`) 
    return re.MatchString(password)
}
func Gendervalidation(fl validator.FieldLevel) bool{
	input := fl.Field().String()
	switch input {
    case "Male", "Female", "Other":
        return true
    default:
        return false
    }
}
func MaritalStatusvalidation(fl validator.FieldLevel) bool{
	input := fl.Field().String()
	switch input {
    case "Single", "Married", "Divorced", "Widowed" :
        return true
    default:
        return false
    }
}
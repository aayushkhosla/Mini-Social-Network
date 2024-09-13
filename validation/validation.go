package validation

import(
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func StrongPassword(fl validator.FieldLevel) bool {
    // password := fl.Field().String()

    // Regex to check for at least one capital letter, one lowercase letter, one digit, and one special character
    // re := regexp.MustCompile(`?=^.{6,10}$?=.*\d)(?=.*[a-z]?=.*[A-Z]?=.*[!@#$%^&amp;*()_+}{&quot;:;'?/&gt;.&lt;,])(?!.*\s).*$`)
    
	
    // return re.MatchString(password)
	return true
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
	// return true
}
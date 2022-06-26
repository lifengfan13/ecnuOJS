package validation

import (
	"fmt"
	"regexp"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidateTelephone(fl validator.FieldLevel) bool {
	telephone := fl.Field().String()
	if ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, telephone); !ok {
		return false
	}
	return true
}
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if ok, _ := regexp.MatchString(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`, email); !ok {
		return false
	}
	return true
}

func ValidateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9]{4,16}$`, username); !ok {
		return false
	}
	return true
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	re := regexp2.MustCompile(`(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9]).{8,24}`, 0)
	if ok, _ := re.MatchString(password); !ok {
		return false
	}
	return true
}

func ValidatePostcode(fl validator.FieldLevel) bool {
	postcode := fl.Field().String()
	if ok, _ := regexp.MatchString(`^\d{6}$`, postcode); !ok {
		return false
	}
	return true
}

func RegisterValidatorFunc(v *validator.Validate, tag string, fn func(fl validator.FieldLevel) bool) {
	_ = v.RegisterValidation(tag, validator.Func(fn))
}

func SetupValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("telephone", ValidateTelephone)
		if err != nil {
			return err
		}
		err = v.RegisterValidation("email", ValidateEmail)
		if err != nil {
			return err
		}
		err = v.RegisterValidation("username", ValidateUsername)
		if err != nil {
			return err
		}
		err = v.RegisterValidation("password", ValidatePassword)
		if err != nil {
			return err
		}
		err = v.RegisterValidation("postcode", ValidatePostcode)
		if err != nil {
			return err
		}

	} else {
		fmt.Println("validator-Validate rise error!")
	}
	fmt.Println("setup validator success")
	return nil
}

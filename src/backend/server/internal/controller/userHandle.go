package controller

import (
	"fmt"
	"server/internal/dao"
	"server/internal/middleware"
	"server/internal/util"
	"server/internal/validation"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandlerRegister(c *gin.Context) {

	param := validation.UserRegisterRequestValidator{}
	err := c.ShouldBind(&param)
	if err != nil {
		// log
		fmt.Println(err)
		util.RespErr(c, util.RegisterUserFormatError)
		return
	}

	isDuplicated := dao.CheckRegisterDuplicate(param.Email, param.Telephone)
	if isDuplicated != util.RequestSuccess {
		//log
		fmt.Println(isDuplicated)
		util.RespErr(c, isDuplicated)
		return
	}
	err = dao.CreateUser(&param)
	if err != nil {
		// log
		fmt.Println(err)
		util.RespErr(c, util.ServerInternalError)
		return
	}
	util.RespSuccessWithoutData(c, util.CreateSuccess)
	// log

}

func HandlerLogin(c *gin.Context) {
	param := validation.UserLoginRequestValidator{}
	err := c.ShouldBind(&param)
	if err != nil {
		// log
		fmt.Println(err)
		util.RespErr(c, util.LoginUserFormatError)
		return
	}

	user, code := dao.SearchUser(param.Telephone)
	if code != util.RequestSuccess {
		//log

		util.RespErr(c, code)
		return
	}

	if ok := checkPassword(param.Password, user.Password); !ok {
		// log
		util.RespErr(c, util.LoginUserPasswordError)
		return
	}

	token, err := middleware.GenToken(user.Telephone)
	if err != nil {
		util.RespErr(c, util.ServerInternalError)
		return
	}

	util.RespSuccess(c, struct{ Token string }{Token: token}, util.RequestSuccess)

}

func checkPassword(requestPass, pass string) bool {
	if ok := bcrypt.CompareHashAndPassword([]byte(pass), []byte(requestPass)); ok != nil {
		return false
	}
	return true
}

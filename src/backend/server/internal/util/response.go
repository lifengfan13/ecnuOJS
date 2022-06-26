package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  Error       `json:"error"`
}

var success map[int]int = map[int]int{
	RequestSuccess: http.StatusOK,
	CreateSuccess:  http.StatusCreated,
	UpdateSuccess:  http.StatusCreated,
	AcceptSuccess:  http.StatusAccepted,
	DeleteSuccess:  http.StatusNoContent,
}

func RespSuccess(c *gin.Context, data interface{}, code int) {

	respStatus := success[code]

	c.JSON(respStatus, Response{data, respStatus, Error{}})

}

func RespSuccessWithoutData(c *gin.Context, code int) {
	RespSuccess(c, struct{}{}, code)
}

// func RespErr(c *gin.Context, errCode int) {

// 	respStatus := ReturnErrStatus(errCode)
// 	respErr := ReturnErrMsg(errCode)

// 	c.JSON(respStatus, Response{struct{}{}, respStatus, *respErr})
// }

func RespErr(c *gin.Context, errCode int) {

	respStatus := ReturnErrStatus(errCode)
	respErr := ReturnErrMsg(errCode)

	c.JSON(respStatus, Response{struct{}{}, respStatus, *respErr})
}

func Errs(c *gin.Context, msg string) {
	err := Error{Msg: msg, Code: 400}
	c.JSON(400, Response{struct{}{}, http.StatusBadRequest, err})
}

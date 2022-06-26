package main

import (
	"fmt"
	"net/http"
	"server/global"
	"server/internal/routers"
	"server/internal/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		// 日志输出
		fmt.Println(err)
	}

	err = global.SetupDB()
	if err != nil {
		fmt.Println(err)
	}

	err = validation.SetupValidator()
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	gin.SetMode(global.ServerSettings.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:    ":" + global.ServerSettings.HttpPort,
		Handler: router,
	}
	s.ListenAndServe()
}

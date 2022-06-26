package controller

import (
	"fmt"
	"server/internal/dao"
	"server/internal/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandlerCreateReviewer(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		log.Error("create reviewer ID can not be empty")
		util.Errs(c, fmt.Sprintf("create reviewer ID can not be empty"))
		return
	}

	log.Debugf("create reviewer ID is: ", id)

	user, err := dao.SearchUserByID(id)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	if user == nil {
		log.Errorf("create reviewer user ID: %s not found", id)
		util.Errs(c, fmt.Sprintf("create reviewer user ID: %s not found", id))
		return
	}

	log.Debugf("create reviewer is: ", user.Username)
	dao.CreateReviewer(user)
	util.RespSuccessWithoutData(c, util.RequestSuccess)
}

func HandlerGetReviewer(c *gin.Context) {
	users := dao.GetReviewers()
	util.RespSuccess(c, users, util.RequestSuccess)
}

func HandlerDeleteReviewer(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		log.Error("delete reviewer ID can not be empty")
		util.Errs(c, fmt.Sprintf("delete reviewer ID can not be empty"))
		return
	}

	log.Debugf("delete reviewer ID is: ", id)

	user, err := dao.SearchUserByID(id)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	if user == nil {
		log.Errorf("delete reviewer user ID: %s not found", id)
		util.Errs(c, fmt.Sprintf("delete reviewer user ID: %s not found", id))
		return
	}

	log.Debugf("delete reviewer is: %s", user.Username)
	dao.DeleteReviewer(user)

	util.RespSuccessWithoutData(c, util.RequestSuccess)
}

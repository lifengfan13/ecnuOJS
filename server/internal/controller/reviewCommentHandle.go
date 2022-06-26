package controller

import (
	"server/internal/dao"
	"server/internal/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandlerCreateReviewComment(c *gin.Context) {
	var id, cmt string
	if id = c.PostForm("id"); id == "" {
		log.Error("create comment id can not be empty")
		util.Errs(c, "create comment id can not be empty")
		return
	}
	if cmt = c.PostForm("comment"); cmt == "" {
		log.Error("create comment cmt can not be empty")
		util.Errs(c, "create comment cmt can not be empty")
		return
	}

	log.Debugf("create review comment reSbm id is: %s, comment is: %s", id, cmt)

	if err := dao.CreateReviewComment(id, cmt); err != nil {
		log.Error(err)
		util.Errs(c, "create review comment failed")
		return
	}

	util.RespSuccessWithoutData(c, util.RequestSuccess)
}

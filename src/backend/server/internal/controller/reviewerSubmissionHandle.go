package controller

import (
	"fmt"
	"server/internal/dao"
	"server/internal/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandlerCreateReviewerSbm(c *gin.Context) {
	rid := c.Query("rid")
	sid := c.Query("sid")
	if rid == "" || sid == "" {
		log.Error("create reviewer ID or submission ID can not be empty")
		util.Errs(c, fmt.Sprintf("create reviewer ID or submission ID can not be empty"))
		return
	}

	log.Debugf("create reviewer ID is: ", rid)
	log.Debugf("create submission ID is: ", sid)

	user, err := dao.SearchUserByID(rid)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}
	sbm, err := dao.SearchSubmissionByID(sid)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	if user == nil || sbm == nil {
		log.Errorf("create reviewer-submission user:%s or submission:%s not found", rid, sid)
		util.Errs(c, fmt.Sprintf("create reviewer-submission user:%s or submission:%s not found", rid, sid))
		return
	}

	reSbm, err := dao.SearchReviewerSbmByRIDSID(rid, sid)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	if reSbm != nil {
		// 恢复软删除
		if reSbm.DeletedAt.Valid {
			if err := dao.RecoverReviewerSbm(reSbm); err != nil {
				log.Error(err)
				util.Errs(c, err.Error())
				return
			}
			log.Debugf("reviewer-submission reviewer is: %s, submission is: %s has been exist, recover it", rid, sid)
			util.RespSuccessWithoutData(c, util.RequestSuccess)
			return
		}
		log.Errorf("reviewer-submission reviewer is: %s, submission is: %s already exist", rid, sid)
		util.Errs(c, fmt.Sprintf("reviewer-submission reviewer is: %s, submission is: %s already exist", rid, sid))
		return
	}

	log.Debugf("create reviewer-submission reviewer is: %s, submission is: %s", rid, sid)
	dao.CreateReviewerSbm(user, sbm)

	util.RespSuccessWithoutData(c, util.RequestSuccess)
}

func HandlerGetReviewerSbms(c *gin.Context) {
	reSbm := dao.GetReviewerSbms()
	util.RespSuccess(c, reSbm, util.RequestSuccess)
}

func HandlerDeleteReviewerSbm(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		log.Error("delete reviewer-submission ID can not be empty")
		util.Errs(c, fmt.Sprintf("delete reviewer-submission ID can not be empty"))
		return
	}

	log.Debugf("delete reviewer-submission ID is: ", id)

	reSbm, err := dao.SearchReviewerSbmByID(id)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	if reSbm == nil {
		log.Errorf("delete reviewer-submission ID: %s not found", id)
		util.Errs(c, fmt.Sprintf("delete reviewer-submission ID: %s not found", id))
		return
	}

	log.Debugf("delete reviewer-submission is: %s", id)
	if err = dao.DeleteReviewerSbm(reSbm); err != nil {
		log.Error(err)
		util.Errs(c, "delete reviewer-submission failed")
		return
	}

	util.RespSuccessWithoutData(c, util.RequestSuccess)
}

package controller

import (
	"crypto/md5"
	"errors"
	"fmt"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/util"
	"server/internal/validation"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandlerCreateSubmission(c *gin.Context) {
	name := c.PostForm("name")
	sbm, err := dao.SearchSubmissionByName(name)
	if err != nil {
		log.Error(err.Error())
		util.Errs(c, err.Error())
		return
	}
	if sbm != nil {
		log.Error("submission already exists")
		util.Errs(c, fmt.Sprintf("submission already exists"))
		return
	}
	param, err := getParams(c, true)
	if err != nil {
		util.Errs(c, err.Error())
		return
	}

	err = dao.CreateSubmission(param)
	if err != nil {
		fmt.Println("create new submission err: " + err.Error())
		util.Errs(c, err.Error())
		return
	}

	sbm, err = dao.SearchSubmissionByName(param.Name)
	if err != nil {
		fmt.Println("search submission by name failed: " + err.Error())
		util.Errs(c, err.Error())
		return
	}

	keyword := c.PostForm("keyword")
	if keyword != "" {
		err = createSubmissionKeyword(keyword, sbm)
		if err != nil {
			util.Errs(c, err.Error())
			return
		}
	}

	util.RespSuccessWithoutData(c, util.CreateSuccess)
}

func HandlerUpdateSubmission(c *gin.Context) {

	name := c.PostForm("name")
	sbm, err := dao.SearchSubmissionByName(name)
	if err != nil {
		fmt.Println("search submission by name failed: " + err.Error())
		util.Errs(c, err.Error())
		return
	} else {
		if sbm == nil {
			fmt.Printf("search submission by name %s not found\n", name)
			util.Errs(c, fmt.Sprintf("search submission by name %s not found\n", name))
			return
		}
	}

	param, err := getParams(c, false)
	if err != nil {
		util.Errs(c, err.Error())
		return
	}

	dao.UpdateSubmission(param, sbm)

	keyword := c.PostForm("keyword")

	if keyword != "" {
		dao.DeleteKeywordBySbmID(sbm.ID)
		err = createSubmissionKeyword(keyword, sbm)
		if err != nil {
			util.Errs(c, err.Error())
			return
		}
	}

	util.RespSuccessWithoutData(c, util.CreateSuccess)
}

func getParams(c *gin.Context, required bool) (*validation.SubmissionCreateValidator, error) {
	param := validation.SubmissionCreateValidator{}
	param.Name = c.PostForm("name")

	param.FirstAuthor = -1
	firstAuthor := c.PostForm("author")
	user, err := dao.SearchUserByUsername(firstAuthor)
	if err != nil {
		return nil, err
	}
	if user == nil && err == nil && required {
		return nil, errors.New("first author user not found")
	}

	if user != nil {
		param.FirstAuthor = int(user.ID)
	}

	param.CorrespondingAuthor = -1
	cauthor := c.PostForm("cauthor")
	user, err = dao.SearchUserByUsername(cauthor)
	if err != nil {
		return nil, err
	}
	if user == nil && err == nil && required {
		return nil, errors.New("corresponding author user not found")
	}
	if user != nil {
		param.CorrespondingAuthor = int(user.ID)
	}

	sauthor := c.PostForm("sauthor")
	param.SecondAuthor = -1
	if sauthor != "" {
		user, err = dao.SearchUserByUsername(sauthor)
		if err == nil && user != nil {
			param.SecondAuthor = int(user.ID)
		}
	}

	tauthor := c.PostForm("tauthor")
	param.SecondAuthor = -1
	if tauthor != "" {
		user, err = dao.SearchUserByUsername(tauthor)
		if err == nil && user != nil {
			param.ThirdAuthor = int(user.ID)
		}
	}

	fauthor := c.PostForm("fauthor")
	param.SecondAuthor = -1
	if fauthor != "" {
		user, err = dao.SearchUserByUsername(fauthor)
		if err == nil && user != nil {
			param.ForthAuthor = int(user.ID)
		}
	}

	param.Topic = -1
	param.Topic, err = strconv.Atoi(c.PostForm("topic"))
	//log.
	log.Debugf("param topic is %v: ", param.Topic)
	if err != nil {
		fmt.Println("topic type convert err")
		return nil, err
	}

	param.FilePath = ""
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("get file from form err: " + err.Error())
		return nil, err
	}
	hashSeed := []byte(param.Name + firstAuthor)
	fileMD5Name := fmt.Sprintf("%x", md5.Sum(hashSeed))
	filePath := "/tmp/paper/" + string(fileMD5Name) + "-" + param.Name
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		fmt.Println("save file err: " + err.Error())
		return nil, err
	}
	param.FilePath = filePath

	return &param, nil
}

func createSubmissionKeyword(keyword string, sbm *model.Submission) error {
	keySlice := strings.Split(keyword, ",")
	for _, k := range keySlice {
		err := dao.CreateKeyword(k, int(sbm.ID))
		if err != nil {
			fmt.Printf("create keyword %s + %v: failed: ", k, sbm.ID)
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func HandlerGetSubmission(c *gin.Context) {
	var userID string
	if userID = c.Query("id"); userID == "" {
		log.Error("userID can not be empty")
		util.Errs(c, fmt.Sprintf("userID can not be empty"))
		return
	}
	log.Debugf("Query submission userID is %s", userID)

	sbms, err := dao.SearchSubmissionByUserID(userID)

	// sbm, err := dao.SearchSubmissionByName(name)
	if err != nil {
		log.Error("search submission by name err: ", err)
		util.Errs(c, "server error")
		return
	}

	util.RespSuccess(c, sbms, util.RequestSuccess)
}

func HandlerDeleteSubmission(c *gin.Context) {
	sbmID := c.Query("id")

	if sbmID == "" {
		log.Error("delete submission ID can not be empty")
		util.Errs(c, fmt.Sprintf("delete submission ID can not be empty"))
		return
	}

	log.Debugf("delete submission ID is: ", sbmID)

	sbm, err := dao.SearchSubmissionByID(sbmID)
	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	if sbm == nil || sbm.IsDel == 1 {
		log.Error("delete submission not exists")
		util.Errs(c, fmt.Sprintf("delete submission not exists"))
		return
	}

	dao.DeleteSubmission(sbm)
	log.Debugf("successfully delete submission: ", sbm.Name)
	util.RespSuccessWithoutData(c, util.RequestSuccess)
}

func HandlerSearchSubmission(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		log.Error("search name can not be empty")
		util.Errs(c, fmt.Sprintf("search name can not be empty"))
		return
	}

	log.Debugf("search name is: ", name)

	sbms, err := dao.SearchSubmissionByNameVague(name)

	if err != nil {
		log.Error(err)
		util.Errs(c, err.Error())
		return
	}

	var retSbms *[]model.Submission

	for _,sbm := range *sbms {
		if s
	}

}

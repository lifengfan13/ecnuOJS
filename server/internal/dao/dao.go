package dao

import (
	"server/global"
	"server/internal/model"
	"server/internal/util"
	"server/internal/validation"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// register

func CreateUser(param *validation.UserRegisterRequestValidator) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(param.Password), 14)
	if err != nil {
		return err
	}
	newUser := model.User{
		Username:     param.Username,
		Email:        param.Email,
		Telephone:    param.Telephone,
		Organization: param.Organization,
		Address:      param.Address,
		Postcode:     param.Address,
		Password:     string(passwordHash),
		Role:         1,
		IsDel:        0,
	}
	return global.DB.Create(&newUser).Error
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckRegisterDuplicate(email, telephone string) int {

	errEmail := global.DB.Where("email = ?", email).First(&model.User{}).Error
	if errEmail == nil {
		return util.RegisterEmailDuplicated
	}

	errTel := global.DB.Where("telephone = ?", telephone).First(&model.User{}).Error
	if errTel == nil {
		return util.RegisterTelephoneDuplicated
	}
	return util.RequestSuccess
}

func SearchUser(telephone string) (*model.User, int) {
	var user model.User
	err := global.DB.Where("telephone = ?", telephone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.LoginUserNotExist
		}
		// log
		return nil, util.ServerInternalError
	}
	return &user, util.RequestSuccess
}

func SearchUserByUsername(name string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ?", name).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func SearchUserByID(id string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

//CreateUser(param *validation.UserRegisterRequestValidator) error {
func CreateSubmission(param *validation.SubmissionCreateValidator) error {
	newSubmission := model.Submission{
		Name:                param.Name,
		FirstAuthor:         param.FirstAuthor,
		CorrespondingAuthor: param.CorrespondingAuthor,
		SecondAuthor:        param.SecondAuthor,
		ThirdAuthor:         param.ThirdAuthor,
		ForthAuthor:         param.ForthAuthor,
		Topic:               param.Topic,
		Status:              0, // 0 :submitted
		FilePath:            param.FilePath,
		IsDel:               0,
	}
	return global.DB.Create(&newSubmission).Error
}

func SearchSubmissionByName(name string) (*model.Submission, error) {
	var sbm model.Submission
	err := global.DB.Where("name = ?", name).First(&sbm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sbm, nil
}

func SearchSubmissionByNameVague(name string) (*[]model.Submission, error) {
	var sbms []model.Submission
	err := global.DB.Where("name = ?", "%"+name+"%").Find(&sbms).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sbms, nil
}

func SearchSubmissionByID(id string) (*model.Submission, error) {
	var sbm model.Submission
	err := global.DB.Where("id = ?", id).First(&sbm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sbm, nil
}

func DeleteSubmission(sbm *model.Submission) {
	global.DB.Delete(&sbm)
	sbm.IsDel = 1
	global.DB.Save(&sbm)
}

func SearchSubmissionByUserID(id string) (*[]model.Submission, error) {
	var sbms []model.Submission
	err := global.DB.Where("first_author = ?", id).Find(&sbms).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sbms, nil
}

func UpdateSubmission(param *validation.SubmissionCreateValidator, sbm *model.Submission) {
	if param.FirstAuthor != -1 {
		sbm.FirstAuthor = param.FirstAuthor
	}
	if param.CorrespondingAuthor != -1 {
		sbm.CorrespondingAuthor = param.CorrespondingAuthor
	}
	if param.SecondAuthor != -1 {
		sbm.SecondAuthor = param.SecondAuthor
	}
	if param.ThirdAuthor != -1 {
		sbm.ThirdAuthor = param.ThirdAuthor
	}
	if param.ForthAuthor != -1 {
		sbm.ForthAuthor = param.ForthAuthor
	}
	if param.Topic != -1 {
		sbm.Topic = param.Topic
	}
	if param.FilePath != "" {
		sbm.FilePath = param.FilePath
	}
	global.DB.Save(&sbm)
}

func CreateKeyword(keyword string, sbmID int) error {
	kw := model.Keyword{
		Name:         keyword,
		SubmissionId: sbmID,
		IsDel:        0,
	}
	return global.DB.Create(&kw).Error
}

func DeleteKeywordBySbmID(sbmID uint) {
	global.DB.Where("submission_id = ?", sbmID).Delete(&model.Keyword{})
}

func CreateReviewer(user *model.User) {
	user.Role = 2
	global.DB.Save(&user)
}

func GetReviewers() *[]model.User {
	var users []model.User
	global.DB.Where("role = ?", 2).Find(&users)
	return &users
}

func DeleteReviewer(user *model.User) {
	user.Role = 1
	global.DB.Save(&user)
}

func CreateReviewerSbm(user *model.User, sbm *model.Submission) {

	reSbm := model.ReviewerSubmission{
		ReviewerId:    int(user.ID),
		SubmissionID:  int(sbm.ID),
		ReviewComment: "",
		Status:        0,
	}

	global.DB.Save(&reSbm)
}

func GetReviewerSbms() *[]model.ReviewerSubmission {
	var reSbms []model.ReviewerSubmission
	global.DB.Find(&reSbms)
	return &reSbms
}

func SearchReviewerSbmByRIDSID(rid, sid string) (*model.ReviewerSubmission, error) {
	var resbm model.ReviewerSubmission
	err := global.DB.Unscoped().Where("reviewer_id = ? AND submission_id = ?", rid, sid).First(&resbm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resbm, nil
}

func SearchReviewerSbmByRID(rid string) (*[]model.ReviewerSubmission, error) {
	var resbm []model.ReviewerSubmission
	err := global.DB.Where("reviewer_id = ?", rid).Find(&resbm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resbm, nil
}

func SearchReviewerSbmBySID(sid string) (*[]model.ReviewerSubmission, error) {
	var resbm []model.ReviewerSubmission
	err := global.DB.Where("submission_id = ?", sid).Find(&resbm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resbm, nil
}

func SearchReviewerSbmByID(id string) (*model.ReviewerSubmission, error) {
	var resbm model.ReviewerSubmission
	err := global.DB.Where("id = ?", id).First(&resbm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resbm, nil
}

func DeleteReviewerSbm(reSbm *model.ReviewerSubmission) error {
	return global.DB.Delete(reSbm).Error
}

func RecoverReviewerSbm(reSbm *model.ReviewerSubmission) error {
	return global.DB.Unscoped().Model(&reSbm).Update("deleted_at", nil).Error
}

func CreateReviewComment(id, cmt string) error {
	var reSbm model.ReviewerSubmission
	if err := global.DB.Where("id = ?", id).First(&reSbm).Error; err != nil {
		return err
	}

	reSbm.Status = 1
	reSbm.ReviewComment = cmt
	return global.DB.Save(&reSbm).Error
}

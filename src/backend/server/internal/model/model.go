package model

import (
	"fmt"
	"server/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Model struct {
// 	ID         uint32 `gorm:"primary_key" json:"id"`
// 	CreatedOn  uint32 `josn:"created_on"`
// 	ModifiedOn uint32 `json:"modified_on"`
// 	DeletedOn  uint32 `json:"deleted_on"`
// 	IsDel      uint8  `json:"is_del"`
// }

type User struct {
	gorm.Model
	Username     string `gorm:"column:username"`
	Email        string `gorm:"uniqueIndex,column:email"`
	Telephone    string `gorm:"uniqueIndex,column:telephone"`
	Organization string `gorm:"column:organization"`
	Address      string `gorm:"column:address"`
	Postcode     string `gorm:"column:postcode"`
	Password     string `gorm:"column:password_hash"`
	Role         uint8  `gorm:"column:role"`
	IsDel        uint8  `gorm:"column:is_del"`
}

func (u User) TableName() string {
	return "user"
}

type Submission struct {
	gorm.Model
	Name                string `gorm:"column:name"`
	FirstAuthor         int    `gorm:"column:first_author"`
	CorrespondingAuthor int    `gorm:"column:corresponding_author"`
	SecondAuthor        int    `gorm:"column:second_author"`
	ThirdAuthor         int    `gorm:"column:third_author"`
	ForthAuthor         int    `gorm:"column:forth_author"`
	Topic               int    `gorm:"column:topic"`
	Status              int    `gorm:"column:status"`
	// 已经投稿（待审核）、被拒稿、被接受
	FilePath string `gorm:"column:file_path"`
	IsDel    uint8  `gorm:"column:is_del"`
}

func (s Submission) TableName() string {
	return "submission"
}

type ReviewerSubmission struct {
	gorm.Model
	ReviewerId    int    `gorm:"column:reviewer_id"`
	SubmissionID  int    `gorm:"column:submission_id"`
	ReviewComment string `gorm:"column:review_comment"`
	Status        int    `gorm:"column:status"`
	// 0: 提交，带评审； 1:已评审 
}

func (r ReviewerSubmission) TableName() string {
	return "reviewer_submission"
}

type Keyword struct {
	gorm.Model
	Name         string `gorm:"column:name"`
	SubmissionId int    `gorm:"column:submission_id"`
	IsDel        uint8  `gorm:"column:is_del"`
}

func (r Keyword) TableName() string {
	return "key_word"
}

func NewDBEngine(dbSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=%v&loc=Local",
		dbSetting.UserName, dbSetting.Password, dbSetting.Host, dbSetting.DBName, dbSetting.Charset, dbSetting.ParserTime,
	)
	fmt.Println(dsn)
	cfg := &gorm.Config{}
	// if ServerSettings.RunMode == "debug" {
	// 	cfg.Logger = logger.Default.LogMode(logger.Info)
	// }
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

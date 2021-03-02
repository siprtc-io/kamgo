package model

import (
	"fmt"
	"github.com/cyruzin/tome"
	"kamgo/model/orm"
)

type KamSubscriber struct {
	orm.Model  `gorm:"-" json:"-"`
	Id         int    `gorm:"column:id" json:"id"`
	Username   string `gorm:"column:username" json:"username"`
	DomainName string `gorm:"column:domain" json:"domain"`
	Password   string `gorm:"column:password" json:"password"`
	Ha1        string `gorm:"column:ha1" json:"ha1"`
	Ha1b       string `gorm:"column:ha1b" json:"ha1b"`
}

type sipEndpoints []*KamSubscriber

// Result type is a struct of posts with pagination.
type Results struct {
	Data *sipEndpoints `json:"subscribers"`
	*tome.Chapter
}

func (kamSub *KamSubscriber) Create() error {
	if err := KamDB().Create(kamSub).Error; err != nil {
		return err
	}
	return nil
}

func (kamSub *KamSubscriber) GetKamSubModel(domain,username string) error {
	if err := KamDB().Where(KamSubscriber{DomainName: domain, Username: username}).First(kamSub).Error; err != nil {
		return  err
	}
	return nil
}

func subscribersCount() int64 {
	var count int64
	sipEndPoint := KamSubscriber{}
	KamDB().Model(&sipEndPoint).Count(&count)
	return count
}

func (kamSub *KamSubscriber) GetAllSipEndpoints(requestURI string, pageNumber int) *Results {
	baseURL := fmt.Sprintf("http://%s", requestURI)
	totalCount := int(subscribersCount())
	// Creating a tome chapter with links.
	chapter := &tome.Chapter{
		// Setting base URL.
		BaseURL: baseURL,
		// Enabling link results.
		Links: true,
		// Page that you captured in params inside you handler.
		NewPage: pageNumber,
		// Total of pages, this usually comes from a SQL query total rows result.
		TotalResults: totalCount,
		Limit:        30,
	}
	// Paginating the results.
	if err := chapter.Paginate(); err != nil {
		return nil
	}

	var sipEPs sipEndpoints
	if pageNumber == 1 {
		KamDB().Order("id desc").Limit(chapter.Limit).
			Find(&sipEPs)
	} else {
		KamDB().Order("id desc").Offset(pageNumber).Limit(chapter.Limit).
			Find(&sipEPs)
	}
	// Mocking results with pagination.
	res := &Results{Data: &sipEPs, Chapter: chapter}
	return res
}


func (kamSub *KamSubscriber) KamUpdate(domain,username string) error {
	if err := KamDB().Where(KamSubscriber{DomainName: domain, Username: username}).Select("username","password").Updates(kamSub).Error; err != nil {
		return err
	}
	return nil
}

//Delete sip endpoint
func (kamSub *KamSubscriber) DeleteKamSubModel(domain, username string) error {
	if err := KamDB().Where(KamSubscriber{DomainName: domain, Username: username}).
		Delete(kamSub).Error; err != nil {
		return err
	}
	return nil
}

func (kamSub KamSubscriber) TableName() string {
	return "subscriber"
}



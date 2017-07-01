package models

import (
	"github.com/astaxie/beego/orm"
)

type Url struct {
	Id               int
	Name             string
	SourceFullUrl    string
	SourceMainDomain string
	UserId           int
	Status           int
	CreatedAt        string
	LastVistTime     string
}

func (url *Url) TableName() string {
	return TableName("url")
}

func FindOne(id int) (*Url, error) {
	url := &Url{
		Id: id,
	}
	err := orm.NewOrm().Read(url)
	if err != nil {
		return nil, err
	}
	return url, err
}

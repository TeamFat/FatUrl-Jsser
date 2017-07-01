package models

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

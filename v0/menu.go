package colonycore

import (
	"github.com/eaciit/orm/v1"
	//"github.com/eaciit/toolkit"
)

func (mn *Menu) TableName() string {
	return "menus"
}

func (mn *Menu) RecordID() interface{} {
	return mn.ID
}

type Menu struct {
	orm.ModelBase
	ID        string  `json:"_id",bson:"_id"`
	AccessId  string  `json:"accessId",bson:"accessId"`
	Title     string  `json:"title",bson:"title"`
	Childrens []*Menu `json:"childrens",bson:"childrens"`
	Link      string  `json:"link",bson:"link"`
}

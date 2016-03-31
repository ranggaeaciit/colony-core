package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Login struct {
	orm.ModelBase
	ID       string `json:"_id",bson:"_id"`
	Password string
	Salt     string
}

func (a *Login) TableName() string {
	return "logins"
}

func (a *Login) RecordID() interface{} {
	return a.ID
}

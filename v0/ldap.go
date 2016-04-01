package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Ldap struct {
	orm.ModelBase
	ID       string `json:"_id",bson:"_id"`
	Address  string
	BaseDN   string
	Filter   string
	Username string
	Password string
}

func (a *Ldap) TableName() string {
	return "ldap"
}

func (a *Ldap) RecordID() interface{} {
	return a.ID
}

package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Ldap struct {
	orm.ModelBase
	ID              string `json:"_id",bson:"_id"`
	Address         string
	BaseDN          string
	FilterUser      string
	FilterGroup     string
	Username        string
	Password        string
	AttributesUser  []string
	AttributesGroup []string
}

func (a *Ldap) TableName() string {
	return "ldap"
}

func (a *Ldap) RecordID() interface{} {
	return a.ID
}

func (a *Ldap) RecordID2() interface{} {
	return a.ID
}

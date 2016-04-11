package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
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

func (a *Login) GetACLConnectionInfo() (string, *dbox.ConnectionInfo) {
	conf, err := toolkit.ToM(GetConfig(CONF_DB_ACL))
	if err != nil {
		return "", nil
	}

	ci := dbox.ConnectionInfo{
		conf.GetString("host"),
		conf.GetString("db"),
		conf.GetString("user"),
		conf.GetString("pass"),
		toolkit.M{}.Set("timeout", 3),
	}

	return conf.GetString("driver"), &ci
}

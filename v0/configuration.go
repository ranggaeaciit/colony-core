package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
)

const (
	CONF_DB_ACL string = "db_acl"
)

type Configuration struct {
	orm.ModelBase
	ID    string `json:"_id",bson:"_id"`
	Value interface{}
}

func (a *Configuration) TableName() string {
	return "configurations"
}

func (a *Configuration) RecordID() interface{} {
	return a.ID
}

func GetConfig(key string, args ...string) interface{} {
	var res interface{} = nil
	if len(args) > 0 {
		res = args[0]
	}

	cursor, err := Find(new(Configuration), dbox.Eq("_id", key))
	if err != nil {
		return res
	}

	if cursor.Count() == 0 {
		return res
	}

	data := Configuration{}
	err = cursor.Fetch(&data, 1, false)
	if err != nil {
		return res
	}

	return data.Value
}

func SetConfig(key string, value interface{}) {
	o := new(Configuration)
	o.ID = key
	o.Value = value
	Save(o)
}

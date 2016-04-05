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
	Key   string
	Value interface{}
}

func (a *Configuration) TableName() string {
	return "configurations"
}

func (a *Configuration) RecordID() interface{} {
	return a.ID
}

func GetConfig(key string) interface{} {
	cursor, err := Find(new(Configuration), dbox.Eq("Key", key))
	if err != nil {
		return nil
	}

	if cursor.Count() == 0 {
		return nil
	}

	data := Configuration{}
	err = cursor.Fetch(&data, 1, false)
	if err != nil {
		return nil
	}

	return data.Value
}

func SetConfig(key string, value interface{}) {
	o := new(Configuration)
	o.Key = key
	o.Value = value
	Save(o)
}

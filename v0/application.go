package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Application struct {
	orm.ModelBase
	ID       string `json:"_id",bson:"_id"`
	AppsName string `json:"AppsName", bson:"AppsName"`
	Enable   bool   `json:"Enable", bson:"Enable"`
	AppPath  string `json:"AppPath", bson:"AppPath"`
}

func (a *Application) TableName() string {
	return "applications"
}

func (a *Application) RecordID() interface{} {
	return a.ID
}

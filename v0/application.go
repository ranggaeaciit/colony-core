package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Application struct {
	orm.ModelBase `json:"-"`
	ID            string `json:"_id"`
	Enable        bool
}

func (a *Application) TableName() string {
	return "applications"
}

func (a *Application) RecordID() interface{} {
	return a.ID
}

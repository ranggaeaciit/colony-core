package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Application struct {
	orm.ModelBase
	ID     string `json:"_id"`
	Enable bool
}

func (a *Application) TableName() string {
	return "applications"
}

var ctxAppn *orm.DataContext

func CtxApplication() *orm.DataContext {
	if ctxAppn == nil {
		c, _ := getConnection(&Application{})
		ctxAppn = orm.New(c)
	}
	return ctxAppn
}

package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

type Connection struct {
	ID                               string `json:"_id",bson:"_id"`
	Driver, Host, UserName, Password string
	Settings                         toolkit.M
}

var ctxConn *orm.DataContext

func CtxConnection() *orm.DataContext {
	if ctxConn == nil {
		c, _ := getConnection(&Connection{})
		ctxConn = orm.New(c)
	}
	return ctxConn
}

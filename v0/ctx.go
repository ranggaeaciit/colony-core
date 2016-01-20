package colonycore

import (
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/jsons"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

var _conn dbox.IConnection
var _ctx *orm.DataContext
var _ctxErr error

func ctx() {
	if _ctx == nil {
		_conn = getConnection()
		_ctx = orm.New(_conn)
	}
	return _ctx
}

func Save(o orm.IModel) error {
	e = initCtx()
	return e
}

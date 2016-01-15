package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

type DataSource struct {
	ID           string `json:"_id",bson:"_id"`
	ConnectionID string
	QueryInfo    toolkit.M
	MetaData     map[string]*FieldInfo
}

type FieldInfo struct {
	ID     string
	Label  string
	Type   string
	Format string
	Lookup *Lookup
}

type Lookup struct {
	ID                    string
	DataSourceID          string
	IDField, DisplayField string
	LookupFields          []string
}

var ctxDs, ctxLookup *orm.DataContext

func CtxDataSource() *orm.DataContext {
	if ctxDs == nil {
		c, _ := getConnection(&DataSource{})
		ctxDs = orm.New(c)
	}
	return ctxDs
}

func CtxLookup() *orm.DataContext {
	if CtxLookup == nil {
		c, _ := getConnection(&Lookup{})
		ctxLookup = orm.New(c)
	}
	return ctxLookup
}

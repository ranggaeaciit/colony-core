package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

type DataSource struct {
	orm.ModelBase
	ID           string `json:"_id",bson:"_id"`
	ConnectionID string
	QueryInfo    toolkit.M
	MetaData     map[string]*FieldInfo
}

func (ds *DataSource) TableName() string {
	return "datasources"
}

func (ds *DataSource) RecordID() interface{} {
	return ds.ID
}

type FieldInfo struct {
	ID     string
	Label  string
	Type   string
	Format string
	Lookup *Lookup
}

type Lookup struct {
	orm.ModelBase
	ID                    string
	DataSourceID          string
	IDField, DisplayField string
	LookupFields          []string
}

func (l *Lookup) TableName() string {
	return "lookups"
}

func (l *Lookup) RecordID() interface{} {
	return l.ID
}

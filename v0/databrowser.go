package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataBrowser struct {
	orm.ModelBase
	ID           string `json:"_id",bson:"_id"`
	BrowserName  string
	Description  string
	ConnectionID string
	TableNames   string
	QueryType    string
	QueryText    string
	MetaData     []*StructInfo
}

type StructInfo struct {
	Field         string
	Label         string
	Format        string
	Align         string
	ShowIndex     int
	Sortable      bool
	SimpleFilter  bool
	AdvanceFilter bool
	Aggregate     string
}

func (b *DataBrowser) TableName() string {
	return "databrowser"
}

func (b *DataBrowser) RecordID() interface{} {
	return b.ID
}

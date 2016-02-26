package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataBrowser struct {
	orm.ModelBase
	ID          string `json:"_id",bson:"_id"`
	BrowserName string
	IDDetails 	string
	Description string
	Connection 	string
	DataBase 	string
	TableNames 	string
	QueryType 	string
	QueryText 	string
	Struct     	[]*StructInfo

}

type StructInfo struct {
	Field     		string 
	Label  			string
	Format   		string
	Align 			string
	ShowIndex 		int32
	Sortable 		bool
	SimpleFilter 	bool
	AdvanceFilter	bool
	Aggregate 		string
}

func (b *DataBrowser) TableName() string {
	return "DataBrowser"
}

func (b *DataBrowser) RecordID() interface{} {
	return b.ID
}

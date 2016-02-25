package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataBrowser struct {
	orm.ModelBase
	ID          string `json:"_id",bson:"_id"`
	BrowserName string
}

func (b *DataBrowser) TableName() string {
	return "DataBrowser"
}

func (b *DataBrowser) RecordID() interface{} {
	return b.ID
}

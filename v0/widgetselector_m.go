package colonycore

import (
	// "github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	// "os"
	"path/filepath"
)

type Selector struct {
	orm.ModelBase
	ID               string         `json:"_id"`
	MasterDataSource string         `json:"masterDataSource"`
	Title            string         `json:"title"`
	Fields           []*FieldDetail `json:"fields"`
}

type FieldDetail struct {
	ID         string `json:"_id"`
	DataSource string `json:"dataSource"`
	Field      string `json:"field"`
}

func (sl *Selector) TableName() string {
	return filepath.Join("widget", "selectors")
}

func (sl *Selector) RecordID() interface{} {
	return sl.ID
}

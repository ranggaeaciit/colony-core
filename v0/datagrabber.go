package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataGrabber struct {
	orm.ModelBase
	ID                      string `json:"_id",bson:"_id"`
	DataSourceOrigin        string
	DataSourceDestination   string
	IgnoreFieldsOrigin      []string
	IgnoreFieldsDestination []string
	Map                     []*Maps
}

type Maps struct {
	FieldOrigin      string `json:"fieldOrigin"`
	FieldDestination string `json:"fieldDestination"`
}

func (c *DataGrabber) TableName() string {
	return "datagrabbers"
}

func (c *DataGrabber) RecordID() interface{} {
	return c.ID
}

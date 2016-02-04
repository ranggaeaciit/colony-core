package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataGrabber struct {
	orm.ModelBase
	ID                      string `json:"_id",bson:"_id"`
	DataSourceOrigin        string
	DataSourceDestination   string
	IgnoreFieldsDestination []string
	IntervalType            string
	GrabInterval            int32
	TimeoutInterval         int32
	Map                     []*Maps
	RunAt                   []string
}

type Maps struct {
	FieldOrigin      string
	FieldDestination string
}

func (c *DataGrabber) TableName() string {
	return "datagrabbers"
}

func (c *DataGrabber) RecordID() interface{} {
	return c.ID
}

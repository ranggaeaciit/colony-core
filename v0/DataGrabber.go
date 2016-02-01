package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Maps struct {
	fieldOrigin		string
	fieldDestination	string
}

type DataGrabber struct {
	orm.ModelBase
	ID		string
	DataSourceOrigin	string
	DataSourceDestination	string
	Map []*Maps
}

func (c *DataGrabber) TableName() string {
	return "DataGrabber"
}

func (c *DataGrabber) RecordID() interface{} {
	return c.ID
}

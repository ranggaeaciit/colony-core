package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataGrabber struct {
	orm.ModelBase
	ID						string	 `json:"_id",bson:"_id"`
	DataSourceOrigin		string
	DataSourceDestination	string
	Map 					[]*Maps
}

type Maps struct {
	fieldOrigin			string `json:"fieldOrigin"`
	fieldDestination	string `json:"fieldDestination"`
}

func (c *DataGrabber) TableName() string {
	return "datagrabber"
}

func (c *DataGrabber) RecordID() interface{} {
	return c.ID
}

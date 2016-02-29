package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type DataGrabber struct {
	orm.ModelBase
	ID string `json:"_id",bson:"_id"`

	DataSourceOrigin      string
	DataSourceDestination string

	ConnectionOrigin      string
	ConnectionDestination string
	TableOrigin           string
	TableDestination      string

	UseInterval         bool
	IntervalType        string
	GrabInterval        int32
	TimeoutInterval     int32
	Maps                []*Map
	RunAt               []string
	PreTransferCommand  string
	PostTransferCommand string
}

type Map struct {
	Destination     string
	DestinationType string
	Source          string
	SourceType      string
}

func (c *DataGrabber) TableName() string {
	return "datagrabbers"
}

func (c *DataGrabber) RecordID() interface{} {
	return c.ID
}

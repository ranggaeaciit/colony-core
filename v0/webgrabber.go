package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

type WebGrabber struct {
	orm.ModelBase
	ID                string `json:"_id",bson:"_id"`
	CallType          string
	SourceType        string
	IntervalType      string
	GrabInterval      int32
	TimeoutInterval   int32
	URL               string
	LogConfiguration  *LogConfiguration
	DataSetting       *DataSetting
	GrabConfiguration toolkit.M
	Parameter         []*Parameter
}

func (ds *WebGrabber) TableName() string {
	return "webgrabbers"
}

func (ds *WebGrabber) RecordID() interface{} {
	return ds.ID
}

type LogConfiguration struct {
	FileName    string
	FilePattern string
	LogPath     string
}

type ConnectionInfo struct {
	dbox.ConnectionInfo
	Collection string
}

type DataSetting struct {
	RowSelector     string
	FilterCondition toolkit.M
	ColumnSettings  []*ColumnSetting

	RowDeleteCondition  toolkit.M
	RowIncludeCondition toolkit.M

	ConnectionInfo  *ConnectionInfo
	DestinationType string
	Name            string
}

type ColumnSetting struct {
	Alias     string
	Index     int
	Selector  string
	ValueType string
}

type Parameter struct {
	Format  string
	Key     string
	Pattern string
	Value   interface{}
}

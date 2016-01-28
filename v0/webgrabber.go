package colonycore

import (
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
	DataSettings      []*DataSetting
	GrabConfiguration *GrabConfiguration
	Parameter         []*Parameter
}

func (ds *WebGrabber) TableName() string {
	return "webgrabber"
}

func (ds *WebGrabber) RecordID() interface{} {
	return ds.ID
}

type LogConfiguration struct {
	FileName    string
	FilePattern string
	LogPath     string
}

type DataSetting struct {
	ColumnSettings     []*ColumnSetting
	ConnectionInfo     *ConnectionInfo
	DestinationType    string
	Name               string
	RowDeleteCondition toolkit.M
	RowSelector        string
}

type ConnectionInfo struct {
	Collection string
	Database   string
	Host       string
}

type ColumnSetting struct {
	Alias    string
	Index    int
	Selector string
}

type GrabConfiguration struct {
	Data toolkit.M
}

type Parameter struct {
	Format  string
	Key     string
	Pattern string
	Value   interface{}
}

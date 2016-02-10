package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

type WebGrabber struct {
	orm.ModelBase
	ID                string            `json:"_id",bson:"_id"`
	IDBackup          string            `json:"nameid",bson:"nameid"`
	CallType          string            `json:"calltype",bson:"calltype"`
	SourceType        string            `json:"sourcetype",bson:"sourcetype"`
	IntervalType      string            `json:"intervaltype",bson:"intervaltype"`
	GrabInterval      int               `json:"grabinterval",bson:"grabinterval"`
	TimeoutInterval   int               `json:"timeoutinterval",bson:"timeoutinterval"`
	URL               string            `json:"url",bson:"url"`
	LogConfiguration  *LogConfiguration `json:"logconf",bson:"logconf"`
	DataSettings      []*DataSetting    `json:"datasettings",bson:"datasettings"`
	GrabConfiguration toolkit.M         `json:"grabconf",bson:"grabconf"`
}

func (ds *WebGrabber) TableName() string {
	return "webgrabbers"
}

func (ds *WebGrabber) RecordID() interface{} {
	return ds.ID
}

type LogConfiguration struct {
	FileName    string `json:"filename",bson:"filename"`
	FilePattern string `json:"filepattern",bson:"filepattern"`
	LogPath     string `json:"logpath",bson:"logpath"`
}

type ConnectionInfo struct {
	Host         string    `json:"host",bson:"host"`
	Database     string    `json:"database",bson:"database"`
	UserName     string    `json:"username",bson:"username"`
	Password     string    `json:"password",bson:"password"`
	Settings     toolkit.M `json:"settings",bson:"settings"`
	Collection   string    `json:"collection",bson:"collection"`
	ConnectionID string    `json:"connectionid",bson:"connectionid"`
}

type DataSetting struct {
	RowSelector         string           `json:"rowselector",bson:"rowselector"`
	ColumnSettings      []*ColumnSetting `json:"columnsettings",bson:"columnsettings"`
	RowDeleteCondition  toolkit.M        `json:"rowdeletecond",bson:"rowdeletecond"`
	RowIncludeCondition toolkit.M        `json:"rowincludecond",bson:"rowincludecond"`
	ConnectionInfo      *ConnectionInfo  `json:"connectioninfo",bson:"connectioninfo"`
	DestinationType     string           `json:"desttype",bson:"desttype"`
	Name                string           `json:"name",bson:"name"`
}

type ColumnSetting struct {
	Alias     string `json:"alias",bson:"alias"`
	Index     int    `json:"index",bson:"index"`
	Selector  string `json:"selector",bson:"selector"`
	ValueType string `json:"valuetype",bson:"valuetype"`
}

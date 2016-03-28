package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
)

func (ds *WebGrabber) TableName() string {
	return "webgrabbers"
}

func (ds *WebGrabber) RecordID() interface{} {
	return ds.ID
}

type WebGrabber struct {
	orm.ModelBase
	ID           string          `json:"_id",bson:"_id"`
	SourceType   string          `json:"sourcetype",bson:"sourcetype"`
	GrabConf     toolkit.M       `json:"grabconf",bson:"grabconf"`
	IntervalConf *IntervalConf   `json:"intervalconf",bson:"intervalconf"`
	LogConf      *LogConf        `json:"logconf",bson:"logconf"`
	HistConf     *HistConf       `json:"histconf",bson:"histconf"`
	DataSettings []*DataSettings `json:"datasettings",bson:"datasettings"`
	Running      bool            `json:"running",bson:"running"`
}

type IntervalConf struct {
	StartTime       string    `json:"starttime",bson:"starttime"`
	ExpiredTime     string    `json:"expiredtime",bson:"expiredtime"`
	IntervalType    string    `json:"intervaltype",bson:"intervaltype"`
	GrabInterval    int       `json:"grabinterval",bson:"grabinterval"`
	TimeoutInterval int       `json:"timeoutinterval",bson:"timeoutinterval"`
	CronConf        toolkit.M `json:"cronconf",bson:"cronconf"`
}

type LogConf struct {
	LogPath     string `json:"logpath",bson:"logpath"`
	FileName    string `json:"filename",bson:"filename"`
	FilePattern string `json:"filepattern",bson:"filepattern"`
}

type HistConf struct {
	Histpath    string `json:"histpath",bson:"histpath"`
	RecPath     string `json:"recpath",bson:"recpath"`
	FileName    string `json:"filename",bson:"filename"`
	FilePattern string `json:"filepattern",bson:"filepattern"`
}

type DataSettings struct {
	Nameid         string            `json:"nameid",bson:"nameid"`
	RowSelector    string            `json:"rowselector",bson:"rowselector"`
	ColumnSettings []*ColumnSettings `json:"columnsettings",bson:"columnsettings"`

	LimitRow       toolkit.M       `json:"limitrow",bson:"limitrow"`
	FilterCond     toolkit.M       `json:"filtercond",bson:"filtercond"`
	DestOutputType string          `json:"destoutputtype",bson:"destoutputtype"`
	DestType       string          `json:"desttype",bson:"desttype"`
	ConnectionInfo *ConnectionInfo `json:"connectioninfo",bson:"connectioninfo"`
}

type ColumnSettings struct {
	Index     int    `json:"index",bson:"index"`
	Alias     string `json:"alias",bson:"alias"`
	Selector  string `json:"selector",bson:"selector"`
	ValueType string `json:"valuetype",bson:"valuetype"`
	AttrName  string `json:"attrname",bson:"attrname"`
}

type ConnectionInfo struct {
	ConnectionID string    `json:"connectionid",bson:"connectionid"`
	Host         string    `json:"host",bson:"host"`
	UserName     string    `json:"username",bson:"username"`
	Password     string    `json:"password",bson:"password"`
	Database     string    `json:"database",bson:"database"`
	Collection   string    `json:"collection",bson:"collection"`
	Settings     toolkit.M `json:"settings",bson:"settings"`
}

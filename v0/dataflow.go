// Package colonycore define the struct that needed by other package in colony
package colonycore

import (
	/*"github.com/eaciit/errorlib"
		// "github.com/eaciit/orm/v1"
		// "log"
		"mime/multipart"
		"strconv"
		"strings"
	    "github.com/eaciit/toolkit"
	*/
	"time"

	"github.com/eaciit/orm/v1"
	tk "github.com/eaciit/toolkit"
)

const (
/*ACTION_TYPE_HIVE     = "HIVE"
ACTION_TYPE_HDFS     = "HDFS"
ACTION_TYPE_SPARK    = "SPARK"
ACTION_TYPE_DECISION = "DECISION"
ACTION_TYPE_SSH      = "SSH"
ACTION_TYPE_KAFKA    = "KAFKA"

FORK_TYPE_ALL       = "ALL"
FORK_TYPE_ONE       = "ONE"
FORK_TYPE_MANDATORY = "MANDATORY"

SSH_OPERATION_MKDIR = "MKDIR"*/
// SSH_OPERATION_* please define
)

// DataFlow to define the flow name and description
// and also have the list of the actions inside the flow
// Actions list of action, the content of the action can be FlowAction or "list of FlowAction" -> for fork action
type DataFlow struct {
	orm.ModelBase
	ID           string       `json:"_id",bson:"_id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	CreatedDate  time.Time    `json:"createddate"`
	LastModified time.Time    `json:"lastmodified"`
	CreatedBy    string       `json:"createdby"`
	Actions      []FlowAction `json:"actions"`
	DataShapes   tk.M         `json:"datashapes"`
	GlobalParam  tk.M         `json:"globalparam"`
}

func (c *DataFlow) TableName() string {
	return "dataflow"
}

func (c *DataFlow) RecordID() interface{} {
	return c.ID
}

/*// ActionBridge, to define the input and output of the previous action and next action
// PrevAction, ID of the previouse action
// NextAction, ID of the next action
// MappingParam, define the output and input paremeter name of the previous and next action
//      e.g.:
//              OUTPUT                INPUT
//              "id"                : "id"
//              "name"              : "name"
//              "address"           : "note"
//              "tag"               : "global.tag"
//              "global.address"    : "address"
type Bridge struct {
	Id           string
	Type         string
	PrevAction   string
	NextAction   string
	MappingParam tk.M
}*/

// FlowAction define the action that exist
// Type refer to cons ACTION_TYPE_*
// Action define the type of action e.g. Spark, HDFS, etc
// Server define which server for the action
// OK next step to go if the result is OK
// KO next step to go if the result is KO or NOT OK or ERROR
// Retry how many time the system will do the process if ERROR happen and then go to KO condition, default will be 3
// Wait interval in minute, default will be 1
// InputParam, will be the list of input that needed by the action
// OutputParam, will be the list of output from the action
// OutputType, will be the type of the output (text, json, *sv, xml)
// OutputPath, path of the output in the hdfs if empty then consider as stdout
type FlowAction struct {
	Id          string      `json:"_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Server      Server      `json:"server"`
	Action      interface{} `json:"action"`
	OK          []string    `json:"ok"`
	KO          []string    `json:"ko"`
	Retry       int         `json:"retry"`
	Interval    int         `json:"interval"`
	FirstAction bool        `json:"firstaction"`
	InputParam  tk.M        `json:"inputparam"`
	OutputParam tk.M        `json:"outputparam"`
	OutputType  string      `json:"outputtype"`
	OutputPath  string      `json:"outputtype"`
}

// ActionHive action for HIVE
// ScriptPath path of the hive script to run
type ActionHive struct {
	ScriptPath string   `json:"scrippath"`
	Params     []string `json:"params"`
}

// ActionHDFS action for HDFS type for server
// Operation type operation
// The HDFS connection will be provided by the server from FlowAction
// Path can be used for the current path, or old path, also can included the filename
// NewPath can be used for move, copy, etc
// Permission e.g. rwx-r--r-
type ActionHDFS struct {
	/*Operation  string `json:"operation"`
	Path       string `json:"path"`
	NewPath    string `json:"newpath"`
	Permission string
	User       string
	Group      string*/
	Command string `json:"command"`
}

// ActionSSH action for SSH type of server
// Operation type operation
// The ssh connection will be provided by the server from FlowAction
// Path can be used for the current path, or old path, also can included the filename
// NewPath can be used for move, copy, etc
// Permission e.g. rwx-r--r-
type ActionSSH struct {
	Command string `json:"command"`
}

// ActionSpark action for SPARK
// Type of the code upload e.g. Java, Scala
// Mode e.g. client, cluster
// Args for other arguments, e.g. --executor-memory 2G
// Application the spark code will be mantain as application
type ActionSpark struct {
	Type      string `json: "type"`
	Master    string `json:"master"`
	Mode      string `json:"node"`
	File      string `json:"file"`
	MainClass string `json:"mainclass"`
	Args      string `json:"args"`
	// AppName   string
	// Args      []string
}

// ActionSpark action for Application
// Type e.g. Java, Scala, GO
// Application refer to application page
// when creating the action with type = application then should follow the application screen standard
type ActionApplication struct {
	Type        string      `json: "type"`
	Application Application `json:"application"`
}

// ActionHadoopStreaming for hadoop streaming
// Mapper name
// Reducer name
// Input can be folder or file
// Output folder
// Files list of file inside the hdfs
// the server type should be hdfs
type ActionHadoopStreaming struct {
	Jar     string   `json:"jar"`
	Mapper  string   `json:"mapper"`
	Reducer string   `json:"reducer"`
	Input   string   `json:"input"`
	Output  string   `json:"output"`
	Files   []string `json:"files"`
	// Params  []string
	Params string `json:"params"`
}

// ActionKafka action for KAFKA
type ActionKafka struct {
	// other field will be different for each action type
}

/*
Doesn't need the Queestion, can be achived by using OK and KO or using decision
type ActionQuestion struct {
	Question string `json:"question"`
	Yes      string `json:"yes"`
	No       string `json:"no"`
	// other field will be different for each action type
}*/

// ActionDecision action to define the list of decision that action have
type ActionDecision struct {
	Conditions []Condition `json:"conditions"`
	// other field will be different for each action type
}

// Condition condition for the action decision
type Condition struct {
	Stat       string `json:"stat"`
	FlowAction string `json:"flowaction"`
}

type ActionStop struct {
	Message string `json:"message"`
}

type ActionShellScript struct {
	Script string `json:"script"`
}

type ActionJavaApp struct {
	Jar string `json:"jar"`
}

type ActionEmail struct {
	To      string `json:"to"`
	Cc      string `json:"cc"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// ActionFork to define the forking condition
// Type refer to cons FORK_TYPE_*
type ActionFork struct {
	Actions []FlowAction `json:"actions"`
	Type    string       `json: "type"`
}

type DataFlowProcess struct {
	orm.ModelBase
	Id          string       `json:"_id"`
	Flow        DataFlow     `json:"dataflow"`
	Steps       []FlowAction `json:"flowaction"`
	StartDate   time.Time    `json:"startdate"`
	EndDate     time.Time    `json:"enddate"`
	StartedBy   string       `json:"startedby"`
	GlobalParam tk.M         `json:"globalparam"`
}

func (c *DataFlowProcess) TableName() string {
	return "dataflowprocess"
}

func (c *DataFlowProcess) RecordID() interface{} {
	return c.Id
}

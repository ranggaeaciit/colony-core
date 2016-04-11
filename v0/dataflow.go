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
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
	tk "github.com/eaciit/toolkit"
	"time"
)

const (
	ACTION_TYPE_HIVE     = "HIVE"
	ACTION_TYPE_HDFS     = "HDFS"
	ACTION_TYPE_SPARK    = "SPARK"
	ACTION_TYPE_DECISION = "DECISION"
	ACTION_TYPE_SSH      = "SSH"
	ACTION_TYPE_KAFKA    = "KAFKA"

	FORK_TYPE_ALL       = "ALL"
	FORK_TYPE_ONE       = "ONE"
	FORK_TYPE_MANDATORY = "MANDATORY"

	SSH_OPERATION_MKDIR = "MKDIR"
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
}

// FlowAction define the action that exist
// Type refer to cons ACTION_TYPE_*
// Action define the type of action e.g. Spark, HDFS, etc
// Server define which server for the action
// OK next step to go if the result is OK
// KO next step to go if the result is KO or NOT OK or ERROR
// Retry how many time the system will do the process if ERROR happen and then go to KO condition, default will be 3
// Wait interval in minute, default will be 1
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
	FirstAction bool
	Context     []ActionContext
}

// ActionHive action for HIVE
// ScriptPath path of the hive script to run
type ActionHive struct {
	ScriptPath string
	Params     []string
}

// ActionHDFS action for HDFS type for server
// Operation type operation
// The HDFS connection will be provided by the server from FlowAction
// Path can be used for the current path, or old path, also can included the filename
// NewPath can be used for move, copy, etc
// Permission e.g. rwx-r--r-
type ActionHDFS struct {
	Operation  string `json:"operation"`
	Path       string `json:"path"`
	NewPath    string `json:"newpath"`
	Permission string
	User       string
	Group      string
}

// ActionSSH action for SSH type of server
// Operation type operation
// The ssh connection will be provided by the server from FlowAction
// Path can be used for the current path, or old path, also can included the filename
// NewPath can be used for move, copy, etc
// Permission e.g. rwx-r--r-
type ActionSSH struct {
	Operation  string `json:"operation"`
	Path       string `json:"path"`
	NewPath    string `json:"newpath"`
	Permission string
	User       string
	Group      string
}

// ActionSpark action for SPARK
// Type of the code upload e.g. Java, Scala
// Mode e.g. client, cluster
// Args for other arguments, e.g. --executor-memory 2G
// Application the spark code will be mantain as application
type ActionSpark struct {
	Type        string `json: "type"`
	Master      string
	Mode        string
	AppName     string
	Args        []string
	Application Application
}

// ActionSpark action for Application
// Type e.g. Java, Scala, GO
// Application refer to application page
// when creating the action with type = application then should follow the application screen standard
type ActionApplication struct {
	Type        string `json: "type"`
	Application Application
}

// ActionHadoopStreaming for hadoop streaming
// Mapper name
// Reducer name
// Input can be folder or file
// Output folder
// Files list of file inside the hdfs
// the server type should be hdfs
type ActionHadoopStreaming struct {
	Mapper  string
	Reducer string
	Input   string
	Output  string
	Files   []string
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
	Result     string `json:"result"`
	FlowAction string `json:"flowaction"`
}

// ActionFork to define the forking condition
// Type refer to cons FORK_TYPE_*
type ActionFork struct {
	Actions []FlowAction `json:"actions"`
	Type    string       `json: "type"`
}

type DataFlowProcess struct {
	Id          string `json:"_id"`
	Flow        DataFlow
	Steps       []FlowAction
	StartDate   time.Time
	EndDate     time.Time
	UserStarted string
}

func (c *DataFlow) TableName() string {
	return "dataflow"
}

func (c *DataFlow) RecordID() interface{} {
	return c.ID
}

type ActionContext struct {
	Keys  interface{}
	Infos []toolkit.M
}

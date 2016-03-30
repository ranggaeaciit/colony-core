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
)

const (
	TYPE_HIVE     = "HIVE"
	TYPE_HDFS     = "HDFS"
	TYPE_SPARK    = "SPARK"
	TYPE_DECISION = "DECISION"
)

type DataFlow struct {
	Id          string       `json:"_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedDate time.Time    `json:"createddate"`
	CreatedBy   string       `json:"createdby"`
	Actions     []FlowAction `json:"actions"`
}

type FlowAction struct {
	Id          string        `json:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        string        `json:"type"`
	Server      Server        `json:"server"`
	Action      interface{}   `json:"action"`
	NextActions []interface{} `json:"nextactions"`
}

type ActionHive struct {
	Id string `json:"_id"`
	// other field will be different for each action type
}

type ActionHDFS struct {
	Id string `json:"_id"`
	// other field will be different for each action type
}

type ActionSpark struct {
	Id string `json:"_id"`
	// other field will be different for each action type
}

type ActionDecision struct {
	Id string `json:"_id"`
	// other field will be different for each action type
}

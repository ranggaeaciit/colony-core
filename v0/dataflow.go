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
	Id          string      `json:"_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Server      Server      `json:"server"`
	Action      interface{} `json:"action"`
	OK          interface{} `json:"ok"`
	KO          interface{} `json:"ko"`
}

type ActionHive struct {
	// other field will be different for each action type
}

type ActionHDFS struct {
	// other field will be different for each action type
}

type ActionSpark struct {
	// other field will be different for each action type
}

type ActionQuestion struct {
	Question string `json:"question"`
	Yes      string `json:"yes"`
	No       string `json:"no"`
	// other field will be different for each action type
}

type ActionDecision struct {
	Conditions []Condition `json:"conditions"`
	// other field will be different for each action type
}

type Condition struct {
	Result     string `json:"result"`
	FlowAction string `json:"flowaction"`
}

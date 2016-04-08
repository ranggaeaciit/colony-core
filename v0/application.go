package colonycore

import (
	"github.com/eaciit/orm/v1"
	"strings"
)

type Application struct {
	orm.ModelBase
	ID            string `json:"_id",bson:"_id"`
	IsInternalApp bool
	AppsName      string
	Port          string // for web type
	Type          string
	ZipName       string
	Enable        bool
	DeployedTo    []string
	Command       interface{}
	Variable      interface{}
}

func (a *Application) TableName() string {
	return "applications"
}

func (a *Application) RecordID() interface{} {
	return a.ID
}

type TreeSource struct {
	ID             int           `json:"_id",bson:"_id"`
	Text           string        `json:"text",bson:"text"`
	Expanded       bool          `json:"expanded",bson:"expanded"`
	SpriteCssClass string        `json:"spriteCssClass",bson:"spriteCssClass"`
	Items          []*TreeSource `json:"items",bson:"items"`
}

type TreeSourceModel struct {
	Text      string             `json:"text"`
	Type      string             `json:"type"`
	Expanded  bool               `json:"expanded"`
	Iconclass string             `json:"iconclass"`
	Ext       string             `json:"ext"`
	Path      string             `json:"path"`
	Items     []*TreeSourceModel `json:"items"`
}

func (a *Application) GetCommand(cmdName string) (bool, string) {
	cmdString := ""

	for _, raw := range a.Command.([]interface{}) {
		each := raw.(map[string]interface{})
		if each["key"].(string) == cmdName {
			cmdString = each["value"].(string)
			break
		}
	}

	for _, raw := range a.Variable.([]interface{}) {
		each := raw.(map[string]interface{})
		if strings.Contains(cmdString, each["key"].(string)) {
			cmdString = strings.Replace(cmdString, each["key"].(string), each["value"].(string), -1)
		}
	}

	if cmdString == "" {
		return false, ""
	}

	return true, cmdString
}

const (
	App_Command_Start         = "start"
	App_Command_Stop          = "stop"
	App_Command_RunningStatus = "running_status"
	App_Command_DeployStatus  = "deploy_status"
	App_Variable_BinaryName   = "BINARY_NAME"
)

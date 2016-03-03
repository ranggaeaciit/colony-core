package colonycore

import (
	"github.com/eaciit/orm/v1" 
)

type Application struct {
	orm.ModelBase
	ID         string `json:"_id",bson:"_id"`
	AppsName   string
	Port       string
	Type       string
	ZipName    string
	Enable     bool
	DeployedTo []string
	Command    interface{}
	Variable   interface{}

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

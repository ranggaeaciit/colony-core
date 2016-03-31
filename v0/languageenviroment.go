package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type LanguageEnvironmentPayload struct {
	ServerId string   `json:"ServerId"`
	Lang     []string `json:"Lang"`
}

type LanguageEnviroment struct {
	orm.ModelBase
	Language  string       `json:"language", bson:"language"`
	Commands  []*Commands  `json:commands, bson:"commands"`
	Installer []*Installer `json:installer, bson:"installer"`
}

type Commands struct {
	Name string `json:"name", bson:"name"`
	Cmd  string `json:"cmd", bson:"cmd"`
}

type Installer struct {
	OS              string `json:"os", bson:"os"`
	InstallerFile   string `json:"installerFile", bson:"installerFile"`
	InstallerSource string `json:"installerSource", bson:"installerSource"`
}

func (le *LanguageEnviroment) TableName() string {
	return "Language"
}

func (le *LanguageEnviroment) RecordID() interface{} {
	return le.Language
}

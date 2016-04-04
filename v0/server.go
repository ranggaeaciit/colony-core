package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Server struct {
	orm.ModelBase
	ID         string       `json:"_id", bson:"_id"`
	OS         string       `json:"os", bson:"os"`
	AppPath    string       `json:"appPath",bson:"appPath"`
	DataPath   string       `json:"dataPath",bson:"dataPath"`
	Host       string       `json:"host", bson:"host"`
	ServerType string       `json:"serverType", bson:"serverType"`
	SSHType    string       `json:"sshtype", bson:"sshtype"`
	SSHFile    string       `json:"sshfile", bson:"sshfile"`
	SSHUser    string       `json:"sshuser", bson:"sshuser"`
	SSHPass    string       `json:"sshpass", bson:"sshpass"`
	CmdExtract string       `json:"cmdextract", bson:"cmdextract"`
	CmdNewFile string       `json:"cmdnewfile", bson:"cmdnewfile"`
	CmdCopy    string       `json:"cmdcopy", bson:"cmdcopy"`
	CmdMkDir   string       `json:"cmdmkdir", bson:"cmdmkdir"`
	HostAlias  []*HostAlias `json:"hostAlias", bson:"hostAlias"`
}

type HostAlias struct {
	IP       string `json:"ip", bson:"ip"`
	HostName string `json:"hostName", bson:"hostName"`
}

func (s *Server) TableName() string {
	return "servers"
}

func (s *Server) RecordID() interface{} {
	return s.ID
}

package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Server struct {
	orm.ModelBase
	ID           string `json:"_id", bson:"_id"`
	Type         string `json:"type", bson:"type"`
	Folder       string `json:"folder", bson:"folder"`
	OS           string `json:"os", bson:"os"`
	Enable       bool   `json:"enable", bson:"enable"`
	SSHType      string `json:"sshtype", bson:"sshtype"`
	SSHFile      string `json:"sshfile", bson:"sshfile"`
	SSHUser      string `json:"sshuser", bson:"sshuser"`
	SSHPass      string `json:"sshpass", bson:"sshpass"`
	CmdExtract   string `json:"extract", bson:"extract"`
	CmdNewFile   string `json:"newfile", bson:"newfile"`
	CmdCopy      string `json:"copy", bson:"copy"`
	CmdDirectory string `json:"dir", bson:"dir"`
}

func (s *Server) TableName() string {
	return "servers"
}

func (s *Server) RecordID() interface{} {
	return s.ID
}

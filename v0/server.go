package colonycore

import (
	"errors"
	"fmt"
	"github.com/eaciit/hdc/hdfs"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/sshclient"
	"golang.org/x/crypto/ssh"
	"strings"
)

type Server struct {
	orm.ModelBase
	ID   string `json:"_id", bson:"_id"`
	OS   string `json:"os", bson:"os"`
	Host string `json:"host", bson:"host"`

	AppPath  string `json:"appPath", bson:"appPath"`
	DataPath string `json:"dataPath", bson:"dataPath"`

	ServiceSSH  *ServiceSSH  `json:"serviceSSH", json:"serviceSSH"`
	ServiceHDFS *ServiceHDFS `json:"serviceHDFS", json:"serviceHDFS"`

	CmdExtract string `json:"cmdextract", bson:"cmdextract"`
	CmdNewFile string `json:"cmdnewfile", bson:"cmdnewfile"`
	CmdCopy    string `json:"cmdcopy", bson:"cmdcopy"`
	CmdMkDir   string `json:"cmdmkdir", bson:"cmdmkdir"`

	InstalledLang []*InstalledLang `json:"installedLang", bson:"installedLang"`

	/* DEPRECATED */

	// ServerType string `json:"serverType", bson:"serverType"`
	// SSHType   string       `json:"sshtype", bson:"sshtype"`
	// SSHFile   string       `json:"sshfile", bson:"sshfile"`
	// SSHUser   string       `json:"sshuser", bson:"sshuser"`
	// SSHPass   string       `json:"sshpass", bson:"sshpass"`
	// HostAlias []*HostAlias `json:"hostAlias", bson:"hostAlias"`
}

type ServerByType struct {
	ServerType  string `json:"serverType"`
	ServerAlias string `json:"serverAlias"`
	Server
}

type ServiceSSH struct {
	Type string `json:"type", bson:"type"`
	File string `json:"file", bson:"file"`
	Host string `json:"host", bson:"host"`
	User string `json:"user", bson:"user"`
	Pass string `json:"pass", bson:"pass"`
}

type ServiceHDFS struct {
	Host      string       `json:"host", bson:"host"`
	User      string       `json:"user", bson:"user"`
	Pass      string       `json:"pass", bson:"pass"`
	HostAlias []*HostAlias `json:"hostAlias", bson:"hostAlias"`
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

func (s *Server) Connect() (sshclient.SshSetting, *ssh.Client, error) {
	client := sshclient.SshSetting{}

	if s.ServiceSSH == nil {
		return client, nil, errors.New("ssh information is not setup yet")
	}

	ssh := s.ServiceSSH
	client.SSHHost = ssh.Host

	if ssh.Type == "File" {
		client.SSHAuthType = sshclient.SSHAuthType_Certificate
		client.SSHKeyLocation = ssh.File
	} else {
		client.SSHAuthType = sshclient.SSHAuthType_Password
		client.SSHUser = ssh.User
		client.SSHPassword = ssh.Pass
	}

	theClient, err := client.Connect()

	return client, theClient, err
}

func (s *Server) IsCommandExists(cmd string) (bool, string, error) {
	setting, _, err := s.Connect()
	if err != nil {
		return false, "", err
	}

	res, err := setting.RunCommandSshAsMap(fmt.Sprintf("which %s", cmd))
	if err != nil {
		return false, "", err
	}

	output := strings.TrimSpace(res[0].Output)
	if output == "" {
		return false, output, errors.New("command not found")
	}

	if resOutput, err := setting.RunCommandSshAsMap(cmd); err == nil {
		output = resOutput[0].Output
	}

	return true, output, nil
}

func (s *Server) Ping(serverType string) (bool, error) {
	if serverType == "node" {
		if _, _, err := s.Connect(); err != nil {
			return false, err
		}
	} else {
		hdfsConfig := hdfs.NewHdfsConfig(s.ServiceHDFS.Host, s.ServiceHDFS.User)
		hdfsConfig.Password = s.ServiceHDFS.Pass

		hadeepes, err := hdfs.NewWebHdfs(hdfsConfig)
		if err != nil {
			return false, err
		}

		if _, err := hadeepes.List("/"); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *Server) DetectInstalledLang() {
	cursorLang, err := Find(new(LanguageEnviroment), nil)
	if err == nil {
		defer cursorLang.Close()

		langs := []*LanguageEnviroment{}
		err = cursorLang.Fetch(&langs, 0, false)
		if err == nil {
			for _, lang := range langs {
				cmd := lang.Commands.GetString("version")
				isExist, output, _ := s.IsCommandExists(cmd)

				l := new(InstalledLang)
				l.IsInstalled = isExist
				l.Lang = lang.Language
				l.Version = output

				s.InstalledLang = append(s.InstalledLang, l)
			}
		}
	}
}

func (s *Server) GetByType() ([]*ServerByType, error) {
	cursor, err := Find(new(Server), nil)
	if err != nil {
		return nil, err
	}

	data := []Server{}
	err = cursor.Fetch(&data, 0, false)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	servers := []*ServerByType{}
	serverTypes := []string{"hdfs", "node"}

	for _, each := range data {
		for _, serverType := range serverTypes {
			if serverType == serverTypes[0] && each.ServiceHDFS.Host == "" {
				continue
			}

			if serverType == serverTypes[1] && each.ServiceSSH.Host == "" {
				continue
			}

			server := new(ServerByType)
			server.Server = each
			server.ServerType = serverType
			server.ServerAlias = fmt.Sprintf("%s (%s)", server.ID, serverType)

			servers = append(servers, server)
		}
	}

	return servers, nil
}

func (s *Server) DetectService() {

}

type ServerLanguage struct {
	ServerID   string
	ServerOS   string
	ServerHost string
	Languages  []*InstalledLang
}

type InstalledLang struct {
	Lang        string
	Version     string
	IsInstalled bool
}

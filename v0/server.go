package colonycore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/eaciit/dbox"
	"github.com/eaciit/hdc/hdfs"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/sshclient"
	"github.com/eaciit/toolkit"
	"golang.org/x/crypto/ssh"
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

func (s *Server) Ping(serviceType string) (bool, error) {
	if serviceType == "ssh" {
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

func (s *Server) GetServerSSH() ([]Server, error) {
	cursorServer, err := Find(new(Server), nil)
	if err != nil {
		return nil, err
	}
	dataServer := []Server{}
	err = cursorServer.Fetch(&dataServer, 0, false)
	if err != nil {
		return nil, err
	}
	defer cursorServer.Close()

	dataServerSSH := []Server{}
	for _, each := range dataServer {
		if each.ServiceSSH == nil {
			continue
		}
		if each.ServiceSSH.Host == "" || each.ServiceSSH.User == "" {
			continue
		}
		dataServerSSH = append(dataServerSSH, each)
	}

	return dataServerSSH, nil
}

func (s *Server) DetectService() {

}

func (s *Server) IsAccessValid(what string) bool {
	if what == "node" {
		if s.ServiceSSH.Type == "Credentials" {
			return (s.ServiceSSH.Host != "") && (s.ServiceSSH.User != "")
		} else {
			return (s.ServiceSSH.Host != "") && (s.ServiceSSH.File != "")
		}
	} else if what == "hdfs" {
		return (s.ServiceHDFS.Host != "") && (s.ServiceHDFS.User != "")
	}

	return false
}

func (s *Server) InstallColonyOnLinux(log *toolkit.LogEngine) error {
	oldServer := new(Server)

	log.AddLog(fmt.Sprintf("Find server ID: %s", s.ID), "INFO")
	cursor, err := Find(new(Server), dbox.Eq("_id", s.ID))
	if err != nil {
		log.AddLog(err.Error(), "ERROR")
		return err
	}
	oldServerAll := []Server{}
	err = cursor.Fetch(&oldServerAll, 0, false)
	if err == nil {
		defer cursor.Close()
		if len(oldServerAll) > 0 {
			oldServer = &oldServerAll[0]
		}
	}

	log.AddLog(fmt.Sprintf("SSH Connect %v", s.ServiceSSH), "INFO")
	sshSetting, client, err := s.Connect()
	if err != nil {
		log.AddLog(err.Error(), "ERROR")
		return err
	}
	defer client.Close()

	setEnvPath := func() error {
		cmd1 := `sed -i '/export EC_APP_PATH/d' ~/.bashrc`
		log.AddLog(cmd1, "INFO")
		sshSetting.GetOutputCommandSsh(cmd1)

		cmd2 := `sed -i '/export EC_DATA_PATH/d' ~/.bashrc`
		log.AddLog(cmd2, "INFO")
		sshSetting.GetOutputCommandSsh(cmd2)

		cmd3 := "echo 'export EC_APP_PATH=" + s.AppPath + "' >> ~/.bashrc"
		log.AddLog(cmd3, "INFO")
		sshSetting.GetOutputCommandSsh(cmd3)

		cmd4 := "echo 'export EC_DATA_PATH=" + s.DataPath + "' >> ~/.bashrc"
		log.AddLog(cmd4, "INFO")
		sshSetting.GetOutputCommandSsh(cmd4)

		return nil
	}

	if oldServer.AppPath == "" || oldServer.DataPath == "" {
		cmdTestUnzip := "unzip"
		log.AddLog(cmdTestUnzip, "INFO")
		unzipRes, err := sshSetting.GetOutputCommandSsh(cmdTestUnzip)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}
		if strings.Contains(strings.ToLower(unzipRes), "not found") {
			log.AddLog("Need to install `unzip` on the server", "ERROR")
			return err
		}

		cmdRmAppPath := fmt.Sprintf("rm -rf %s", s.AppPath)
		log.AddLog(cmdRmAppPath, "INFO")
		sshSetting.GetOutputCommandSsh(cmdRmAppPath)

		cmdMkdirAppPath := fmt.Sprintf(`mkdir -p "%s"`, s.AppPath)
		log.AddLog(cmdMkdirAppPath, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(cmdMkdirAppPath)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		cmdRmDataPath := fmt.Sprintf("rm -rf %s", s.DataPath)
		log.AddLog(cmdRmDataPath, "INFO")
		sshSetting.GetOutputCommandSsh(cmdRmDataPath)

		cmdMkdirDataPath := fmt.Sprintf(`mkdir -p "%s"`, s.DataPath)
		log.AddLog(cmdMkdirDataPath, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(cmdMkdirDataPath)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		appDistSrc := filepath.Join(os.Getenv("EC_DATA_PATH"), "dist", "app-root.zip")
		err = sshSetting.SshCopyByPath(appDistSrc, s.AppPath)
		log.AddLog(fmt.Sprintf("scp from %s to %s", appDistSrc, s.AppPath), "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		appDistSrcDest := filepath.Join(s.AppPath, "app-root.zip")
		appDistSrcDest = strings.Replace(appDistSrcDest, "\\", "/", -1)
		unzipAppCmd := strings.Replace(strings.Replace(s.CmdExtract, "%1", appDistSrcDest, -1), "%2", s.AppPath, -1)
		log.AddLog(unzipAppCmd, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(unzipAppCmd)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		rmTempAppPath := fmt.Sprintf("rm -rf %s", appDistSrcDest)
		_, err = sshSetting.GetOutputCommandSsh(rmTempAppPath)
		log.AddLog(rmTempAppPath, "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		dataDistSrc := filepath.Join(os.Getenv("EC_DATA_PATH"), "dist", "data-root.zip")
		err = sshSetting.SshCopyByPath(dataDistSrc, s.DataPath)
		log.AddLog(fmt.Sprintf("scp from %s to %s", dataDistSrc, s.DataPath), "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		dataDistSrcDest := filepath.Join(s.DataPath, "data-root.zip")
		dataDistSrcDest = strings.Replace(dataDistSrcDest, "\\", "/", -1)
		unzipDataCmd := strings.Replace(strings.Replace(s.CmdExtract, "%1", dataDistSrcDest, -1), "%2", s.DataPath, -1)
		log.AddLog(unzipDataCmd, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(unzipDataCmd)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		rmTempDataPath := fmt.Sprintf("rm -rf %s", dataDistSrcDest)
		_, err = sshSetting.GetOutputCommandSsh(rmTempDataPath)
		log.AddLog(rmTempDataPath, "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		osArchCmd := "uname -m"
		log.AddLog(osArchCmd, "INFO")
		osArchRes, err := sshSetting.GetOutputCommandSsh(osArchCmd)
		osArchRes = strings.TrimSpace(osArchRes)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		for _, each := range []string{"sedotand", "sedotans", "sedotanw"} {
			src := filepath.Join(os.Getenv("EC_APP_PATH"), "cli", "dist", fmt.Sprintf("linux_%s", osArchRes), each)
			dst := filepath.Join(s.AppPath, "cli", each)

			rmSedotanCmd := fmt.Sprintf("rm -rf %s", dst)
			log.AddLog(rmSedotanCmd, "INFO")
			_, err := sshSetting.GetOutputCommandSsh(rmSedotanCmd)
			if err != nil {
				log.AddLog(err.Error(), "ERROR")
				return err
			}

			log.AddLog(fmt.Sprintf("scp %s to %s", src, dst), "INFO")
			err = sshSetting.SshCopyByPath(src, dst)
			if err != nil {
				log.AddLog(err.Error(), "ERROR")
				return err
			}

			chmodCliCmd := fmt.Sprintf("chmod 755 %s", dst)
			log.AddLog(chmodCliCmd, "INFO")
			_, err = sshSetting.GetOutputCommandSsh(chmodCliCmd)
			if err != nil {
				log.AddLog(err.Error(), "ERROR")
				return err
			}
		}

		checkPathCmd := fmt.Sprintf("ls %s", s.AppPath)
		isPathCreated, err := sshSetting.GetOutputCommandSsh(checkPathCmd)
		log.AddLog(checkPathCmd, "INFO")
		if err != nil || strings.TrimSpace(isPathCreated) == "" {
			errString := fmt.Sprintf("Invalid path. %s", err.Error())
			log.AddLog(errString, "ERROR")
			return errors.New(errString)
		}

		if err := setEnvPath(); err != nil {
			return err
		}
	} else if oldServer.AppPath != s.AppPath {
		moveDir := fmt.Sprintf(`mv %s %s`, oldServer.AppPath, s.AppPath)
		log.AddLog(moveDir, "INFO")
		_, err := sshSetting.GetOutputCommandSsh(moveDir)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		if err := setEnvPath(); err != nil {
			return err
		}
	} else if oldServer.DataPath != s.DataPath {
		moveDir := fmt.Sprintf(`mv %s %s`, oldServer.DataPath, s.DataPath)
		log.AddLog(moveDir, "INFO")
		_, err := sshSetting.GetOutputCommandSsh(moveDir)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		if err := setEnvPath(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) InstallColonyOnWindows(log *toolkit.LogEngine) error {
	oldServer := new(Server)

	log.AddLog(fmt.Sprintf("Find server ID: %s", s.ID), "INFO")
	cursor, err := Find(new(Server), dbox.Eq("_id", s.ID))
	if err != nil {
		log.AddLog(err.Error(), "ERROR")
		return err
	}
	oldServerAll := []Server{}
	err = cursor.Fetch(&oldServerAll, 0, false)
	if err == nil {
		defer cursor.Close()
		if len(oldServerAll) > 0 {
			oldServer = &oldServerAll[0]
		}
	}

	log.AddLog(fmt.Sprintf("SSH Connect %v", s.ServiceSSH), "INFO")
	sshSetting, client, err := s.Connect()
	if err != nil {
		log.AddLog(err.Error(), "ERROR")
		return err
	}
	defer client.Close()

	setEnvPath := func() error {
		cmd1 := fmt.Sprintf(`setx EC_APP_PATH "%s"`, s.AppPath)
		log.AddLog(cmd1, "INFO")
		sshSetting.GetOutputCommandSsh(cmd1)

		cmd2 := fmt.Sprintf(`setx EC_DATA_PATH "%s"`, s.DataPath)
		log.AddLog(cmd2, "INFO")
		sshSetting.GetOutputCommandSsh(cmd2)

		return nil
	}

	if oldServer.AppPath == "" || oldServer.DataPath == "" {
		cmdTestUnzip := "unzip"
		log.AddLog(cmdTestUnzip, "INFO")
		unzipRes, err := sshSetting.GetOutputCommandSsh(cmdTestUnzip)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}
		if strings.Contains(strings.ToLower(unzipRes), "not recognized") {
			log.AddLog("Need to install `unzip` on the server", "ERROR")
			return err
		}

		cmdRmAppPath := fmt.Sprintf("rmdir /S /Q %s", s.AppPath)
		log.AddLog(cmdRmAppPath, "INFO")
		sshSetting.GetOutputCommandSsh(cmdRmAppPath)

		cmdMkdirAppPath := fmt.Sprintf(`mkdir "%s"`, s.AppPath)
		log.AddLog(cmdMkdirAppPath, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(cmdMkdirAppPath)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		cmdRmDataPath := fmt.Sprintf("rmdir /S /Q %s", s.DataPath)
		log.AddLog(cmdRmDataPath, "INFO")
		sshSetting.GetOutputCommandSsh(cmdRmDataPath)

		cmdMkdirDataPath := fmt.Sprintf(`mkdir "%s"`, s.DataPath)
		log.AddLog(cmdMkdirDataPath, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(cmdMkdirDataPath)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		appDistSrc := filepath.Join(os.Getenv("EC_DATA_PATH"), "dist", "app-root.zip")
		err = sshSetting.SshCopyByPath(appDistSrc, s.AppPath)
		log.AddLog(fmt.Sprintf("scp from %s to %s", appDistSrc, s.AppPath), "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		appDistSrcDest := filepath.Join(s.AppPath, "app-root.zip")
		appDistSrcDest = strings.Replace(appDistSrcDest, "\\", "/", -1)
		unzipAppCmd := strings.Replace(strings.Replace(s.CmdExtract, "%1", appDistSrcDest, -1), "%2", s.AppPath, -1)
		log.AddLog(unzipAppCmd, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(unzipAppCmd)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		rmTempAppPath := fmt.Sprintf("rmdir /S /Q %s", appDistSrcDest)
		_, err = sshSetting.GetOutputCommandSsh(rmTempAppPath)
		log.AddLog(rmTempAppPath, "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		dataDistSrc := filepath.Join(os.Getenv("EC_DATA_PATH"), "dist", "data-root.zip")
		err = sshSetting.SshCopyByPath(dataDistSrc, s.DataPath)
		log.AddLog(fmt.Sprintf("scp from %s to %s", dataDistSrc, s.DataPath), "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		dataDistSrcDest := filepath.Join(s.DataPath, "data-root.zip")
		dataDistSrcDest = strings.Replace(dataDistSrcDest, "\\", "/", -1)
		unzipDataCmd := strings.Replace(strings.Replace(s.CmdExtract, "%1", dataDistSrcDest, -1), "%2", s.DataPath, -1)
		log.AddLog(unzipDataCmd, "INFO")
		_, err = sshSetting.GetOutputCommandSsh(unzipDataCmd)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		rmTempDataPath := fmt.Sprintf("rmdir /S /Q %s", dataDistSrcDest)
		_, err = sshSetting.GetOutputCommandSsh(rmTempDataPath)
		log.AddLog(rmTempDataPath, "INFO")
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		osArchCmd := "echo %PROCESSOR_ARCHITECTURE%"
		log.AddLog(osArchCmd, "INFO")
		osArchRes, err := sshSetting.GetOutputCommandSsh(osArchCmd)
		osArchRes = strings.TrimSpace(osArchRes)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}
		if osArchRes != "x86" {
			osArchRes = "x86_64"
		}

		for _, each := range []string{"sedotand", "sedotans", "sedotanw"} {
			src := filepath.Join(os.Getenv("EC_APP_PATH"), "cli", "dist", fmt.Sprintf("windows_%s", osArchRes), each)
			dst := filepath.Join(s.AppPath, "cli", each)

			rmSedotanCmd := fmt.Sprintf("rmdir /S /Q %s", dst)
			log.AddLog(rmSedotanCmd, "INFO")
			_, err := sshSetting.GetOutputCommandSsh(rmSedotanCmd)
			if err != nil {
				log.AddLog(err.Error(), "ERROR")
				return err
			}

			log.AddLog(fmt.Sprintf("scp %s to %s", src, dst), "INFO")
			err = sshSetting.SshCopyByPath(src, dst)
			if err != nil {
				log.AddLog(err.Error(), "ERROR")
				return err
			}

			chmodCliCmd := fmt.Sprintf("cacls %s /g everyone:f 755", dst)
			log.AddLog(chmodCliCmd, "INFO")
			_, err = sshSetting.GetOutputCommandSsh(chmodCliCmd)
			if err != nil {
				log.AddLog(err.Error(), "ERROR")
				return err
			}
		}

		checkPathCmd := fmt.Sprintf("dir %s", s.AppPath)
		isPathCreated, err := sshSetting.GetOutputCommandSsh(checkPathCmd)
		log.AddLog(checkPathCmd, "INFO")
		if err != nil || strings.TrimSpace(isPathCreated) == "" {
			errString := fmt.Sprintf("Invalid path. %s", err.Error())
			log.AddLog(errString, "ERROR")
			return errors.New(errString)
		}

		if err := setEnvPath(); err != nil {
			return err
		}
	} else if oldServer.AppPath != s.AppPath {
		moveDir := fmt.Sprintf(`move %s %s`, oldServer.AppPath, s.AppPath)
		log.AddLog(moveDir, "INFO")
		_, err := sshSetting.GetOutputCommandSsh(moveDir)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		if err := setEnvPath(); err != nil {
			return err
		}
	} else if oldServer.DataPath != s.DataPath {
		moveDir := fmt.Sprintf(`move %s %s`, oldServer.DataPath, s.DataPath)
		log.AddLog(moveDir, "INFO")
		_, err := sshSetting.GetOutputCommandSsh(moveDir)
		if err != nil {
			log.AddLog(err.Error(), "ERROR")
			return err
		}

		if err := setEnvPath(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ToggleSedotanService(op string, id string) (bool, error) {
	sshSetting, client, err := s.Connect()
	if err != nil {
		return false, err
	}
	defer client.Close()

	data := new(Server)
	cursor, err := Find(new(Server), dbox.Eq("_id", id))
	if err != nil {
		return false, err
	}
	dataAll := []Server{}
	err = cursor.Fetch(&dataAll, 0, false)
	if err != nil {
		return false, err
	}
	defer cursor.Close()

	if len(dataAll) == 0 {
		return false, errors.New("Server not found")
	}

	data = &dataAll[0]

	pgrepSedotanCmd, err := sshSetting.GetOutputCommandSsh("pgrep sedotand")
	if err != nil {
		// do something
	}
	isOn := false
	pid := strings.TrimSpace(pgrepSedotanCmd)
	if pid != "" {
		isOn = true
	}

	if strings.Contains(op, "stat") {
		return isOn, nil
	}

	if strings.Contains(op, "stop") {
		if pid != "" {
			killProcessCmd := fmt.Sprintf("kill -9 %s", pid)
			_, err = sshSetting.GetOutputCommandSsh(killProcessCmd)
			if err != nil {
				// do something
			}

			if !strings.Contains(op, "start") && err != nil {
				return isOn, err
			}
		}
	}

	if strings.Contains(op, "start") {
		sedotanConfigArg := fmt.Sprintf(`-config="%s"`, filepath.Join(data.AppPath, "config", "webgrabbers.json"))
		sedotanLogArg := fmt.Sprintf(`-logpath="%s"`, filepath.Join(data.DataPath, "daemon"))
		runSedotanCmd := fmt.Sprintf("cd %s && ./sedotand %s %s", filepath.Join(data.AppPath, "cli"), sedotanConfigArg, sedotanLogArg)

		cRunCommand := make(chan string, 1)

		go func() {
			res, err := sshSetting.RunCommandSsh(runSedotanCmd)
			fmt.Println("cmd    ->", runSedotanCmd, "output ->", res)

			if err != nil {
				cRunCommand <- err.Error()
			} else {
				cRunCommand <- ""
			}
		}()

		errorMessage := ""
		select {
		case receiveRunCommandOutput := <-cRunCommand:
			errorMessage = receiveRunCommandOutput
		case <-time.After(time.Second * time.Duration(5)):
			errorMessage = ""
		}

		if strings.TrimSpace(errorMessage) != "" {
			return false, errors.New(errorMessage)
		}
	}

	return isOn, nil
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

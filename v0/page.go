package colonycore

import (
	// "errors"
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	// "github.com/eaciit/sshclient"
	"github.com/eaciit/toolkit"
	// "golang.org/x/crypto/ssh"
	"encoding/json"
	"os"
	"path/filepath"
	// "strings"
	"fmt"
	"io/ioutil"
	// "time"
)

type Page struct {
	orm.ModelBase
	ID    string `json:"_id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (p *Page) TableName() string {
	return "pages"
}

func (p *Page) RecordID() interface{} {
	return p.ID
}

func (p *Page) GetPages(search string) []Page {
	var query *dbox.Filter

	if search != "" {
		query = dbox.Contains("_id", search)
	}

	data := []Page{}
	cursor, err := Find(new(Page), query)
	if err != nil {
		return data
	}
	if err := cursor.Fetch(&data, 0, false); err != nil {
		return data
	}

	defer cursor.Close()
	return data
}

type PageDetail struct {
	orm.ModelBase
	ID          string        `json:"_id"`
	DataSources []string      `json:"dataSources"`
	Widgets     []*WidgetPage `json:"widgets"`
	ParentMenu  string        `json:"parentMenu"`
	Title       string        `json:"title"`
	URL         string        `json:"url"`
	ThemeColor  string        `json:"themeColor"`
	StyleSheet  string        `json:"styleSheet"`
}

type WidgetPage struct {
	ID           string    `json:"_id"`
	WidgetID     string    `json:"widgetId"`
	Title        string    `json:"title"`
	PositionType string    `json:"position"`
	X            int       `json:"x"`
	Y            int       `json:"y"`
	Height       int       `json:"height"`
	Width        int       `json:"width"`
	DataSources  toolkit.M `json:"dataSources"`
	Config       toolkit.M `json:"config"`
}

// func RandomIDWithPrefix(prefix string) string {
// 	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
// 	return toolkit.Sprintf("%s%d", prefix, timestamp)
// }

// func (mp *MapPage) Get(search string) ([]MapPage, error) {
// 	var query *dbox.Filter

// 	if search != "" {
// 		query = dbox.Contains("_id", search)
// 	}

// 	data := []MapPage{}
// 	cursor, err := Find(new(MapPage), query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := cursor.Fetch(&data, 0, false); err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close()
// 	return data, nil
// }

// func (mp *MapPage) GetById() error {
// 	if err := Get(mp, mp.ID); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (mp *MapPage) Save(payload toolkit.M) error {
// 	mp.ID = payload.Get("_id", "").(string)
// 	mp.Title = payload.Get("title", "").(string)
// 	mp.URL = payload.Get("url", "").(string)
// 	mp.FileName = mp.ID + ".json"
// 	page := new(Page)
// 	page.ID = mp.ID
// 	page.Title = mp.Title
// 	page.URL = mp.URL

// 	// page.ParentMenu = payload.Get("parentMenu", "").(string)
// 	// page.URL = payload.Get("url", "").(string)
// 	if err := Save(mp); err != nil {
// 		return err
// 	}
// 	if err := page.Save(payload, false, ""); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (mp *MapPage) Delete(pagePath string) error {
// 	jsonPath := filepath.Join(pagePath, mp.ID+".json")
// 	if err := Delete(mp); err != nil {
// 		return err
// 	}
// 	if err := os.Remove(jsonPath); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (pd *PageDetail) TableName() string {
	basePath := os.Getenv("EC_APP_PATH")
	fileName := fmt.Sprintf("%s.json", pd.ID)
	return filepath.Join(basePath, "config", "pages", fileName)
}

func (pd *PageDetail) Get() (*PageDetail, error) {
	bytes, err := ioutil.ReadFile(pd.TableName())
	if err != nil {
		return nil, err
	}

	res := new(PageDetail)
	if err := json.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (pd *PageDetail) RecordID() interface{} {
	return pd.ID
}

func (pd *PageDetail) Save() error {
	p := new(Page)
	p.ID = pd.ID
	p.Title = pd.Title
	p.URL = pd.URL
	if err := Save(p); err != nil {
		return err
	}

	data, err := json.Marshal(pd)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(pd.TableName(), data, 0666); err != nil {
		return err
	}

	return nil
}

func (pd *PageDetail) Remove() error {
	p := new(Page)
	p.ID = pd.ID
	if err := Delete(p); err != nil {
		return err
	}

	if err := os.Remove(pd.TableName()); err != nil {
		return err
	}

	return nil
}

// func (p *Page) Get(search string) ([]Page, error) {
// 	var query *dbox.Filter

// 	if search != "" {
// 		query = dbox.Contains("_id", search)
// 	}

// 	data := []Page{}
// 	cursor, err := Find(new(Page), query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := cursor.Fetch(&data, 0, false); err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close()
// 	return data, nil
// }

// func (p *Page) GetById() error {
// 	if err := Get(p, p.ID); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *Page) SaveNewWidget(payload toolkit.M, widgetPath string) (*WidgetPage, error) {
// 	wp := new(WidgetPage)
// 	wp.ID = payload.Get("widgetPageId", "").(string)
// 	wp.WidgetID = payload.Get("widgetId", "").(string)
// 	widget := new(Widget)
// 	widget.ID = wp.WidgetID
// 	widget.GetById()
// 	wp.ConfigDefault = widget.Config

// 	extractDest := filepath.Join(widgetPath, widget.ID)
// 	path, err := GetWidgetPath(extractDest)
// 	if path == "" {
// 		return nil, errors.New("directory doesn't contains index.html")
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	getConfigFile := filepath.Join(path, "config-widget.json")
// 	result, err := GetJsonFile(getConfigFile)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if result != nil { /*going to select dataSources field only from config-widget.json*/
// 		for _, value := range result[0].Get("dataSources").([]interface{}) {
// 			wp.DataSources = append(wp.DataSources, value.(map[string]interface{}))
// 		}
// 	}
// 	return wp, nil
// }

// func (p *Page) Save(payload toolkit.M, isDesigner bool, widgetPath string) error {
// 	if isDesigner {
// 		p.ID = payload.Get("_id", "").(string)
// 		if err := p.GetById(); err != nil {
// 			return err
// 		}
// 		var datasources []string
// 		wpArray := p.Widget
// 		if payload.Get("dataSourceId") != nil { /*save configuration page*/
// 			for _, ds := range payload.Get("dataSourceId", "").([]interface{}) {
// 				datasources = append(datasources, ds.(string))
// 			}
// 			p.DataSources = datasources
// 		}

// 		var mode string
// 		if payload.Get("mode") != nil {
// 			mode = payload.Get("mode", "").(string)
// 		}
// 		if mode == "new widget" { /*save new widget*/
// 			newWidget, err := p.SaveNewWidget(payload, widgetPath)
// 			wpArray = append(wpArray, newWidget)
// 			if err != nil {
// 				return err
// 			}
// 		} else if mode == "delete widget" {
// 			widget := p.Widget
// 			wpArray = nil
// 			for _, val := range widget {
// 				if payload.Get("widgetPageId", "") == val.ID {
// 					continue
// 				}
// 				wpArray = append(wpArray, val)
// 			}
// 		} else if mode == "save widget" {
// 			if payload.Get("widget") != nil {
// 				getwp := payload.Get("widget").(map[string]interface{})

// 				wp := &WidgetPage{}
// 				toolkit.UnjsonFromString(toolkit.JsonString(getwp), wp)

// 				p.ID = payload.Get("pageId", "").(string)
// 				if err := p.GetById(); err != nil {
// 					return err
// 				}
// 				widget := p.Widget
// 				wpArray = nil
// 				for _, val := range widget {
// 					if wp.ID == val.ID {
// 						val = wp
// 					}
// 					wpArray = append(wpArray, val)
// 				}
// 			}
// 		}
// 		p.Widget = wpArray
// 		// page.ThemeColor = payload.Get("themeColor", "").(string)
// 		// page.SendFiles(EC_DATA_PATH, payload.Get("serverId", "").(string))
// 	}

// 	if err := Save(p); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *Page) Delete(payload toolkit.M, path string) error {
// 	return nil
// }

// func (p *Page) GetServerPathSeparator(server *Server) string {
// 	if server.ServerType == "windows" {
// 		return "\\\\"
// 	}

// 	return `/`
// }

// func (p *Page) DefineCommand(server *Server, sourceZipPath string, destZipPath string, appID string) (string, string, string, error) {
// 	var ext string
// 	if strings.Contains(server.CmdExtract, "7z") || strings.Contains(server.CmdExtract, "zip") {
// 		ext = ".zip"
// 	} else if strings.Contains(server.CmdExtract, "tar") {
// 		ext = ".tar"
// 	} else if strings.Contains(server.CmdExtract, "gz") {
// 		ext = ".gz"
// 	}
// 	sourceZipFile := toolkit.Sprintf("%s%s", sourceZipPath, ext)
// 	destZipFile := toolkit.Sprintf("%s%s", destZipPath, ext)
// 	var unzipCmd string
// 	// cmd /C 7z e -o %s -y %s
// 	if server.ServerType == "windows" {
// 		unzipCmd = toolkit.Sprintf("cmd /C %s", server.CmdExtract)
// 		unzipCmd = strings.Replace(unzipCmd, `%1`, destZipPath, -1)
// 		unzipCmd = strings.Replace(unzipCmd, `%2`, destZipFile, -1)
// 	} else {
// 		unzipCmd = strings.Replace(server.CmdExtract, `%1`, destZipFile, -1)
// 		unzipCmd = strings.Replace(unzipCmd, `%2`, destZipPath, -1)
// 	}

// 	return unzipCmd, sourceZipFile, destZipFile, nil

// }

// func (p *Page) CopyFileToServer(server *Server, sourcePath string, destPath string, appID string, log *toolkit.LogEngine) error {
// 	var serverPathSeparator string
// 	if strings.Contains(destPath, "/") {
// 		serverPathSeparator = `/`
// 	} else {
// 		serverPathSeparator = "\\\\"
// 	}
// 	destZipPath := strings.Join([]string{destPath, appID}, serverPathSeparator)
// 	unzipCmd, sourceZipFile, destZipFile, err := p.DefineCommand(server, sourcePath, destZipPath, appID)

// 	log.AddLog(toolkit.Sprintf("Connect to server %v", server), "INFO")
// 	sshSetting, sshClient, err := p.connectSSH(server)
// 	defer sshClient.Close()

// 	log.AddLog(unzipCmd, "INFO") /*compress file on local colony manager*/
// 	if strings.Contains(sourceZipFile, ".zip") {
// 		err = toolkit.ZipCompress(sourcePath, sourceZipFile)
// 	} else if strings.Contains(sourceZipFile, ".tar") {
// 		err = toolkit.TarCompress(sourcePath, sourceZipFile)
// 	}
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	rmCmdZip := toolkit.Sprintf("rm -rf %s", destZipFile)
// 	log.AddLog(rmCmdZip, "INFO")
// 	_, err = sshSetting.GetOutputCommandSsh(rmCmdZip) /*delete zip file on server before copy file*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	log.AddLog(toolkit.Sprintf("scp from %s to %s", sourceZipFile, destPath), "INFO")
// 	err = sshSetting.SshCopyByPath(sourceZipFile, destPath) /*copy zip file from colony manager to server*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	rmCmdZipOutput := toolkit.Sprintf("rm -rf %s", destZipPath)
// 	log.AddLog(rmCmdZipOutput, "INFO")
// 	_, err = sshSetting.GetOutputCommandSsh(rmCmdZipOutput) /*delete folder before extract zip file on server*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	mkdirDestCmd := toolkit.Sprintf("%s %s%s%s", server.CmdMkDir, destZipPath, serverPathSeparator, appID)
// 	log.AddLog(mkdirDestCmd, "INFO")
// 	_, err = sshSetting.GetOutputCommandSsh(mkdirDestCmd) /*make new dest folder on server for folder extraction*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	chmodDestCmd := toolkit.Sprintf("chmod -R 755 %s%s%s", destZipPath, serverPathSeparator, appID)
// 	log.AddLog(chmodDestCmd, "INFO")
// 	_, err = sshSetting.GetOutputCommandSsh(chmodDestCmd) /*set chmod on new folder extraction*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	log.AddLog(unzipCmd, "INFO")
// 	_, err = sshSetting.GetOutputCommandSsh(unzipCmd) /*extract zip file to server*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	log.AddLog(toolkit.Sprintf("remove %s", sourceZipFile), "INFO")
// 	err = os.Remove(sourceZipFile) /*remove zip file from local colony manager*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}

// 	log.AddLog(rmCmdZip, "INFO")
// 	_, err = sshSetting.GetOutputCommandSsh(rmCmdZip) /*delete zip file on server after folder extraction*/
// 	if err != nil {
// 		log.AddLog(err.Error(), "ERROR")
// 		return err
// 	}
// 	return nil
// }

// func (p *Page) SendFiles(EC_DATA_PATH string, serverid string) error {
// 	path := filepath.Join(EC_DATA_PATH, "widget", "log")
// 	log, _ := toolkit.NewLog(false, true, path, "sendfile-%s", "20060102-1504")

// 	for _, wValue := range p.Widget {
// 		appID := wValue.ID
// 		log.AddLog("Get widget with ID: "+appID, "INFO")
// 		widget := new(Widget)
// 		err := Get(widget, appID)
// 		if err != nil {
// 			log.AddLog(err.Error(), "ERROR")
// 			return err
// 		}

// 		log.AddLog("Get server with ID: "+serverid, "INFO")
// 		server := new(Server)
// 		err = Get(server, serverid)
// 		if err != nil {
// 			log.AddLog(err.Error(), "ERROR")
// 			return err
// 		}

// 		serverPathSeparator := p.GetServerPathSeparator(server)
// 		sourcePath := filepath.Join(EC_DATA_PATH, "widget", appID)
// 		destPath := strings.Join([]string{server.AppPath, "src", "widget"}, serverPathSeparator)

// 		if server.OS == "windows" {
// 			if strings.Contains(server.CmdExtract, "7z") || strings.Contains(server.CmdExtract, "zip") {
// 				err = p.CopyFileToServer(server, sourcePath, destPath, appID, log)
// 				if err != nil {
// 					log.AddLog(err.Error(), "ERROR")
// 					return err
// 				}
// 			} else {
// 				message := "currently only zip/7z command which is supported"
// 				log.AddLog(message, "ERROR")
// 				return err
// 			}
// 		} else {
// 			if strings.Contains(server.CmdExtract, "tar") || strings.Contains(server.CmdExtract, "zip") {
// 				err = p.CopyFileToServer(server, sourcePath, destPath, appID, log)
// 				if err != nil {
// 					log.AddLog(err.Error(), "ERROR")
// 					return err
// 				}
// 			} else {
// 				message := "currently only zip/tar command which is supported"
// 				log.AddLog(message, "ERROR")
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func (p *Page) connectSSH(payload *Server) (sshclient.SshSetting, *ssh.Client, error) {
// 	client := sshclient.SshSetting{}
// 	client.SSHHost = payload.Host

// 	if payload.SSHType == "File" {
// 		client.SSHAuthType = sshclient.SSHAuthType_Certificate
// 		client.SSHKeyLocation = payload.SSHFile
// 	} else {
// 		client.SSHAuthType = sshclient.SSHAuthType_Password
// 		client.SSHUser = payload.SSHUser
// 		client.SSHPassword = payload.SSHPass
// 	}

// 	theClient, err := client.Connect()

// 	return client, theClient, err
// }

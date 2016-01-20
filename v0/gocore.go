package colonycore

import (
	"errors"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/jsons"
	//"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
	"os"
	"path/filepath"
	"strings"
)

var ConfigPath string

func validateConfig() error {
	if ConfigPath == "" {
		return errors.New("gocore.validateConfig: ConfigPath is empty")
	}
	_, e := os.Stat(ConfigPath)
	if e != nil {
		return errors.New("gocore.validateConfig: " + e.Error())
	}
	return nil
}

func getJsonFilePath(o interface{}) string {
	tns := strings.Split(strings.ToLower(toolkit.TypeName(o)), ".")
	fn := tns[len(tns)-1] + ".json"
	return filepath.Join(ConfigPath, fn)
}

func getConnection() (dbox.IConnection, error) {
	if e := validateConfig(); e != nil {
		return nil, errors.New("gocore.GetConnection: " + e.Error())
	}

	//jsonpath := getJsonFilePath(o)
	c, e := dbox.NewConnection("jsons", &dbox.ConnectionInfo{ConfigPath, "", "", "", toolkit.M{}.Set("newfile", true)})
	if e != nil {
		return nil, errors.New("gocore.GetConnection: " + e.Error())
	}
	e = c.Connect()
	if e != nil {
		return nil, errors.New("gocore.GetConnection: Connect: " + e.Error())
	}
	return c, nil
}

/*
func Populate(ms interface{}) error {
	o, e := toolkit.GetEmptySliceElement(ms)
	if e != nil {
		return errors.New("colonycore.Populate: " + e.Error())
	}
	c, e := getConnection(o)
	if e != nil {
		toolkit.Printf("Error: %s\n", e.Error())
		return nil
	}
	//ims := []toolkit.M{}
	ctx := orm.New(c)
	defer ctx.Close()
	var model orm.IModel
	model = new(Application)
	cursor, e := ctx.Find(model, nil)
	if e != nil {
		return errors.New(toolkit.Sprintf("colonycore.Populate: no cursor %s \n", e.Error()))
	}
	cursor.Fetch(&ms, 0, false)
	defer cursor.Close()
	return nil
}
*/

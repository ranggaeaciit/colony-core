package gocoretest

import (
	//"fmt"

	"github.com/eaciit/colony-core/v0"
	"github.com/eaciit/toolkit"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveApp(t *testing.T) {
	wd, _ := os.Getwd()
	colonycore.ConfigPath = filepath.Join(wd, "../config")

	appn := colonycore.CtxApplication().NewModel(&colonycore.Application{}).(*colonycore.Application)
	appn.ID = "test"
	appn.Enable = true
	//toolkit.Printf("ID of new appn: %s\n", toolkit.Id(appn))

	e := colonycore.CtxApplication().Save(appn)
	if e != nil {
		t.Errorf("Save Appn: " + e.Error())
	}
}

func TestLoadApp(t *testing.T) {
	//t.Skip()
	apps := []colonycore.Application{}
	c, e := colonycore.CtxApplication().Find(&colonycore.Application{}, nil)
	if e != nil {
		t.Errorf("Load appn:" + e.Error())
		return
	}
	e = c.Fetch(&apps, 0, false)
	if e != nil {
		t.Errorf("Load appn: fetch " + e.Error())
	}
	toolkit.Printf("Data:\n%v\n", toolkit.JsonString(apps))
}

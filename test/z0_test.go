package gocoretest

import (
	//"fmt"
	"github.com/eaciit/colony-core/v0"
	"github.com/eaciit/dbox"
	"github.com/eaciit/toolkit"
	"os"
	"path/filepath"
	"testing"
)

var e error

func TestSaveApp(t *testing.T) {
	wd, _ := os.Getwd()
	colonycore.ConfigPath = filepath.Join(wd, "../config")
	for i := 1; i <= 5; i++ {
		appn := new(colonycore.Application)
		appn.ID = toolkit.Sprintf("appn%d", i)
		appn.Enable = true
		e = colonycore.Save(appn)
		if e != nil {
			t.Fatalf("Save %s fail: %s", appn.ID, e.Error())
		}
	}

	appn := new(colonycore.Application)
	e := colonycore.Get(appn, "appn5")
	if e != nil {
		t.Fatal(e)
	}

	appn.ID = "appn3"
	e = colonycore.Delete(appn)
	if e != nil {
		t.Fatal(e)
	}
}

func TestLoadApp(t *testing.T) {
	//t.Skip()
	apps := []colonycore.Application{}
	c, e := colonycore.Find(new(colonycore.Application), dbox.Lte("_id", "appn4"))
	if e != nil {
		t.Errorf("Load appn fail:" + e.Error())
		return
	}
	e = c.Fetch(&apps, 0, false)
	if e != nil {
		t.Error("Fetching appn fail:" + e.Error())
	}
	toolkit.Printf("Applications: %s\n", toolkit.JsonString(apps))
}

func TestSaveQuery(t *testing.T) {
	var e error
	for i := 1; i <= 5; i++ {
		ds := new(colonycore.DataSource)
		ds.ID = toolkit.Sprintf("ds%d", i)
		ds.ConnectionID = "conn1"
		ds.QueryInfo = toolkit.M{}
		ds.MetaData = nil
		e = colonycore.Save(ds)
		if e != nil {
			t.Fatalf("Save datasource fail. " + e.Error())
		}
	}

	var dss []colonycore.DataSource
	c, e := colonycore.Find(new(colonycore.DataSource), nil)
	if e != nil {
		t.Fatalf("Load ds fail: " + e.Error())
	}

	e = c.Fetch(&dss, 0, true)
	if e != nil {
		t.Fatalf("Ftech ds fail: " + e.Error())
	}
	if len(dss) != 5 {
		t.Fatal("Fetch ds fail. Got %d records only", len(dss))
	}
	toolkit.Println("Data:", toolkit.JsonString(dss))
}

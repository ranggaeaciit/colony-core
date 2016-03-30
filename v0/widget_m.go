package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Widget struct {
	orm.ModelBase
	ID           string   `json:"_id"`
	Title        string   `json:"title"`
	DataSourceID []string `json:"dataSourceId"`
	Description  string   `json:"description"`
}

func (w *Widget) TableName() string {
	return "widgets"
}

func (w *Widget) RecordID() interface{} {
	return w.ID
}

func (w *Widget) Get(search string) ([]Widget, error) {
	var query *dbox.Filter

	if search != "" {
		query = dbox.Contains("_id", search)
	}

	data := []Widget{}
	cursor, err := Find(new(Widget), query)
	if err != nil {
		return nil, err
	}
	if err := cursor.Fetch(&data, 0, false); err != nil {
		return nil, err
	}
	defer cursor.Close()
	return data, nil
}

func (w *Widget) GetById() error {
	if err := Get(w, w.ID); err != nil {
		return err
	}
	return nil
}

func (w *Widget) Save() error {
	if err := Save(w); err != nil {
		return err
	}
	return nil
}

func (w *Widget) ExtractFile(compressedSource string, fileName string) error {
	compressedFile := filepath.Join(compressedSource, fileName)
	extractDest := filepath.Join(compressedSource, w.ID)

	if runtime.GOOS == "windows" {
		exec.Command("cmd", "-c", "rmdir", "/s", "/q", extractDest).Run()
	} else {
		exec.Command("rm", "-rf", extractDest).Run()
	}

	if strings.Contains(fileName, ".tar.gz") {
		if err := toolkit.TarGzExtract(compressedFile, extractDest); err != nil {
			return err
		}
	} else if strings.Contains(fileName, ".gz") {
		if err := toolkit.GzExtract(compressedFile, extractDest); err != nil {
			return err
		}
	} else if strings.Contains(fileName, ".tar") {
		if err := toolkit.TarExtract(compressedFile, extractDest); err != nil {
			return err
		}
	} else if strings.Contains(fileName, ".zip") {
		if err := toolkit.ZipExtract(compressedFile, extractDest); err != nil {
			return err
		}
	}
	os.Remove(compressedFile)

	return nil
}

func (w *Widget) Delete() error {
	if err := Delete(w); err != nil {
		return err
	}
	return nil
}

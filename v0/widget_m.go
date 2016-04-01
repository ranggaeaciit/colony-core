package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Widget struct {
	orm.ModelBase
	ID           string     `json:"_id"`
	Title        string     `json:"title"`
	DataSourceId []string   `json:"dataSourceId"`
	Description  string     `json:"description"`
	Config       toolkit.Ms `json:"config"`
	Params       toolkit.M  `json:"params"`
	URL          string     `json:"url"`
}

type DataSourceWidget struct {
	ID   string      `json:"_id"`
	Data []toolkit.M `json:"data"`
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

func GetPath(root string) (string, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	var _path string
	var indexPath []string
	var configPath []string
	var assetPath []string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		_path, filename := filepath.Split(path)
		if strings.Compare(filename, "index.html") == 0 {
			indexPath = append(indexPath, _path)
		}
		if strings.Compare(filename, "config.json") == 0 {
			configPath = append(configPath, _path)
		}

		if strings.Compare(filename, "assets") == 0 {
			assetPath = append(assetPath, _path)
		}
		return nil
	}
	if err = filepath.Walk(absRoot, walkFunc); err != nil {
		if err.Error() == "found" {
			return _path, nil
		} else {
			return "", err
		}
	}

	for _, valIndex := range indexPath {
		for _, valConfig := range configPath {
			for _, valAsset := range assetPath {
				if valIndex == valConfig && valConfig == valAsset {
					_path = valConfig
					break
				}
			}

		}
	}
	return _path, nil
}

func (w *Widget) ExtractFile(compressedSource string, fileName string) (toolkit.Ms, error) {
	compressedFile := filepath.Join(compressedSource, fileName)
	extractDest := filepath.Join(compressedSource, w.ID)

	if err := os.RemoveAll(extractDest); err != nil {
		return nil, err
	}

	if strings.Contains(fileName, ".tar.gz") {
		if err := toolkit.TarGzExtract(compressedFile, extractDest); err != nil {
			return nil, err
		}
	} else if strings.Contains(fileName, ".gz") {
		if err := toolkit.GzExtract(compressedFile, extractDest); err != nil {
			return nil, err
		}
	} else if strings.Contains(fileName, ".tar") {
		if err := toolkit.TarExtract(compressedFile, extractDest); err != nil {
			return nil, err
		}
	} else if strings.Contains(fileName, ".zip") {
		if err := toolkit.ZipExtract(compressedFile, extractDest); err != nil {
			return nil, err
		}
	}

	if err := os.Remove(compressedFile); err != nil {
		return nil, err
	}

	path, err := GetPath(extractDest)
	if err != nil {
		return nil, err
	}

	urlPath := filepath.ToSlash(path)
	splitPath := strings.SplitAfter(urlPath, "/data-root/widget/")
	w.URL = strings.Join([]string{w.URL, "res-widget", splitPath[1]}, "/")

	getConfigFile := filepath.Join(path, "config.json")
	content, err := ioutil.ReadFile(getConfigFile)
	if err != nil {
		return nil, err
	}

	result := toolkit.Ms{}
	err = toolkit.Unjson(content, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (w *Widget) Delete(compressedSource string) error {
	extractDest := filepath.Join(compressedSource, w.ID)
	if err := Delete(w); err != nil {
		return err
	}

	if err := os.RemoveAll(extractDest); err != nil {
		return err
	}
	return nil
}

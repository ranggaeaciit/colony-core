package colonycore

import (
	"github.com/eaciit/toolkit"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetPath(root string, _filename string) ([]string, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	var pathList []string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		_path, filename := filepath.Split(path)
		if filename == _filename {
			pathList = append(pathList, _path)
		}
		return nil
	}
	if err = filepath.Walk(absRoot, walkFunc); err != nil {
		return nil, err
	}
	return pathList, nil
}

func GetWidgetPath(path string) (string, error) {
	indexPath, err := GetPath(path, "index.html")
	if err != nil {
		return "", err
	}
	configPath, err := GetPath(path, "config.json")
	if err != nil {
		return "", err
	}
	assetPath, err := GetPath(path, "assets")
	if err != nil {
		return "", err
	}
	var _path string
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

func GetJsonFile(pathfile string) (toolkit.Ms, error) {
	content, err := ioutil.ReadFile(pathfile)
	if err != nil {
		return nil, err
	}
	result := toolkit.Ms{}
	if err := toolkit.Unjson(content, &result); err != nil {
		return nil, err
	}
	return result, nil
}

package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"path/filepath"
)

type Selector struct {
	orm.ModelBase
	ID               string         `json:"_id"`
	MasterDataSource string         `json:"masterDataSource"`
	Title            string         `json:"title"`
	Fields           []*FieldDetail `json:"fields"`
}

type FieldDetail struct {
	ID         string `json:"_id"`
	DataSource string `json:"dataSource"`
	Field      string `json:"field"`
}

func (sl *Selector) TableName() string {
	return filepath.Join("widget", "selectors")
}

func (sl *Selector) RecordID() interface{} {
	return sl.ID
}

func (sl *Selector) Get(search string) ([]Selector, error) {
	var query *dbox.Filter

	if search != "" {
		query = dbox.Contains("_id", search)
	}

	data := []Selector{}
	cursor, err := Find(new(Selector), query)
	if err != nil {
		return nil, err
	}
	if err := cursor.Fetch(&data, 0, false); err != nil {
		return nil, err
	}
	defer cursor.Close()
	return data, nil
}

func (sl *Selector) GetById() error {
	if err := Get(sl, sl.ID); err != nil {
		return err
	}
	return nil
}

func (sl *Selector) Save() error {
	if err := Save(sl); err != nil {
		return err
	}
	return nil
}

func (sl *Selector) Delete() error {
	if err := Delete(sl); err != nil {
		return err
	}
	return nil
}

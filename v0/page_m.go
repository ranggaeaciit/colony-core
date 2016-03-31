package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
)

type Page struct {
	orm.ModelBase
	ID         string     `json:"_id"`
	PanelID    []*PanelID `json:"panelId"`
	ParentMenu string     `json:"parentMenu"`
	Title      string     `json:"title"`
	URL        string     `json:"url"`
}

type PanelID struct {
	ID       string      `json:"_id"`
	WidgetID []*WidgetID `json:"widgetId"`
}

type WidgetID struct {
	ID string `json:"_id"`
}

func (p *Page) TableName() string {
	return "pages"
}

func (p *Page) RecordID() interface{} {
	return p.ID
}

func (p *Page) Get(search string) ([]Page, error) {
	var query *dbox.Filter

	if search != "" {
		query = dbox.Contains("_id", search)
	}

	data := []Page{}
	cursor, err := Find(new(Page), query)
	if err != nil {
		return nil, err
	}
	if err := cursor.Fetch(&data, 0, false); err != nil {
		return nil, err
	}
	defer cursor.Close()
	return data, nil
}

func (p *Page) GetById() error {
	if err := Get(p, p.ID); err != nil {
		return err
	}
	return nil
}

func (p *Page) Save() error {
	if err := Save(p); err != nil {
		return err
	}
	return nil
}

func (p *Page) Delete() error {
	if err := Delete(p); err != nil {
		return err
	}
	return nil
}

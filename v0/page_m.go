package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
)

type Page struct {
	orm.ModelBase
	ID          string   `json:"_id"`
	DataSources []string `json:"dataSources"`
	// Panel       []*PanelPage `json:"panel"`
	Widget     []*WidgetPage `json:"widget"`
	ParentMenu string        `json:"parentMenu"`
	Title      string        `json:"title"`
	URL        string        `json:"url"`
	ThemeColor string        `json:"themeColor"`
}

/*type PanelPage struct {
	ID     string        `json:"_id"`
	Title  string        `json:"title"`
	Offset string        `json:"offset"`
	Width  int           `json:"width"`
	Widget []*WidgetPage `json:"widget"`
}*/

type WidgetPage struct {
	ID            string     `json:"_id"`
	Title         string     `json:"title"`
	WidgetType    string     `json:"widgetType"`
	DataSourceID  string     `json:"dataSourceId"`
	SettingWidget toolkit.Ms `json:"settingWidget"`
	Height        int        `json:"height"`
	Width         int        `json:"width"`
	PanelWidgetID string     `json:"panelWidgetID"`
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

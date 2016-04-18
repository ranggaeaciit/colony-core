package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"os"
	"path/filepath"
)

type MapChart struct {
	orm.ModelBase
	ID        string `json:"_id"`
	ChartName string `json:"chartName"`
	FileName  string `json:"fileName"`
}

type Chart struct {
	orm.ModelBase
	ID             string          `json:"_id"`
	Title          string          `json:"title"`
	DataSourceID   string          `json:"dataSourceID"`
	DataSource     *DataSourceData `json: "dataSource"`
	ChartArea      *ChartArea      `json:"chartArea"`
	Outsider       *OutSider       `json:"outsider"`
	SeriesDefaults *SeriesDefault  `json:"seriesDefaults"`
	Series         []*Series       `json:"series"`
	CategoryAxis   *CategoryAxis   `json:"categoryAxis"`
	ValueAxis      *ValueAxis      `json:"valueAxis"`
	Tooltip        *Tooltip        `json:"tooltip"`
	Legend         *Legend         `json:"legend"`
}

type DataSourceData struct {
	Data []*map[string]interface{} `json:"data"`
}

type OutSider struct {
	WidthMode           string `json:"widthMode"`
	HeightMode          string `json:"heightMode"`
	ValueAxisUseMaxMode bool   `json:"valueAxisUseMaxMode"`
	ValueAxisUseMinMode bool   `json:"valueAxisUseMinMode"`
}

type ChartArea struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type Legend struct {
	Visible bool `json:"visible"`
}

type Series struct {
	Field string `json:"field"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type ValueAxis struct {
	Max            int            `json:"max"`
	Min            int            `json:"min"`
	Line           *Line          `json:"line"`
	MinorGridLines *MinorGridLine `json:"minorGridLines"`
	Labels         *Label         `json:"labels"`
}

type Line struct {
	Visible bool `json: "visible"`
}

type MinorGridLine struct {
	Visible bool `json: "visible"`
}

type Label struct {
	Rotation int `json: "rotation"`
}

type CategoryAxis struct {
	Field string `json:"field"`
}

type Tooltip struct {
	Visible  bool   `json:"visible"`
	Template string `json:"template"`
}

type SeriesDefault struct {
	Type string `json: "type"`
}

func (mc *MapChart) TableName() string {
	return filepath.Join("widget", "mapcharts")
}

func (mc *MapChart) RecordID() interface{} {
	return mc.ID
}

func (mc *MapChart) Get(search string) ([]MapChart, error) {
	var query *dbox.Filter

	if search != "" {
		query = dbox.Contains("_id", search)
	}

	mapchart := []MapChart{}
	cursor, err := Find(new(MapChart), query)
	if err != nil {
		return mapchart, err
	}

	err = cursor.Fetch(&mapchart, 0, false)
	if err != nil {
		return mapchart, err
	}
	defer cursor.Close()
	return mapchart, nil
}

func (mc *MapChart) Delete() error {
	if err := Delete(mc); err != nil {
		return err
	}
	return nil
}

func (c *Chart) TableName() string {
	return filepath.Join("widget", "chart", c.ID)
}

func (c *Chart) RecordID() interface{} {
	return c.ID
}

func (c *Chart) GetById() error {
	if err := Get(c, c.ID); err != nil {
		return err
	}
	return nil
}

func (c *Chart) Save() error {
	newChart := MapChart{}
	mapchart, err := newChart.Get("")
	if err != nil {
		return err
	}

	var isUpdate bool

	for _, eachRaw := range mapchart {
		if eachRaw.FileName == c.ID+".json" {
			eachRaw.ChartName = c.Title
			isUpdate = true
			newChart = eachRaw
		}
	}

	if !isUpdate {
		newChart.ID = c.ID
		newChart.FileName = c.ID + ".json"
		newChart.ChartName = c.Title
	}

	if err := Save(&newChart); err != nil {
		return err
	}

	if err := Save(c); err != nil {
		return err
	}
	return nil
}

func (c *Chart) Remove() error {
	_file := filepath.Join(ConfigPath, "widget", "chart", c.ID+".json")
	if err := os.Remove(_file); err != nil {
		return err
	}
	return nil
}

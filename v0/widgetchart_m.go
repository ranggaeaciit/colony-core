package colonycore

import (
	// "github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	// "os"
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
	ID                string          `json:"_id"`
	Outsiders         *Outsiders      `json:"outsiders"`
	Title             string          `json:"title"`
	DataSourceID      string          `json:"dataSourceID"`
	ChartArea         *ChartArea      `json:"chartArea"`
	dataSource        *DataSources    `json:"dataSource"`
	Legend            *Legend         `json:"legend"`
	SeriesDefaultType string          `json:"seriesDefaultType"`
	Series            *Series         `json:"series"`
	ValueAxis         *ValueAxis      `json:"valueAxis"`
	CategoryAxis      []*CategoryAxis `json:"categoryAxis"`
	Tooltip           *Tooltip        `json:"tooltip"`
}

type Outsiders struct {
	WidthMode           string `json:"widthMode"`
	HeightMode          string `json:"heightMode"`
	ValueAxisUseMaxMode bool   `json:"valueAxisUseMaxMode"`
	ValueAxisUseMinMode bool   `json:"valueAxisUseMinMode"`
}

type ChartArea struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type DataSources struct {
	data []string
}

type Legend struct {
	Visible bool `json:"visible"`
}

type Series struct {
	Field string `json:"field"`
	Name  string `json:"name"`
	Types bool   `json:"types"`
}

type ValueAxis struct {
	Max            int  `json:"max"`
	Min            int  `json:"min"`
	Types          bool `json:"types"`
	Line           bool `json:"line"`
	MinorGridLines bool `json:"minorGridLines"`
	LabelsRotation int  `json:"labelsRotation"`
}

type CategoryAxis struct {
	Field string `json:"field"`
}

type Tooltip struct {
	Visible  bool   `json:"visible"`
	Template string `json:"template"`
}

func (mc *MapChart) TableName() string {
	return filepath.Join("widget", "mapcharts")
}

func (mc *MapChart) RecordID() interface{} {
	return mc.ID
}

func (c *Chart) TableName() string {
	return filepath.Join("widget", "chart", c.ID)
}

func (c *Chart) RecordID() interface{} {
	return c.ID
}

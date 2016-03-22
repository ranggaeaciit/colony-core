package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Grid struct {
	orm.ModelBase
	ID           string            `json:"id"`
	Title        string            `json:"title"`
	DataSourceID string            `json:"dataSourceKey"`
	Aggregate    []AggregateColumn `json:"aggregate"`
	Outsider     Outsider          `json:"outsider"`
	PageSize     int               `json:"pageSize"`
	Groupable    bool              `json:"groupable"`
	Sortable     bool              `json:"sortable"`
	Filterable   bool              `json:"filterable"`
	Pageable     Pageable          `json:"pageable"`
	Columns      []Column          `json:"columns"`
	ColumnMenu   bool              `json:"columnMenu"`
	Toolbar      []string          `json:"toolbar"`
	Pdf          ExportGrid        `json:"pdf"`
	Excel        ExportGrid        `json:"excel"`
}

type Outsider struct {
	// IdGrid        string `json:"idGrid"`
	// Title         string `json:"title"`
	// DataSourceKey string `json:"dataSourceKey"`
	VisiblePDF   bool `json:"visiblePDF"`
	VisibleExcel bool `json:"visibleExcel"`
}

type ExportGrid struct {
	AllPages string `json:"allPages"`
	FileName string `json:"fileName"`
}

/*type DataSource struct {
	Aggregate []AggregateColumn `json:"aggregate"`
}*/

type AggregateColumn struct {
	Field     string `json:"field"`
	Aggregate string `json:"aggregate"`
}

type Pageable struct {
	Refresh     bool `json:"refresh"`
	PageSize    bool `json:"pageSize"`
	ButtonCount int  `json:"buttonCount"`
}

type Column struct {
	Template         string           `json:"template"`
	Field            string           `json:"field"`
	Title            string           `json:"title"`
	Format           string           `json:"format"`
	Width            string           `json:"width"`
	Menu             bool             `json:"menu"`
	HeaderTemplate   string           `json:"headerTemplate"`
	HeaderAttributes HeaderAttributes `json:"headerAttributes"`
	FooterTemplate   string           `json:"footerTemplate"`
}

type HeaderAttributes struct {
	Class string `json:"class"`
	Style string `json:"style"`
}

type MapGrid struct {
	orm.ModelBase
	ID   int           `json:"id"`
	Data []DataMapGrid `json:"data"`
}

type DataMapGrid struct {
	orm.ModelBase
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (mg *MapGrid) TableName() string {
	return "mapgrids"
}

func (mg *MapGrid) RecordID() interface{} {
	return mg.ID
}

func (dmg *DataMapGrid) TableName() string {
	return "datamapgrids"
}

func (dmg *DataMapGrid) RecordID() interface{} {
	return dmg.ID
}

func (g *Grid) TableName() string {
	return "grids"
}

func (g *Grid) RecordID() interface{} {
	return g.ID
}

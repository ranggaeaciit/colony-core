package colonycore

import (
	"github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
	"path/filepath"
)

/*========================WIDGET GRID ================================*/
type MapGrid struct {
	orm.ModelBase
	ID       string `json:"_id"`
	GridName string `json:"gridName"`
	FileName string `json:"fileName"`
}

type Grid struct {
	orm.ModelBase
	ID           string             `json:"_id"`
	Title        string             `json:"title"`
	DataSourceID string             `json:"dataSourceID"`
	Aggregate    []*AggregateColumn `json:"aggregate"`
	Outsider     *Outsider          `json:"outsider"`
	PageSize     int                `json:"pageSize"`
	Groupable    bool               `json:"groupable"`
	Sortable     bool               `json:"sortable"`
	Filterable   bool               `json:"filterable"`
	Pageable     *Pageable          `json:"pageable"`
	Columns      []*Column          `json:"columns"`
	ColumnMenu   bool               `json:"columnMenu"`
	Toolbar      []string           `json:"toolbar"`
	Pdf          ExportGrid         `json:"pdf"`
	Excel        ExportGrid         `json:"excel"`
}

type Outsider struct {
	VisiblePDF   bool `json:"visiblePDF"`
	VisibleExcel bool `json:"visibleExcel"`
}

type ExportGrid struct {
	AllPages string `json:"allPages"`
	FileName string `json:"fileName"`
}

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
	Template         string            `json:"template"`
	Field            string            `json:"field"`
	Title            string            `json:"title"`
	Format           string            `json:"format"`
	Width            string            `json:"width"`
	Menu             bool              `json:"menu"`
	HeaderTemplate   string            `json:"headerTemplate"`
	HeaderAttributes *HeaderAttributes `json:"headerAttributes"`
	FooterTemplate   string            `json:"footerTemplate"`
}

type HeaderAttributes struct {
	Class string `json:"class"`
	Style string `json:"style"`
}

/*type MapGrid struct {
	orm.ModelBase
	ID   int           `json:"_id"`
	Data []DataMapGrid `json:"data"`
}*/

/*======================== END OF WIDGET GRID ================================*/

/*========================WIDGET SELECTOR ================================*/

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

/*======================== END OF WIDGET SELECTOR ================================*/

/*======================== WIDGET CHART ================================*/

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
	dataSource        toolkit.M       `json:"dataSource"`
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

/*======================== END OF WIDGET CHART ================================*/

func (sl *Selector) TableName() string {
	return filepath.Join("widget", "selectors")
}

func (sl *Selector) RecordID() interface{} {
	return sl.ID
}

func (mg *MapGrid) TableName() string {
	return filepath.Join("widget", "mapgrids")
}

func (mg *MapGrid) RecordID() interface{} {
	return mg.ID
}

/*func (dmg *DataMapGrid) TableName() string {
	return filepath.Join("widget", "datamapgrids")
}

func (dmg *DataMapGrid) RecordID() interface{} {
	return dmg.ID
}*/

func (g *Grid) TableName() string {
	return filepath.Join("widget", "grid", g.ID)
}

func (g *Grid) RecordID() interface{} {
	return g.ID
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

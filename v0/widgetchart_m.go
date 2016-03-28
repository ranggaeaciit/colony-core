package colonycore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm/v1"
	"os"
	"path/filepath"
)

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
	Pdf          *ExportGrid        `json:"pdf"`
	Excel        *ExportGrid        `json:"excel"`
}

type AggregateColumn struct {
	Field     string `json:"field"`
	Aggregate string `json:"aggregate"`
}

type Outsider struct {
	VisiblePDF   bool `json:"visiblePDF"`
	VisibleExcel bool `json:"visibleExcel"`
}

type Pageable struct {
	Refresh     bool `json:"refresh"`
	PageSize    bool `json:"pageSize"`
	ButtonCount int  `json:"buttonCount"`
}

type ExportGrid struct {
	AllPages bool   `json:"allPages"`
	FileName string `json:"fileName"`
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

func (mg *MapGrid) TableName() string {
	return filepath.Join("widget", "mapgrids")
}

func (mg *MapGrid) RecordID() interface{} {
	return mg.ID
}

func (mg *MapGrid) Get(search string) ([]MapGrid, error) {
	var query *dbox.Filter

	if search != "" {
		query = dbox.Contains("_id", search)
	}

	mapgrid := []MapGrid{}
	cursor, err := Find(new(MapGrid), query)
	if err != nil {
		return mapgrid, err
	}

	err = cursor.Fetch(&mapgrid, 0, false)
	if err != nil {
		return mapgrid, err
	}
	defer cursor.Close()
	return mapgrid, nil
}

func (mg *MapGrid) Delete() error {
	if err := Delete(mg); err != nil {
		return err
	}
	return nil
}

func (g *Grid) TableName() string {
	return filepath.Join("widget", "grid", g.ID)
}

func (g *Grid) RecordID() interface{} {
	return g.ID
}

func (g *Grid) GetById() error {
	if err := Get(g, g.ID); err != nil {
		return err
	}
	return nil
}

func (g *Grid) Save() error {
	newGrid := MapGrid{}
	mapgrid, err := newGrid.Get("")
	if err != nil {
		return err
	}

	var isUpdate bool

	for _, eachRaw := range mapgrid {
		if eachRaw.FileName == g.ID+".json" {
			eachRaw.GridName = g.Title
			isUpdate = true
			newGrid = eachRaw
		}
	}

	if !isUpdate {
		newGrid.ID = g.ID
		newGrid.FileName = g.ID + ".json"
		newGrid.GridName = g.Title
	}

	if err := Save(&newGrid); err != nil {
		return err
	}

	if err := Save(g); err != nil {
		return err
	}
	return nil
}

func (g *Grid) Remove() error {
	_file := filepath.Join(ConfigPath, "widget", "grid", g.ID+".json")
	if err := os.Remove(_file); err != nil {
		return err
	}
	return nil
}

package colonycore

import (
	// "github.com/eaciit/orm/v1"
	// "github.com/eaciit/toolkit"
	"time"
)

type FileInfo struct {
	Name         string      `json:"name", bson:"name"`
	Size         float64     `json:"size", bson:"size"`
	User         string      `json:"user", bson:"user"`
	Group        string      `json:"group", bson:"group"`
	Permissions  string      `json:"permissions", bson:"permissions"`
	CreatedDate  time.Time   `json:"createddate", bson:"createddate"`
	LastModified time.Time   `json:"lastmodified", bson:"lastmodified"`
	Type         string      `json:"type", bson:"type"`
	Sub          []*FileInfo `json:"sub", bson:"sub"`
}

/*func Construct(line string, result []FileInfo) {
	if !toolkit.IsPointer(result) {
		e = errorlib.Error("", "", "Fetch", "Model object should be pointer")
		return
	}

}*/

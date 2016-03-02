package colonycore

import (
	// "github.com/eaciit/orm/v1"
	"github.com/eaciit/errorlib"
	"github.com/eaciit/toolkit"
	"strconv"
	"strings"
	"time"
)

/*const (
	FILE_TYPE = map[string]string{
		"-": "FILE",
		"d": "DIRECTORY",
		"l": "SYMLINK",
	}
)*/

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
	SubCount     int64       `json:"subcount", bson:"subcount"`
}

func Construct(lines string) ([]FileInfo, error) {
	var result []FileInfo

	aLine := strings.Split(lines, "\n")

	if len(aLine) > 2 {
		for _, val := range aLine[1:] {
			res, e := parse(val)
			if e != nil {
				return result, e
			} else {

				result = append(result, res)
			}
		}

		return result, nil
	} else {
		res, e := parse(aLine[1])
		result = append(result, res)
		return result, e
	}

}

func parse(line string) (result FileInfo, e error) {
	if line != "" {
		cols := strings.Split(strings.Trim(line, " "), "||")

		// log.Printf("--------- cols: %v\n%v\n", len(cols), cols)

		if len(cols) == 9 {
			result.Type = strings.TrimSpace(cols[0][:1])
			result.Sub = nil

			result.Permissions = strings.TrimSpace(cols[0][1:])

			subCount, _ := strconv.ParseInt(strings.TrimSpace(cols[1]), 10, 64)
			result.SubCount = subCount

			result.User = strings.TrimSpace(cols[2])

			result.Group = strings.TrimSpace(cols[3])

			size, _ := strconv.ParseFloat(strings.TrimSpace(cols[4]), 64)
			result.Size = size

			var lastModified time.Time
			str := strings.TrimSpace(cols[5]) + "-" + strings.TrimSpace(cols[6]) + "-" + strings.TrimSpace(cols[7])

			// log.Printf("str: %v\n", str)

			if len(strings.TrimSpace(cols[7])) == 5 {
				str = str + "-" + strconv.Itoa(time.Now().Year())
				// log.Printf("str: %v\n", str)
				lastModified = toolkit.String2Date(str, "MMM-dd-H:mm-YYYY")
			} else {
				lastModified = toolkit.String2Date(str, "MMM-dd-YYYY")
			}

			result.LastModified = lastModified

			result.Name = strings.TrimSpace(cols[8])
		} else {
			e = errorlib.Error("", "", "parse", "column is not valid")
		}
	}

	return
}

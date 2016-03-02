package colonycore

import (
	"github.com/eaciit/errorlib"
	// "github.com/eaciit/orm/v1"
	"github.com/eaciit/toolkit"
	"strconv"
	"strings"
	"time"
)

const (
	/*FILE_TYPE = map[string]string{
		"-": "FILE",
		"d": "DIRECTORY",
		"l": "SYMLINK",
	}*/
	DELIMITER = "/"
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
	SubCount     int64       `json:"subcount", bson:"subcount"`
	IsDir        bool        `json:"isdir", bson:"isdir"`
	Path         string      `json:"path", bson:"path"`
}

func ConstructFileInfo(lines string, path string) ([]FileInfo, error) {
	var result []FileInfo

	aLine := strings.Split(lines, "\n")

	if len(aLine) > 2 {
		for _, val := range aLine[1:] {
			if val != "" {
				res, e := parse(val, path)
				if e != nil {
					return result, e
				} else {
					result = append(result, res)
				}
			}
		}

		return result, nil
	} else {
		res, e := parse(aLine[1], path)
		result = append(result, res)
		return result, e
	}

}

func parse(line string, path string) (result FileInfo, e error) {
	if line != "" {
		cols := strings.Split(strings.Trim(line, " "), "||")

		if len(cols) == 9 {
			result.Type = strings.TrimSpace(cols[0][:1])
			result.Sub = nil

			if result.Type == "d" {
				result.IsDir = true
			}

			result.Permissions = strings.TrimSpace(cols[0][1:])

			subCount, _ := strconv.ParseInt(strings.TrimSpace(cols[1]), 10, 64)
			result.SubCount = subCount

			result.User = strings.TrimSpace(cols[2])

			result.Group = strings.TrimSpace(cols[3])

			size, _ := strconv.ParseFloat(strings.TrimSpace(cols[4]), 64)
			result.Size = size

			var lastModified time.Time
			str := strings.TrimSpace(cols[5]) + "-" + strings.TrimSpace(cols[6]) + "-" + strings.TrimSpace(cols[7])

			if len(strings.TrimSpace(cols[7])) == 5 {
				str = str + "-" + strconv.Itoa(time.Now().Year())
				lastModified = toolkit.String2Date(str, "MMM-d-H:mm-YYYY")
			} else {
				lastModified = toolkit.String2Date(str, "MMM-d-YYYY")
			}

			result.LastModified = lastModified

			result.Name = strings.TrimSpace(cols[8])

			if strings.LastIndex(path, DELIMITER) == -1 {
				result.Path = path + DELIMITER + result.Name
			} else {
				result.Path = path + result.Name
			}

		} else {
			e = errorlib.Error("", "", "parse", "column is not valid")
		}
	}

	return
}

package compare

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"sorting/folder"
	"strings"
)

type parsedJSON struct {
	fileName string
	parsed   map[string]map[string]interface{}
}

type differenceInfo struct {
	fileName string
	length   int
}

// Start ...
func Start(JSONAndPath []folder.JSONAndPath) {

	var sliceOfJSON = getParsedJSON(JSONAndPath)

	count := 0

	JSONItemPrev := sliceOfJSON[0]
	JSONItemNext := sliceOfJSON[1]

	for key, _ := range JSONItemPrev.parsed {
		var prev = differenceInfo{JSONItemPrev.fileName, len(JSONItemPrev.parsed[key])}
		var next = differenceInfo{JSONItemNext.fileName, len(JSONItemNext.parsed[key])}

		if prev.length != next.length {
			count++
			fmt.Printf("\n - Ключи \"%s\" отличаются: \n %s\n", key, showDifference(prev, next))
		}
	}

	if count == 0 {
		fmt.Printf("\n + Файлы соответствуют друг другу \n\n")
	} else {
		fmt.Printf("\n - Количество различий: %d\n\n", count)
	}
}

func getParsedJSON(sortingJSONFiles []folder.JSONAndPath) []parsedJSON {

	var result []parsedJSON

	for _, item := range sortingJSONFiles {
		var output map[string]map[string]interface{}
		if err := json.Unmarshal(item.Content, &output); err != nil {
			log.Fatal(err)
		}
		var name = path.Base(item.Path)
		result = append(result, parsedJSON{name, output})
	}

	return result
}

func showDifference(info ...differenceInfo) string {

	var result []string

	for i := 0; i < len(info); i++ {
		name, length := info[i].fileName, info[i].length
		result = append(result, fmt.Sprintf("(%s: %d)", name, length))
	}

	return strings.Join(result, " / ")
}

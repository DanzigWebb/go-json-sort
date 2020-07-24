package compare

import (
	"encoding/json"
	"fmt"
	"log"
)

// Start ...
func Start(sortingJSONFiles [][]byte) {

	var sliceOfJSON = getParsedJSON(sortingJSONFiles)

	count := 0

	for key, element := range sliceOfJSON[0] {
		prev, next := len(element), len(sliceOfJSON[1][key])
		if prev != next {
			count++
			fmt.Printf("\n - Ключи \"%s\" отличаются %s\n", key, showDifference(prev, next))
		}
	}

	if count == 0 {
		fmt.Printf("\n + Файлы соответствуют друг другу \n\n")
	} else {
		fmt.Printf("\n - Количество различий: %d\n\n", count)
	}
}

func getParsedJSON(sortingJSONFiles [][]byte) []map[string]map[string]interface{} {
	var resultSlice []map[string]map[string]interface{}

	for i := 0; i < len(sortingJSONFiles); i++ {
		var result map[string]map[string]interface{}
		if err := json.Unmarshal(sortingJSONFiles[i], &result); err != nil {
			log.Fatal(err)
		}
		resultSlice = append(resultSlice, result)
	}

	return resultSlice
}

func showDifference(prev, next int) string {
	return fmt.Sprintf("(%d / %d)", prev, next)
}

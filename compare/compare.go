package compare

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"sorting/folder"
	"strings"
)

type JSONParsed map[string]map[string]interface{}

type JSONFile struct {
	fileName string
	parsed   JSONParsed
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

	for key := range JSONItemPrev.parsed {
		var prev = differenceInfo{JSONItemPrev.fileName, len(JSONItemPrev.parsed[key])}
		var next = differenceInfo{JSONItemNext.fileName, len(JSONItemNext.parsed[key])}

		if prev.length != next.length {
			count++
			fmt.Printf("\n - Ключи \"%s\" отличаются: \n%s\n", key, showDifference(prev, next))
		}

		showUnresolvedKeys(JSONItemPrev, JSONItemNext, key)
	}

	if count == 0 {
		fmt.Printf("\n + Кол-во ключей одинаковое \n\n")
	} else {
		fmt.Printf("\n - Различий строк в колонках: %d\n\n", count)
	}
}

func getParsedJSON(sortingJSONFiles []folder.JSONAndPath) []JSONFile {

	var result []JSONFile

	for _, item := range sortingJSONFiles {
		var output JSONParsed
		if err := json.Unmarshal(item.Content, &output); err != nil {
			log.Fatal(err)
		}
		var name = path.Base(item.Path)
		result = append(result, JSONFile{name, output})
	}

	return result
}

func showDifference(info ...differenceInfo) string {

	var result []string

	for i := 0; i < len(info); i++ {
		result = append(result, fmt.Sprintf("%d", info[i].length))
	}

	return strings.Join(result, "/")
}

func showUnresolvedKeys(prevItem, nextItem JSONFile, key string) {

	getMaxMin := func() (JSONFile, JSONFile) {
		if len(prevItem.parsed[key]) > len(nextItem.parsed[key]) {
			return prevItem, nextItem
		} else {
			return nextItem, prevItem
		}
	}

	getUnresolved := func(first, second JSONParsed, key string) string {
		var output []string
		for item := range first[key] {
			if second[key][item] == nil {
				output = append(output, item)
			}
		}
		var unresolvedKeys string
		for _, str := range output {
			unresolvedKeys += fmt.Sprintf("\n • %s.%s", key, str)
		}

		return unresolvedKeys
	}

	max, min := getMaxMin()

	minUnresolvedKeys := getUnresolved(max.parsed, min.parsed, key)
	maxUnresolvedKeys := getUnresolved(min.parsed, max.parsed, key)

	if len(minUnresolvedKeys) > 0 {
		fmt.Printf("в %s отсутствуют: %s \n", min.fileName, minUnresolvedKeys)
	}

	if len(maxUnresolvedKeys) > 0 {
		fmt.Printf("\nв %s отсутствуют: %s \n", max.fileName, maxUnresolvedKeys)
	}

	if len(minUnresolvedKeys) > 0 || len(maxUnresolvedKeys) > 0 {
		fmt.Printf("______________________\n")
	}
}

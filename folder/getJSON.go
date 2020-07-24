package folder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"sorting/folderpath"
)

// GetJSONAndPathes returns output JSON files and pathes to original files in slices
func GetJSONAndPathes(langs []string) ([][]byte, []string) {
	var sortingJSON [][]byte
	var fileNames []string

	for _, file := range langs {
		path := filepath.Join(folderpath.Path, file)
		fileNames = append(fileNames, path)
		initialFile := readJSONFile(path)
		outputFile := outputJSON(initialFile)
		sortingJSON = append(sortingJSON, outputFile)
	}

	return sortingJSON, fileNames
}

// FindJSONInFolder return JSON files in current folder
func FindJSONInFolder() []string {
	folderFiles, _ := ioutil.ReadDir(folderpath.Path)

	var JSONFiles []string
	for _, file := range folderFiles {
		if filepath.Ext(file.Name()) == ".json" {
			JSONFiles = append(JSONFiles, file.Name())
		}
	}

	return JSONFiles
}


func readJSONFile(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func outputJSON(file []byte) []byte {
	var result map[string]interface{}

	if err := json.Unmarshal(file, &result); err != nil {
		log.Fatal(err)
	}

	sortJSON, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	return sortJSON
}
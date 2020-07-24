package folder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"sorting/flags"
)

// GetJSONAndPathes returns output JSON files and pathes to original files in slices
func GetJSONAndPathes(JSONFilesName []string) ([][]byte, []string) {
	var sortingJSON [][]byte
	var filePathes []string

	for _, file := range JSONFilesName {
		path := filepath.Join(flags.Path, file)
		filePathes = append(filePathes, path)
		initialFile := ReadJSONFile(path)
		outputFile := OutputJSON(initialFile)
		sortingJSON = append(sortingJSON, outputFile)
	}

	return sortingJSON, filePathes
}

// FindJSONInFolder return JSON files in current folder
func FindJSONInFolder() []string {
	folderFiles, _ := ioutil.ReadDir(flags.Path)

	var JSONFiles []string
	for _, file := range folderFiles {
		if filepath.Ext(file.Name()) == ".json" {
			JSONFiles = append(JSONFiles, file.Name())
		}
	}

	return JSONFiles
}

// ReadJSONFile return byte file of json
func ReadJSONFile(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

// OutputJSON returns sort JSON file
func OutputJSON(file []byte) []byte {
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

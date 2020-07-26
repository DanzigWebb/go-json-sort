package folder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"sorting/flags"
)

type JSONAndPath struct {
	Path    string
	Content []byte
}

// GetJSONAndPath returns output JSON files and pathes to original files in slices
func GetJSONAndPath(JSONFilesName []string) []JSONAndPath {
	var output []JSONAndPath

	for _, file := range JSONFilesName {
		path := filepath.Join(flags.Path, file)
		outputFile := OutputJSON(path)
		output = append(output, JSONAndPath{path, outputFile})
	}

	return output
}

// FindJSONFiles return JSON files in current folder
func FindJSONFiles() []string {
	folderFiles, _ := ioutil.ReadDir(flags.Path)

	var JSONFiles []string
	for _, file := range folderFiles {
		if filepath.Ext(file.Name()) == ".json" {
			JSONFiles = append(JSONFiles, file.Name())
		}
	}

	return JSONFiles
}

// OutputJSON returns sort JSON file
func OutputJSON(path string) []byte {
	var file = readJSONFile(path)
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

func readJSONFile(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

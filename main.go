package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "i18n", "path to folder")
}

func main() {
	flag.Parse()

	folder := configPath
	langs := findJSONInFolder(folder)

	for _, langFile := range langs {
		path := filepath.Join(folder, langFile)
		initialFile := readJSONFile(path)
		outputFile := outputJSON(initialFile)
		ioutil.WriteFile(path, outputFile, 0)
	}

	fmt.Printf("%d файла/ов отсортировано", len(langs))
}

func findJSONInFolder(pathToFolder string) []string {
	fileInfo, _ := ioutil.ReadDir(pathToFolder)

	var JSONFiles []string
	for _, file := range fileInfo {
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

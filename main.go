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
	pathToFolder string
)

func init() {
	flag.StringVar(&pathToFolder, "path", "i18n", "path to folder")
}

func main() {
	flag.Parse()

	langs := findJSONInFolder(pathToFolder)

	if len(langs) == 0 {
		fmt.Printf("Файлы с расширением .json не найдены \n")
		return
	}

	for _, file := range langs {
		path := filepath.Join(pathToFolder, file)
		initialFile := readJSONFile(path)
		outputFile := outputJSON(initialFile)
		ioutil.WriteFile(path, outputFile, 0777)
	}

	fmt.Printf("%d файла/ов отсортировано \n", len(langs))
}

func findJSONInFolder(pathToFolder string) []string {
	folderFiles, _ := ioutil.ReadDir(pathToFolder)

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

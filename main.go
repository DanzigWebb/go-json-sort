package main

import (
	"fmt"
	"io/ioutil"
	"sorting/compare"
	"sorting/flags"
	"sorting/folder"
)

func main() {
	flags.Parse()

	JSONNames := folder.FindJSONFiles()

	if len(JSONNames) == 0 {
		fmt.Println("Файлы json не найдены")
		return
	}

	outputFiles := folder.GetJSONAndPath(JSONNames)

	compare.Start(outputFiles)

	for _, file := range outputFiles {
		if err := ioutil.WriteFile(file.Path, file.Content, 0777); err != nil {
			fmt.Println("An error occurred:", err)
			return
		}
	}

	fmt.Printf("\n%d файла/ов отсортировано \n\n", len(JSONNames))
}

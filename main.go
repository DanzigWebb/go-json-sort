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
		fmt.Printf("Файлы json не найдены \n")
		return
	}

	sortingJSONSlice, filePathes := folder.GetJSONAndPath(JSONNames)

	compare.Start(sortingJSONSlice)

	for i, file := range sortingJSONSlice {
		ioutil.WriteFile(filePathes[i], file, 0777)
	}

	fmt.Printf("\n%d файла/ов отсортировано \n\n", len(JSONNames))
}

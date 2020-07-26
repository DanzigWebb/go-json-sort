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

	sortingJSONSlice, filePathSlice := folder.GetJSONAndPath(JSONNames)

	compare.Start(sortingJSONSlice)

	for i, file := range sortingJSONSlice {
		if err := ioutil.WriteFile(filePathSlice[i], file, 0777); err != nil {
			fmt.Println("An error occurred:", err)
			return
		}
	}

	fmt.Printf("\n%d файла/ов отсортировано \n\n", len(JSONNames))
}

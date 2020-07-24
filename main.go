package main

import (
	"fmt"
	"io/ioutil"
	"sorting/compare"
	"sorting/flags"
	"sorting/folder"
)

func main() {
	flags.Init()

	JSONFilesName := folder.FindJSONInFolder()

	if len(JSONFilesName) == 0 {
		fmt.Printf("Файлы json не найдены \n")
		return
	}

	sortingJSONSlice, filePathes := folder.GetJSONAndPathes(JSONFilesName)

	compare.Start(sortingJSONSlice)

	for i, file := range sortingJSONSlice {
		ioutil.WriteFile(filePathes[i], file, 0777)
	}

	fmt.Printf("\n%d файла/ов отсортировано \n\n", len(JSONFilesName))
}

package main

import (
	"fmt"
	"io/ioutil"
	"sorting/folder"
	"sorting/folderpath"
)

func main() {
	folderpath.Init()

	langs := folder.FindJSONInFolder()

	if len(langs) == 0 {
		fmt.Printf("Файлы с расширением .json не найдены \n")
		return
	}

	sortingJSON, fileNames := folder.GetJSONAndPathes(langs)

	for i, file := range sortingJSON {
		ioutil.WriteFile(fileNames[i], file, 0777)
	}

	fmt.Printf("%d файла/ов отсортировано \n", len(langs))
}

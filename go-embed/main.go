package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed test/version.txt
var version string

//go:embed test/img.png
var logo []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)
	err := ioutil.WriteFile("img_baru.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	dirEntries, _ := path.ReadDir("test/files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println(string(file))
		}
	}

}

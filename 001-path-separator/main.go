package main

import (
	"fmt"
	"path"
)

func main() {
	// var dir, file string
	// var file string

	dir, file := path.Split("css/main.css") // short version
	_, file = path.Split("css/main.css")

	fmt.Println("dir :", dir)
	fmt.Println("file :", file)
}

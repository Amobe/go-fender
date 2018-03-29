package main

import (
	"fender"
	"fender/convertor"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	file := argsWithoutProg[0]
	if file != "" {
		fender.FileHandler(file, convertor.Conv2Dos)
	}
}

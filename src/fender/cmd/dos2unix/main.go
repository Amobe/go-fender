package main

import (
	"bufio"
	"fender"
	"fender/convertor"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	file := argsWithoutProg[0]
	if file != "" {
		fileList, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(fileList)
		for scanner.Scan() {
			file := scanner.Text()
			fmt.Println(file)
			fender.FileHandler(file, convertor.Conv2Unix)
		}
	}
}

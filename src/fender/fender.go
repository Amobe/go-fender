package fender

import (
	"io"
	"io/ioutil"
	"os"
)

func isFileExist(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return false
}

func removeFile(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func backupFile(originPath string, backupPath string) {
	fileInfo, err := os.Stat(originPath)
	if err != nil {
		panic(err)
	}

	origin, err := os.Open(originPath)
	if err != nil {
		panic(err)
	}
	defer origin.Close()

	backup, err := os.OpenFile(backupPath, os.O_RDWR|os.O_CREATE, fileInfo.Mode())
	if err != nil {
		panic(err)
	}
	defer backup.Close()

	_, err = io.Copy(backup, origin)
	if err != nil {
		panic(err)
	}
}

func convertFile(src string, des string, convFunc func([]byte) []byte) {
	fileInfo, err := os.Stat(src)
	if err != nil {
		panic(err)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	data, _ := ioutil.ReadAll(srcFile)
	newData := convFunc(data)
	ioutil.WriteFile(des, newData, fileInfo.Mode())
}

func FileHandler(originPath string, convFunc func([]byte) []byte) {
	backupPath := originPath + ".bk"

	if !isFileExist(originPath) {
		return
	}

	backupFile(originPath, backupPath)
	removeFile(originPath)
	convertFile(backupPath, originPath, convFunc)
	removeFile(backupPath)
}

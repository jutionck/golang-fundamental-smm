package data

import (
	"fmt"
	"golang-fundamental/file/utils"
	"os"
)

func WriteToFile(fileLocation string, input string) {
	//os
	//ioutil

	outputFile, outputError := os.OpenFile(fileLocation, os.O_APPEND|os.O_RDWR, 0644)
	if outputError != nil {
		fmt.Println(outputError.Error())
		return
	}

	defer outputFile.Close()
	_, err := outputFile.WriteString(input + "\n")
	if utils.IsError(err) {
		return
	}

	err = outputFile.Sync()
	if utils.IsError(err) {
		return
	}
}

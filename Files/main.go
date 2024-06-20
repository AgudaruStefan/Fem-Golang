package main

import (
	"fmt"
	"learningo.com/go/files/fileutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}
}

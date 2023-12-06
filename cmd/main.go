package main

import (
	"os"
	"strings"
)

func doesFileExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func validateFile(fileName string) bool {
	linuxBlackList := [10]rune{'/', '>', '<', '|', ':', '&'}

	for i := 0; i < len(fileName); i++ {
		if strings.ContainsRune(fileName, linuxBlackList[i]) {
			return false
		}
	}

	return true
}

func rename(oldPath string, newPath string) {
	if !validateFile(newPath) || !validateFile(newPath) {
		panic("This type of character is not allowed when naming a file")
	}
	if !doesFileExist(oldPath) {
		panic("File does not exits")
	}

	renameFile := os.Rename(oldPath, newPath)
	if renameFile != nil {
		panic("Error renaming file please try again")
	}
}

func main() {
	args := os.Args

	if len(args) < 3 {
		return
	}
	rename(os.Args[1], os.Args[2])
}

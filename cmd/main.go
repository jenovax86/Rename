package main

import (
	"fmt"
	"os"
	"strings"
)

func doesFileExist(fileName string) {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		fmt.Printf("%v file does not exits\n", fileName)
	}
	return
}

func validateFile(fileName string) bool {
	linuxBlackList := [10]rune{'/', '>', '<', '|', ':', '&'}

	for i := 0; i < len(fileName); i++ {
		if strings.ContainsRune(fileName, linuxBlackList[i]) {
			return true
		}
	}

	return false
}

func rename(oldPath string, newPath string) {
	if validateFile(newPath) == true {
		panic("This type of character is not allowed when naming a file")
	}
	doesFileExist(oldPath)

	renameFile := os.Rename(oldPath, newPath)
	if renameFile != nil {
		panic("Error renaming file please try again")
	}
}

func main() {
	rename(os.Args[1], os.Args[2])
}


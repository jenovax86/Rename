package main

import (
	"bufio"
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

func rename() {
	fmt.Println("Enter The filename You want to rename and preferred name")
	read := bufio.NewReader(os.Stdin)

	input, err := read.ReadString('\n')

	if err != nil {
		panic(err)
	}

	input = strings.TrimSuffix(input, "\n")
	words := strings.Fields(input)
	if validateFile(words[1]) == true {
		panic("This type of character is not allowed when naming a file")
	}
	doesFileExist(words[0])

	renameFile := os.Rename(words[0], words[1])
	if renameFile != nil {
		panic("Error renaming file please try again")
	}
}

func main() {
	rename()
}

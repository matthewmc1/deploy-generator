package internal

import (
	"fmt"
	"os"

	"github.com/h2non/filetype"
)

func parseConfig() {

}

func validateConfig() {

}

func checkFileType(file string) {
	check, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}

	buff := make([]byte, 512)
	_, err = check.Read(buff)
	if err != nil {
		fmt.Println("Error reading content to determine filetype")
		os.Exit(1)
	}

	kind, err := filetype.Match(buff)
	if err != nil {
		fmt.Println("Error reading content to determine filetype")
		os.Exit(1)
	}
	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}

func Generator(file string) {
	checkFileType(file)
}

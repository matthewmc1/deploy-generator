package internal

import (
	"fmt"
	"os"
	"validator/service-definition/data"

	"github.com/h2non/filetype"
)

func parseConfig() {

}

func validateConfig() {

}

func checkFileType(file string) string {
	check, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
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
		fmt.Printf("Error reading content to determine filetype %v", kind)
		os.Exit(1)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading data from file")
		os.Exit(1)
	}

	fmt.Println(string(data))

	return string(data)
}

func Generator(file string) {
	d := checkFileType(file)
	service := data.Decode(d)

	fmt.Println(service.Labels)
}

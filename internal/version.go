package internal

import (
	"fmt"
	"io"
	"os"
)

func Version() string {
	file, err := os.Open("VERSION")
	if err != nil {
		fmt.Println("Error reading version file:", err)
		os.Exit(1)
	}
	defer file.Close() // Ensure file closure

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading version file:", err)
		os.Exit(1)
	}

	return string(data)
}

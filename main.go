package main

import (
	"fmt"
	"os"
	"validator/service-definition/cmd"
)

func main() {
	if err := cmd.BaseCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

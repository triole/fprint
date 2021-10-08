package main

import (
	"fmt"
	"os"

	figure "github.com/triole/go-figure"
)

func main() {
	stdin := getStdin()
	parseArgs()

	var msgArr []string
	if CLI.Message == "" {
		msgArr = stdin
	} else {
		msgArr = []string{CLI.Message}
	}

	if CLI.ListFonts == true {
		figure.PrintFontList()
		os.Exit(0)
	}

	msg := initConfig(CLI.Config, msgArr)

	if CLI.PrintExamples == true {
		fontList := figure.GetFontList()
		for _, font := range fontList {
			fmt.Printf("\n\nFont %q\n", font)
			msg.Fnt = font
			printMessage(msg)
		}
	} else {
		printMessage(msg)
	}
}

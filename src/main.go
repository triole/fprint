package main

import (
	"fmt"
	"os"

	figure "github.com/triole/go-figure"
)

func main() {
	stdin := getStdin()
	argparse()

	var msgArr []string
	if *argsMessage == "" {
		msgArr = stdin
	} else {
		msgArr = []string{*argsMessage}
	}

	if *argsList == true {
		figure.PrintFontList()
		os.Exit(0)
	}

	msg := initConfig(*argsConfig, msgArr)

	if *argsExample == true {
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

package main

import (
	"fmt"

	figure "github.com/triole/go-figure"
)

func main() {
	argparse()

	if *argsList == true {
		figure.PrintFontList()
	} else if *argsExample == true {
		fontList := figure.GetFontList()
		for _, font := range fontList {
			fmt.Printf("\n\nFont %q\n", font)
			cprint(*argsMessage, font, *argsColour)
		}
	} else {
		cprint(*argsMessage, *argsFont, *argsColour)
	}
}

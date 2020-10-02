package main

import (
	figure "github.com/triole/go-figure"
)

func aprint(msg string, font string) {
	fig := figure.NewFigure(msg, font, true)
	fig.Print()
	println("")
}

func cprint(msg string, font string, colour string) {
	fig := figure.NewColorFigure(msg, font, colour, true)
	fig.Print()
	println("")
}

func listFonts() {

}

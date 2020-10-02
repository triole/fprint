package main

import (
	"fmt"

	figure "github.com/triole/go-figure"
)

func printMessage(msg tMessage) {
	printIf(msg.Pre)
	fig := figure.NewColorFigure(msg.Txt, msg.Fnt, msg.Col, true)
	fig.Print()
	printIf(msg.Post)
	println("")
}

func printIf(s string) {
	if s != "" {
		fmt.Printf("\n%s\n\n", s)
	}
}

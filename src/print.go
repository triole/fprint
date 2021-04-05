package main

import (
	"fmt"

	figure "github.com/triole/go-figure"
)

func printMessage(msg tConf) {
	printIf(msg.Pre)
	for _, el := range msg.Txt {
		fig := figure.NewColorFigure(el, msg.Fnt, msg.Col, true)
		fmt.Printf("\n")
		fig.Print()
	}
	printIf(msg.Post)
	fmt.Printf("\n\n")
}

func printIf(s string) {
	if s != "" {
		fmt.Printf("\n%s\n", s)
	}
}

package main

import (
	figure "github.com/triole/go-figure"
)

func aprint(msg tMessage) {
	fig := figure.NewFigure(msg.Txt, msg.Fnt, true)
	fig.Print()
	println("")
}

func cprint(msg tMessage) {
	fig := figure.NewColorFigure(msg.Txt, msg.Fnt, msg.Col, true)
	fig.Print()
	println("")
}

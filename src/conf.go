package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func initConfig(fil string) (msg tMessage) {
	if fil != "" {
		msg = readConfigFile(fil)
	}
	msg = pushArgsIntoConfig(msg)
	return
}

func readConfigFile(fil string) (msg tMessage) {
	content, err := ioutil.ReadFile(fil)
	checkx(err, "Unable to read file")

	_, err = toml.Decode(string(content), &msg)
	checkx(err, "Unable to decode")
	return
}

func pushArgsIntoConfig(msg tMessage) tMessage {
	if *argsMessage != "" {
		msg.Txt = *argsMessage
	}
	if *argsColour != "" {
		msg.Col = *argsColour
	}
	if *argsFont != "" {
		msg.Fnt = *argsFont
	}
	if *argsPreText != "" {
		msg.Pre = *argsPreText
	}
	if *argsPostText != "" {
		msg.Post = *argsPostText
	}

	// apply default values if anything is missing
	if msg.Txt == "" {
		msg.Txt = "sphinx of black quartz"
	}
	if msg.Col == "" {
		msg.Col = "white"
	}
	if msg.Fnt == "" {
		msg.Fnt = "big"
	}
	return msg
}

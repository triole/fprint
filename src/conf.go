package main

import (
	"fmt"
	"io/ioutil"
	"os"

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
	if err != nil {
		fmt.Printf("Unable to read file. %s\n", err)
		os.Exit(1)
	}
	if _, err := toml.Decode(string(content), &msg); err != nil {
		fmt.Printf("Unable to decode. %s\n", err)
		os.Exit(1)
	}
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
		msg.Txt = *argsPreText
	}
	if *argsPostText != "" {
		msg.Txt = *argsPostText
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

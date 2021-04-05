package main

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	defaultMsg = "sphinx of black quartz"
)

func setVal(itf interface{}, alt string, def interface{}) (r string) {
	switch v := itf.(type) {
	case string:
		r = v
	}
	if r == "" && alt != "" {
		r = alt
	}
	if r == "" && def != nil {
		r = def.(string)
	}
	return
}

func initConfig(fil string, msgArr []string) (conf tConf) {
	var ctoml tConfToml

	if _, err := os.Stat(fil); err == nil {
		content, err := ioutil.ReadFile(fil)
		checkx(err, "Unable to read file "+fil)
		_, err = toml.Decode(string(content), &ctoml)
		checkx(err, "Unable to decode "+fil)
	}

	conf.Col = setVal(ctoml.Col, *argsColour, "white")
	conf.Fnt = setVal(ctoml.Fnt, *argsFont, "big")
	conf.Pre = setVal(ctoml.Pre, *argsPreText, nil)
	conf.Post = setVal(ctoml.Post, *argsPostText, nil)

	if ctoml.Txt != nil {
		for _, el := range ctoml.Txt {
			conf.Txt = append(conf.Txt, el.(string))
		}
	}
	if len(msgArr) != 0 {
		conf.Txt = msgArr
	}
	if len(conf.Txt) == 0 {
		conf.Txt = []string{defaultMsg}
	}

	return
}

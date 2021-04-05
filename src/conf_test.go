package main

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	AssertReadConfig("../test/all.toml", []string{"hello world", "red", "big"}, t)
	AssertReadConfig("../test/no_font.toml", []string{"hello world", "red", "big"}, t)
	AssertReadConfig("../test/only_col.toml", []string{defaultMsg, "red", "big"}, t)
	AssertReadConfig("../test/only_text.toml", []string{"hello world", "white", "big"}, t)
	AssertReadConfig("../test/empty.toml", []string{defaultMsg, "white", "big"}, t)
}

func AssertReadConfig(fil string, arr []string, t *testing.T) {
	asr := tConf{
		Txt: []string{arr[0]},
		Col: arr[1],
		Fnt: arr[2],
	}
	msg := initConfig(fil, []string{})
	if msg.Txt[0] != asr.Txt[0] || msg.Col != asr.Col || msg.Fnt != asr.Fnt {
		t.Errorf("Message assert failed for %q: %q != %q", fil, msg, asr)
	}
}

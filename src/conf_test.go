package main

import "testing"

func TestReadConfig(t *testing.T) {
	AssertReadConfig("../test/all.toml", []string{"hello world", "red", "big"}, t)
	AssertReadConfig("../test/no_font.toml", []string{"hello world", "red", ""}, t)
	AssertReadConfig("../test/only_col.toml", []string{"", "red", ""}, t)
	AssertReadConfig("../test/only_text.toml", []string{"hello world", "", ""}, t)
}

func AssertReadConfig(fil string, arr []string, t *testing.T) {
	asr := tMessage{
		Txt: arr[0],
		Col: arr[1],
		Fnt: arr[2],
	}
	msg := readConfigFile(fil)
	if msg.Txt != asr.Txt || msg.Col != asr.Col || msg.Fnt != asr.Fnt {
		t.Errorf("Message assert failed for %q: %q != %q", fil, msg, asr)
	}
}

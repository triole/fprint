package main

type tMessage struct {
	Txt  string `toml:"text"`
	Col  string `toml:"colour"`
	Fnt  string `toml:"font"`
	Pre  string `toml:"pretext"`
	Post string `toml:"posttext"`
}

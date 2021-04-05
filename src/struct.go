package main

type tConfToml struct {
	Txt  []interface{} `toml:"text"`
	Col  interface{}   `toml:"colour"`
	Fnt  interface{}   `toml:"font"`
	Pre  interface{}   `toml:"pretext"`
	Post interface{}   `toml:"posttext"`
}

type tConf struct {
	Txt  []string
	Col  string
	Fnt  string
	Pre  string
	Post string
}

package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	BUILDTAGS      string
	appName        = "fprint"
	appMainVersion = "0.1"
	appDescription = "figlet print"

	app          = kingpin.New(appName, appDescription)
	argsMessage  = app.Arg("message", "message to print").String()
	argsColour   = app.Flag("colour", "text color").Short('c').String()
	argsFont     = app.Flag("font", "font").Short('f').String()
	argsPreText  = app.Flag("pre text", "text before the message").Short('p').String()
	argsPostText = app.Flag("post text", "text after the message").Short('s').String()
	argsList     = app.Flag("list", "list available fonts").Short('l').Default("false").Bool()
	argsExample  = app.Flag("example", "print message in all fonts available").Short('x').Default("false").Bool()
	argsConfig   = app.Flag("config", "load settings from config").Short('k').Default("").String()
)

func argparse() {
	env := tEnv{
		Name:        appName,
		MainVersion: appMainVersion,
		Description: appDescription,
	}
	app.Version(makeInfoString(env, parseBuildtags(BUILDTAGS)))
	app.VersionFlag.Short('V')
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

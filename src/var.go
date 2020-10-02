package main

import (
	"olibs/environment"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	BUILDTAGS      string
	appName        = "aprint"
	appMainversion = "0.1"
	appDescription = "Ascii print"
	env            = environment.Init(appName, appMainversion, appDescription, BUILDTAGS)

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
	// argparse
	app.Version(env.AppInfoString)
	app.VersionFlag.Short('V')
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

package main

import (
	"runtime"
	"sort"
	"strings"
)

type tEnv struct {
	Name        string
	MainVersion string
	Subversion  string
	Description string
	Author      string
	Githash     string
	Builddate   string
	Loc         string
}

func makeInfo(name string, mversion string, desc string) tEnv {
	return tEnv{
		Name:        name,
		MainVersion: mversion,
		Description: desc,
	}
}

func parseBuildtags(buildtags string) (m map[string]string) {
	m = make(map[string]string)
	arr := strings.Split(buildtags, ",")
	for _, s := range arr {
		s = strings.TrimLeft(s, "{")
		s = strings.TrimRight(s, "}")
		s = strings.TrimSpace(s)
		p := strings.Split(s, ":")
		// join to keep the ':' of the time stamp in build date
		key := strings.TrimSpace(p[0])
		val := strings.TrimSpace(strings.Join(p[1:], ":"))
		m[key] = val
	}
	return
}

func makeInfoString(env tEnv, buildtags map[string]string) (s string) {
	var sorter []string
	// detect maximum length of key strings
	maxlength := len(runtime.Version()) + 2
	for k := range buildtags {
		if len(k) > maxlength {
			maxlength = len(k)
		}
		sorter = append(sorter, k)
	}
	// manual part of string assembly
	s += "\n" + strings.Title(env.Name) + "\n"
	s += env.Description + "\n\n"
	s += makeEntryKey("version", maxlength) + env.MainVersion + "." + buildtags["_subversion"] + "\n"
	s += makeEntryKey("go runtime", maxlength) + runtime.Version() + "\n"
	s += makeEntryKey("", maxlength) + "\n"
	// automatic part
	sort.Strings(sorter)
	for _, k := range sorter {
		if strings.HasPrefix(k, "_") == false {
			s += makeEntryKey(k, maxlength) + buildtags[k] + "\n"
		}
	}
	return
}

func makeEntryKey(s string, i int) (r string) {
	if s != "" {
		s += ":"
	}
	for len(s) <= i+1 {
		s += " "
	}
	r = strings.Title(s)
	return
}

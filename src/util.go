package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkp(err error) {
	if err != nil {
		panic(err)
	}
}

func checkx(err error, msg string) {
	if err != nil {
		fmt.Printf("%s. %s\n", msg, err)
		os.Exit(1)
	}
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getStdin() (s string) {
	if isInputFromPipe() == true {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanner.Scan()
			text := scanner.Text()
			if len(text) != 0 {
				s = text
			} else {
				break
			}
		}
	}
	return
}

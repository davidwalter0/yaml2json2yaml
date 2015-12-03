package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var version = flag.Bool("version", false, "print build and git commit as a version string")

var Build string  // from the build ldflag options
var Commit string // from the build ldflag options

func init() {
	flag.Parse()

	if *version {
		array := strings.Split(os.Args[0], "/")
		me := array[len(array)-1]
		fmt.Println(me, "Build:", Build, "Commit:", Commit)
		os.Exit(0)
	}
}

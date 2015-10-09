/*
  yaml2json: use as a filter or with a file argument. writes to stdout
  refactored with transform and added json filter
*/

package main

import (
	"flag"
	"fmt"
	"github.com/davidwalter0/transform"
	"io/ioutil"
	"os"
)

func main() {
	err := _main()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	os.Exit(0)
}

var compress = flag.Bool("compress", false, "--compress alias for unformat.")
var indent = flag.Int("indent", 2, "--indent n, where n is the number of spaces to indent each level.")
var file = flag.String("file", "", "--file=json file to convert to yaml")

func _main() error {
	flag.Parse()

	var data []byte
	var input []byte
	var err error
	if len(*file) > 0 {
		input, err = ioutil.ReadFile(*file)
		if err != nil {
			fmt.Println(err, "error reading the file argument given")
			os.Exit(1)
		}
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
	}
	if *compress {
		data, err = transform.Yaml2Json(input)
	} else {
		data, err = transform.Yaml2JsonIndent(input)
	}
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(data)
	return err
}

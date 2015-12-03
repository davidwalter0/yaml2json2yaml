/*
  json2yaml: use as a filter or with a file argument. writes to stdout
*/

package main

import (
	"flag"
	"fmt"
	"github.com/davidwalter0/transform"
	"io/ioutil"
	"os"
)

var file = flag.String("file", "", "--file=yaml source file to convert to json")

func main() {
	err := _main()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	os.Exit(0)
}

func _main() error {
	var err error
	var data []byte
	var input []byte

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

	data, err = transform.Json2Yaml(input)
	_, err = os.Stdout.Write(data)
	return err
}

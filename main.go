package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/b4b4r07/go-colon"
	"github.com/b4b4r07/go-filter"
	"github.com/mattn/go-pipeline"
)

func main2() {
	result, err := colon.Parse(os.Getenv("ENHANCD_FILTER"))
	if err != nil {
		panic(err)
	}
	filter, err := result.Executable().First()
	if err != nil {
		panic(err)
	}

	logfile := filepath.Join(os.Getenv("ENHANCD_DIR"), "enhancd.log")
	out, err := pipeline.Output(
		[]string{"cat", logfile},
		filter.Args,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(out))
}

func main() {
	logfile := filepath.Join(os.Getenv("ENHANCD_DIR"), "enhancd.log")
	contents, err := ioutil.ReadFile(logfile)
	if err != nil {
		panic(err)
	}
	text := string(contents)
	lines, err := filter.Run(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", lines)
}

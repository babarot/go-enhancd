package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/b4b4r07/go-colon"
	"github.com/mattn/go-pipeline"
)

func main() {
	result, err := colon.Parse(os.Getenv("ENHANCD_FILTER"))
	if err != nil {
		panic(err)
	}
	logfile := filepath.Join(os.Getenv("ENHANCD_DIR"), "enhancd.log")
	filter := result.Executable().One()
	out, err := pipeline.Output(
		[]string{"cat", logfile},
		filter.Attr.Args,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(out))
}

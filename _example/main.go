package main

import (
	"fmt"

	"github.com/b4b4r07/go-enhancd"
)

func main() {
	history, err := enhancd.NewHistory()
	if err != nil {
		panic(err)
	}
	lines, err := history.Select()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", lines)
	fmt.Printf("%#v\n", lines.Reverse())
}

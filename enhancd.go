package enhancd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/b4b4r07/go-filter"
)

type Enhancd interface {
	Select() ([]string, error)
}

type Lines []string

type History struct {
	File   string
	Entire []byte
	Lines  Lines
}

func NewHistory() (*History, error) {
	dir := os.Getenv("ENHANCD_DIR")
	if dir == "" {
		return &History{}, errors.New("ENHANCD_DIR: not set")
	}
	file := filepath.Join(dir, "enhancd.log")
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return &History{}, err
	}
	var lines []string
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return &History{
		File:   file,
		Entire: contents,
		Lines:  lines,
	}, nil
}

func (h *History) Select() (Lines, error) {
	lines, err := filter.Run(string(h.Entire))
	if err != nil {
		return Lines{}, err
	}
	return lines, nil
}

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Lines) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l *Lines) Reverse() *Lines {
	sort.Sort(sort.Reverse(l))
	return l
}

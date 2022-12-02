package io

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(fname string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		panic(err)
	}

	return string(b)
}

func Read1DString(fname, separator string, skipLast bool) []string {
	s := ReadFile(fname)
	rawLines := strings.Split(s, separator)

	if skipLast {
		return rawLines[:len(rawLines)-1]
	}
	return rawLines
}

func Read2DString(fname, separator, separator2 string, skipLast bool) [][]string {
	rawLines := Read1DString(fname, separator, skipLast)
	lines := make([][]string, len(rawLines))
	for i, l := range rawLines {
		lines[i] = strings.Split(strings.TrimSpace(l), separator2)
	}
	return lines
}

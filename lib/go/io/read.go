package io

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/khinshankhan/advent/lib/go/util"
)

func ReadFile(pathParts ...string) string {
	p := filepath.Join(pathParts...)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func ReadRelativeFile(fname string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return ReadFile(dir, fname)
}

func Read1DString(fname, separator string, skipLast bool) []string {
	return util.TransformString1D(
		ReadRelativeFile(fname),
		separator,
		skipLast,
	)
}

func Read2DString(fname, separator, separator2 string, skipLast bool) [][]string {
	return util.TransformString2D(
		ReadRelativeFile(fname),
		separator,
		separator2,
		skipLast,
	)
}

func ReadTestFile(pathParts ...string) string {
	p := filepath.Join(pathParts...)
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return ""
	}
	return ReadFile(p)
}

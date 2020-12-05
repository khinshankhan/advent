package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadRelativeFile1DString will return an array of strings from a given file separated by separator
func ReadRelativeFile1DString(fname, separator string) ([]string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		return nil, err
	}

	rawLines := strings.Split(string(b), separator)
	lines := rawLines[:len(rawLines)-1]
	return lines, nil
}

// ReadRelativeFile2DString will return an array of strings from a given file separated by separator
func ReadRelativeFile2DString(fname, separator, separator2 string) ([][]string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		return nil, err
	}

	rawLines := strings.Split(string(b), separator)
	rawLines = rawLines[:len(rawLines)-1]

	lines := make([][]string, len(rawLines))
	for i, l := range rawLines {
		lines[i] = strings.Split(l, separator2)
	}
	return lines, nil
}

// ReadRelativeFile1DInt will return an array of ints from a given file separated by separator
func ReadRelativeFile1DInt(fname, separator string) ([]int, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		return nil, err
	}

	rawLines := strings.Split(string(b), separator)
	lines := rawLines[:len(rawLines)-1]

	nums := make([]int, len(lines))
	for i, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

// ReadRelativeFile2DInt will return an array of ints from a given file separated by separator
func ReadRelativeFile2DInt(fname, separator string) ([]int, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	b, err := ioutil.ReadFile(filepath.Join(dir, fname))
	if err != nil {
		return nil, err
	}

	rawLines := strings.Split(string(b), separator)
	lines := rawLines[:len(rawLines)-1]

	nums := make([]int, len(lines))
	for i, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

// Max returns max of two ints
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Min returns min of two ints
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

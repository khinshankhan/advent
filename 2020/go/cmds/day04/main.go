package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	lines, err := customReadRelativeFile("../data/day04.txt", "\n\n")
	if err != nil {
		panic(err)
	}

	validAttributes := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}
	parta, partb := 0, 0
	for _, passport := range lines {
		passport = strings.ReplaceAll(passport, "\n", " ")
		passportAttributes := strings.Split(passport, " ")
		keyCountA, keyCountB := 0, 0
		for _, passportAttribute := range passportAttributes {
			passportKey := strings.Split(passportAttribute, ":")
			if validAttributes[passportKey[0]] {
				keyCountA += 1
				if verifyKey(passportKey[0], passportKey[1]) {
					keyCountB += 1
				}
			}
		}

		if keyCountA == len(validAttributes) {
			parta += 1
		}
		if keyCountB == len(validAttributes) {
			partb += 1
		}
	}
	fmt.Println(parta)
	fmt.Println(partb)
}

func verifyKey(attributeKey, attributeValue string) bool {
	switch attributeKey {
	case "byr":
		return len(attributeValue) == 4 && attributeValue >= "1920" && attributeValue <= "2002"
	case "iyr":
		return len(attributeValue) == 4 && attributeValue >= "2010" && attributeValue <= "2020"
	case "eyr":
		return len(attributeValue) == 4 && attributeValue >= "2020" && attributeValue <= "2030"
	case "hgt":
		if strings.HasSuffix(attributeValue, "cm") {
			return attributeValue[:len(attributeValue)-2] >= "150" && attributeValue[:len(attributeValue)-2] <= "193"
		} else if strings.HasSuffix(attributeValue, "in") {
			return attributeValue[:len(attributeValue)-2] >= "59" && attributeValue[:len(attributeValue)-2] <= "76"
		}
		return false
	case "hcl":
		val, _ := regexp.MatchString("^#[A-Za-z0-9]{6}$", attributeValue)
		return val
	case "ecl":
		switch attributeValue {
		case
			"amb",
			"blu",
			"brn",
			"gry",
			"grn",
			"hzl",
			"oth":
			return true
		}
		return false
	case "pid":
		val, _ := regexp.MatchString("^[0-9]{9}$", attributeValue)
		return val
	default:
		fmt.Println("debug: invalid key", attributeKey)
		return false
	}
}

func customReadRelativeFile(fname, separator string) ([]string, error) {
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
	return rawLines, nil
}

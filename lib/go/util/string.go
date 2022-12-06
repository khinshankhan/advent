package util

import "strings"

func TransformString1D(s, separator string, skipLast bool) []string {
	rawLines := strings.Split(s, separator)

	if skipLast {
		return rawLines[:len(rawLines)-1]
	}
	return rawLines
}

func TransformString2D(s, separator, separator2 string, skipLast bool) [][]string {
	rawLines := TransformString1D(s, separator, skipLast)
	lines := make([][]string, len(rawLines))
	for i, l := range rawLines {
		lines[i] = strings.Split(l, separator2)
	}
	return lines
}

package util

import (
	"strings"
)

func TransformString1D(s, separator string, trimSpace bool) []string {
	lines := strings.Split(s, separator)

	if !trimSpace {
		return lines
	}

	return cleanSpaces1D(lines)
}

func TransformString2D(
	s,
	separator,
	separator2 string,
	trimSpace,
	trimInnerSpace bool,
) [][]string {
	rawLines := TransformString1D(s, separator, trimSpace)

	matrix := make([][]string, len(rawLines))
	for i, rl := range rawLines {
		lines := strings.Split(rl, separator2)
		if trimInnerSpace {
			lines = cleanSpaces1D(lines)
		}
		matrix[i] = lines
	}
	return matrix
}

func cleanSpaces1D(lines []string) []string {
	cleanedLines := make([]string, len(lines))
	for i, e := range lines {
		cleanedLines[i] = strings.TrimSpace(e)
	}
	return cleanedLines
}

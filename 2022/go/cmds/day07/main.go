package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
	"github.com/khinshankhan/advent/lib/go/math"
	"github.com/khinshankhan/advent/lib/go/util"
)

func main() {
	s := io.ReadRelativeFile("../data/day07.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

type File struct {
	Name string
	Size int
}
type FsNode struct {
	Name          string
	Files         []File
	Dirs          []string
	RecursiveSize int
}

func SubDirName(p, c string) string {
	if p == "/" {
		p += c
	} else if c == "/" {
		p = "/"
	} else {
		p = p + "/" + c
	}
	return p
}

func parse(s string) map[string]FsNode {
	input := util.TransformString2D(strings.TrimSpace(s), "$ ", "\n", true, true)

	p := ""
	fs := make(map[string]FsNode)

	for _, lines := range input {
		if lines[0] == "" {
			continue
		}

		cmd := strings.Split(lines[0], " ")
		switch cmd[0] {
		case "cd":
			if cmd[1] == ".." {
				pParts := strings.Split(p, "/")
				p = strings.Join(pParts[:len(pParts)-1], "/")
			} else {
				p = SubDirName(p, cmd[1])
				if _, ok := fs[p]; !ok {
					fs[p] = FsNode{}
				}
			}
		case "ls":
			for _, file := range lines[1:] {
				info := strings.Split(file, " ")

				n := fs[p]
				if info[0] == "dir" {
					n.Dirs = append(n.Dirs, info[1])
				} else {
					sz := util.FromStringToInt(info[0])
					file := File{info[1], sz}
					n.Files = append(n.Files, file)
				}
				fs[p] = n
			}
		}
	}

	dirNames := []string{}
	for n := range fs {
		dirNames = append(dirNames, n)
	}

	// ensure 'innermost' dirs are listed first, which will help 'build up'
	// since they shouldn't have nested dirs, creating 'base cases'
	sort.Slice(dirNames, func(i, j int) bool {
		l1, l2 := len(dirNames[i]), len(dirNames[j])
		if l1 != l2 {
			return l1 > l2
		}
		return dirNames[i] > dirNames[j]
	})

	for _, dirName := range dirNames {
		rsz := 0
		for _, file := range fs[dirName].Files {
			rsz += file.Size
		}
		for _, dir := range fs[dirName].Dirs {
			file := fs[SubDirName(dirName, dir)]
			rsz += file.RecursiveSize
		}

		updatedDir := fs[dirName]
		updatedDir.RecursiveSize = rsz
		fs[dirName] = updatedDir
	}

	return fs
}

func parta(fs map[string]FsNode) int {
	sum := 0
	for _, dir := range fs {
		if dir.RecursiveSize <= 100000 {
			sum += dir.RecursiveSize
		}
	}

	return sum
}

func partb(fs map[string]FsNode) int {
	needed := (70000000 - 30000000 - fs["/"].RecursiveSize) * -1

	acceptable := []int{}
	for _, dir := range fs {
		if dir.RecursiveSize >= needed {
			acceptable = append(acceptable, dir.RecursiveSize)
		}
	}

	return math.Min(acceptable...)
}

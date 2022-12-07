package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khinshankhan/advent/lib/go/conv"
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
	Name  string
	Files []File
	Dirs  []string
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
					sz := conv.FromStringToInt(info[0])
					file := File{info[1], sz}
					n.Files = append(n.Files, file)
				}
				fs[p] = n
			}
		}
	}

	return fs
}

func getDirSizes(fs map[string]FsNode) map[string]File {
	dirs := []string{}
	for n := range fs {
		dirs = append(dirs, n)
	}

	sort.Slice(dirs, func(i, j int) bool {
		l1, l2 := len(dirs[i]), len(dirs[j])
		if l1 != l2 {
			return l1 > l2
		}
		return dirs[i] > dirs[j]
	})

	dirSizes := make(map[string]File)
	for _, key := range dirs {
		sz := 0
		for _, dir := range fs[key].Dirs {
			file := dirSizes[SubDirName(key, dir)]
			sz += file.Size
		}
		for _, file := range fs[key].Files {
			sz += file.Size
		}

		dirSizes[key] = File{key, sz}
	}

	return dirSizes
}

func parta(fs map[string]FsNode) int {
	dirSizes := getDirSizes(fs)

	sum := 0
	for _, dir := range dirSizes {
		if dir.Size <= 100000 {
			sum += dir.Size
		}
	}

	return sum
}

func partb(fs map[string]FsNode) int {
	szs := getDirSizes(fs)
	needed := (70000000 - 30000000 - szs["/"].Size) * -1

	acceptable := []int{}
	for _, dir := range szs {
		if dir.Size >= needed {
			acceptable = append(acceptable, dir.Size)
		}
	}

	return math.Min(acceptable...)
}

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
	s := io.ReadRelativeFile("../data/day11.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parta(monkeys Monkeys) int64 {
	worryDec := func(old int64) int64 { return old / 3 }
	return playRounds(monkeys, worryDec, 20)
}

func partb(monkeys Monkeys) int64 {
	getDivisor := func(m Monkey) int64 { return m.Divisor }
	d := math.ProductNums(util.Map1D(getDivisor, monkeys)...)
	worryDec := func(old int64) int64 { return old % d }
	return playRounds(monkeys, worryDec, 10000)
}

func playRounds(monkeys Monkeys, worryDec func(int64) int64, rounds int) int64 {
	m := make(Monkeys, len(monkeys))
	for i := range monkeys {
		m[i] = monkeys[i].Clone()
	}

	for ri := 0; ri < rounds; ri++ {
		m = m.Play(worryDec)
	}

	sort.Slice(m, func(i, j int) bool { return m[i].Inspected > m[j].Inspected })
	return m[0].Inspected * m[1].Inspected
}

type Monkeys []Monkey

func (mos *Monkeys) Play(worryDec func(int64) int64) Monkeys {
	m := *mos
	for i := 0; i < len(m); i++ {
		for _, item := range m[i].Items {
			worry := m[i].WorryInc(item)
			worry = worryDec(worry)
			next := m[i].Throw(worry)
			m[next].Items = append(m[next].Items, worry)
		}
		m[i].Items = []int64{}
	}
	return m
}

func (mos Monkeys) Clone() Monkeys {
	for i := range mos {
		mos[i] = mos[i].Clone()
	}
	return mos
}

type Monkey struct {
	Items     []int64
	WorryInc  func(int64) int64
	throw     func(int64) int
	Divisor   int64
	Inspected int64
}

func (mo *Monkey) Throw(worry int64) int {
	mo.Inspected += 1
	return mo.throw(worry)
}

// not true deep copy, but ops and throw shouldnt change
func (mo Monkey) Clone() Monkey {
	items := util.Clone1D(mo.Items)
	mo.Items = items
	return mo
}

func parse(s string) Monkeys {
	split := util.TransformString2D(strings.TrimSpace(s), "\n\n", "\n", true, true)
	ret := make(Monkeys, len(split))

	getNum := func(s string) int64 {
		if s == "old" {
			return -1
		}
		return int64(util.FromStringToInt(s))
	}

	getLastNum := func(s string) int {
		split := strings.Split(s, " ")
		return util.FromStringToInt(split[len(split)-1])
	}

	for i := range split {
		lines := split[i]

		strItems := util.TransformString1D(strings.Split(lines[1], ": ")[1], ", ", true)
		items := make([]int64, len(strItems))
		for i, e := range strItems {
			items[i] = int64(util.FromStringToInt(e))
		}

		exp := strings.Split(lines[2], " = ")[1]
		opParts := strings.Split(exp, " ")
		n1, n2, op := getNum(opParts[0]), getNum(opParts[2]), opParts[1]
		o1, o2 := opParts[0] == "old", opParts[2] == "old"

		d, t, f := getLastNum(lines[3]), getLastNum(lines[4]), getLastNum(lines[5])

		worryInc := func(old int64) int64 {
			if o1 {
				n1 = old
			}
			if o2 {
				n2 = old
			}

			switch op {
			// no - / since worry always increases
			case "+":
				return n1 + n2
			case "*":
				return n1 * n2
			default:
				panic(op)
			}
		}

		throw := func(worry int64) int {
			if worry%int64(d) == 0 {
				return t
			}
			return f
		}

		ret[i] = Monkey{items, worryInc, throw, int64(d), 0}
	}

	return ret
}

var sample = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

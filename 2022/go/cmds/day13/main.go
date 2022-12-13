package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/khinshankhan/advent/lib/go/io"
)

func main() {
	s := io.ReadRelativeFile("../data/day13.txt")
	input := parse(s)

	fmt.Println(parta(input))
	fmt.Println(partb(input))
}

func parta(input []Packet) int {
	sum := 0
	for i := 0; i < len(input); i += 2 {
		l, r := input[i], input[i+1]
		ordered, determined := comparePackets(l, r)
		if ordered && determined {
			sum += (i / 2) + 1
		}
	}
	return sum
}

func partb(input []Packet) int {
	two, six := 2, 6
	pack2, pack6 := []Packet{{num: &two}}, []Packet{{num: &six}}
	packets := append(input, Packet{packets: &pack2}, Packet{packets: &pack6})

	sort.Slice(packets, func(i, j int) bool {
		ordered, determined := comparePackets(packets[i], packets[j])
		return ordered && determined
	})

	product := 1
	for i, packet := range packets {
		if packet.packets != nil &&
			len(*packet.packets) == 1 &&
			((*packet.packets)[0].num == &two ||
				(*packet.packets)[0].num == &six) {
			product *= i + 1
		}
	}

	return product
}

func comparePackets(l, r Packet) (bool, bool) {
	switch {
	case l.num != nil && r.num != nil:
		if *l.num == *r.num {
			return false, false
		}
		return *l.num < *r.num, true
	case l.num != nil && r.packets != nil:
		newPacket := []Packet{{num: l.num}}
		ordered, determined := comparePackets(Packet{nil, &newPacket}, r)
		return ordered, determined
	case r.num != nil && l.packets != nil:
		newPacket := []Packet{{num: r.num}}
		ordered, determined := comparePackets(l, Packet{nil, &newPacket})
		return ordered, determined
	default: // both are lists
		for i := 0; i < len(*l.packets); i++ {
			if i == len(*r.packets) {
				return false, true
			}
			ordered, determined := comparePackets((*l.packets)[i], (*r.packets)[i])
			if determined {
				return ordered, determined
			}
		}
		greaterLength := len(*l.packets) < len(*r.packets) // l is out of items, but is r?
		return greaterLength, greaterLength
	}
}

type Packet struct {
	num     *int
	packets *[]Packet
}

func unmarshalPacket(raw interface{}) Packet {
	switch v := raw.(type) {
	case float64:
		num := int(v)
		return Packet{num: &num}
	case []interface{}:
		inner := make([]Packet, len(v))
		for i, item := range v {
			inner[i] = unmarshalPacket(item)
		}
		return Packet{packets: &inner}
	}
	fmt.Println(raw)
	panic("unknown")
}

func parse(s string) []Packet {
	lines := "[" + strings.Join(strings.Fields(s), ",") + "]"

	var rawPackets []interface{}
	err := json.Unmarshal([]byte(lines), &rawPackets)
	if err != nil {
		panic(err)
	}

	packets := make([]Packet, len(rawPackets))
	for i, rawPacket := range rawPackets {
		packets[i] = unmarshalPacket(rawPacket)
	}

	return packets
}

var sample = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

package test

import (
	"fmt"
	"testing"
)

type TestCase[I any, A any, A2 any] struct {
	Name string

	Input string
	Ans1  A
	Ans2  A2

	Skip1 bool
	Skip2 bool
}

func (tc *TestCase[I, A, A2]) Run(
	t *testing.T,
	parse func(string) I,
	p1 func(I) A,
	p2 func(I) A2,
	eq func(A, A) bool,
	eq2 func(A2, A2) bool,
	first bool,
) {
	name := tc.Name
	if first {
		name += "Part1"
	} else {
		name += "Part2"
	}

	t.Run(name, func(t *testing.T) {
		skip := (first && tc.Skip1) || (!first && tc.Skip2)
		if skip && tc.Input == "" {
			t.Skip("Input for day doesn't exist")
		}
		if skip {
			t.Skip(fmt.Sprintf("Purposefully skipping %s", name))
		}

		input := parse(tc.Input)
		if first {
			res := p1(input)
			ans := tc.Ans1
			if !eq(ans, res) {
				t.Logf("expected: %v, got: %v", ans, res)
				t.Fail()
			}
		} else {
			res := p2(input)
			ans := tc.Ans2
			if !eq2(ans, res) {
				t.Logf("expected: %v, got: %v", ans, res)
				t.Fail()
			}
		}
	})
}

type TestCases[I any, A any, A2 any] []TestCase[I, A, A2]

func (tcs *TestCases[I, A, A2]) Run(
	t *testing.T,
	parse func(string) I,
	p1 func(I) A,
	p2 func(I) A2,
	eq func(A, A) bool,
	eq2 func(A2, A2) bool,
) {
	for _, tc := range *tcs {
		tc.Run(
			t,
			parse,
			p1,
			p2,
			eq,
			eq2,
			true,
		)
		tc.Run(
			t,
			parse,
			p1,
			p2,
			eq,
			eq2,
			false,
		)
	}
}

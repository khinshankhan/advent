package test

import "testing"

type TestCase[RI any, I any, A any] struct {
	Name string

	Input RI
	Ans1  A
	Ans2  A

	Skip1 bool
	Skip2 bool
}

func (tc *TestCase[RI, I, A]) Run(
	t *testing.T,
	parse func(RI) I,
	p1 func(I) A,
	p2 func(I) A,
	eq func(A, A) bool,
	first bool,
) {
	name := tc.Name
	if first {
		name += "Part1"
	} else {
		name += "Part2"
	}

	t.Run(name, func(t *testing.T) {
		if (first && tc.Skip1) || (!first && tc.Skip2) {
			t.Skip("Input for day doesn't exist")
		}

		var res A
		var ans A
		input := parse(tc.Input)
		if first {
			res = p1(input)
			ans = tc.Ans1
		} else {
			res = p2(input)
			ans = tc.Ans2
		}
		if !eq(ans, res) {
			t.Logf("expected: %v, got: %v", ans, res)
			t.Fail()
		}
	})
}

type TestCases[RI any, I any, A any] []TestCase[RI, I, A]

func (tcs *TestCases[RI, I, A]) Run(
	t *testing.T,
	parse func(RI) I,
	p1 func(I) A,
	p2 func(I) A,
	eq func(A, A) bool,
) {
	for _, tc := range *tcs {
		tc.Run(
			t,
			parse,
			p1,
			p2,
			eq,
			true,
		)
		tc.Run(
			t,
			parse,
			p1,
			p2,
			eq,
			false,
		)
	}
}

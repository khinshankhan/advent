package test

import "testing"

type TestCase[I any, A any] struct {
	Name string

	Input I
	Ans   A

	First bool
	Skip  bool
}

func (tc *TestCase[I, A]) Run(
	t *testing.T,
	p1 func(I) A,
	p2 func(I) A,
	eq func(A, A) bool,
) {
	name := tc.Name
	if tc.First {
		name += "Part1"
	} else {
		name += "Part2"
	}

	t.Run(name, func(t *testing.T) {
		if tc.Skip {
			t.Skip("Input for day doesn't exist")
		}

		var res A
		if tc.First {
			res = p1(tc.Input)
		} else {
			res = p2(tc.Input)
		}
		if !eq(tc.Ans, res) {
			t.Logf("expected: %v, got: %v", tc.Ans, res)
			t.Fail()
		}
	})
}

type TestCases[I any, A any] []TestCase[I, A]

func (tcs *TestCases[I, A]) Run(
	t *testing.T,
	p1 func(I) A,
	p2 func(I) A,
	eq func(A, A) bool,
) {
	for _, tc := range *tcs {
		tc.Run(
			t,
			p1,
			p2,
			eq,
		)
	}
}

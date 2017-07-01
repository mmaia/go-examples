package mathapp

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := [] struct {
		in, want float64
	} {
		{9, 3},
		{4, 2},
	}
	for _, c:= range cases {
		got := Sqrt(c.in)
		if got != c.want {
			t.Errorf("Sqrt(%q) == %q, want %q ", c.in, got, c.want)
		}
	}
}

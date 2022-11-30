package layout

import (
	"fmt"
	"reflect"
	"slava0135/nanoservice/rules"
	"strings"
	"testing"
)

type linkedSquaresTest struct {
	name string
	args Point
	want []Point
}

func TestLinkedCornerSquares(t *testing.T) {
	tests := []linkedSquaresTest{
		{"zero corner", Point{0, 0}, []Point{{1, 1}}},
		{"n corner", Point{rules.N - 1, rules.N - 1}, []Point{{8, 8}}},
		{"side y", Point{1, 0}, []Point{{2, 1}, {0, 1}}},
		{"side x", Point{0, 1}, []Point{{1, 0}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedCornerSquares(tt.args.X, tt.args.Y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedCornerSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedSideSquares(t *testing.T) {
	tests := []linkedSquaresTest{
		{"zero corner", Point{0, 0}, []Point{{1, 0}, {0, 1}}},
		{"n corner", Point{rules.N - 1, rules.N - 1}, []Point{{8, 9}, {9, 8}}},
		{"side y", Point{1, 0}, []Point{{0, 0}, {2, 0}, {1, 1}}},
		{"side x", Point{0, 1}, []Point{{0, 0}, {1, 1}, {0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedSideSquares(tt.args.X, tt.args.Y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedSideSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedSquares(t *testing.T) {
	tests := []linkedSquaresTest{
		{"zero corner", Point{0, 0}, []Point{{1, 1}, {1, 0}, {0, 1}}},
		{"n corner", Point{rules.N - 1, rules.N - 1}, []Point{{8, 8}, {8, 9}, {9, 8}}},
		{"side y", Point{1, 0}, []Point{{2, 1}, {0, 1}, {0, 0}, {2, 0}, {1, 1}}},
		{"side x", Point{0, 1}, []Point{{1, 0}, {1, 2}, {0, 0}, {1, 1}, {0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedSquares(tt.args.X, tt.args.Y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLayout_GridParsing(t *testing.T) {
	l := simpleLayout()
	
	s := l.String()
	pl, _, err := ParseLayout(s[:len(s)-1])
	if err != nil {
		t.Fatalf("parsing layout shouldn't fail: %v", err)
	}
	if reflect.DeepEqual(l, pl) != true {
		t.Fatalf("initial and parsed layouts should be the same:\n%v\n%v", l, pl)
	}
}

func TestLayout_ShipParsing(t *testing.T) {
	l := simpleLayout()
	
	ships := ShipsFromLayout(l)

	var sb strings.Builder
	for _, s := range ships {
		sb.WriteString(fmt.Sprintf("%v\n", s))
	}

	pl, _, err := ParseLayout(sb.String()[:sb.Len()-1])
	if err != nil {
		t.Fatalf("parsing layout shouldn't fail: %v", err)
	}
	if reflect.DeepEqual(l, pl) != true {
		t.Fatalf("initial and parsed layouts should be the same:\n%v\n%v", l, pl)
	}
}

func simpleLayout() Layout {
	l := Layout{}
	
	l[3][3] = true
	l[3][4] = true

	l[8][8] = true
	
	l[1][0] = true
	l[2][0] = true
	l[3][0] = true

	return l
}
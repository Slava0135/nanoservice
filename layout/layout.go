package layout

import (
	"fmt"
	"slava0135/nanoservice/rules"
	"strings"
)

type Layout [rules.N][rules.N]bool

type Point struct {
	X, Y uint
}

type Ship struct {
	P1, P2 Point
}

func (l Layout) String() string {
	var sb strings.Builder
	for _, v := range l {
		for _, v := range v {
			if v {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('o')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func NewShip(x1, y1, x2, y2 uint) Ship {
	return Ship{Point{x1, y1}, Point{x2, y2}}
}

func ParseLayout(s string) (layout Layout, ships []Ship, err error) {
	for i, line := range strings.Split(s, "\n") {
		var x1, y1, x2, y2 uint
		_, err := fmt.Sscanf(line, "{{%d %d} {%d %d}}", &x1, &y1, &x2, &y2)
		if err != nil {
			return layout, ships, fmt.Errorf("failed to parse line %d: %v", i+1, err)
		}
		if x1 != x2 && y1 != y2 {
			return layout, ships, fmt.Errorf("failed to parse line %d: points are not on the same line", i+1)
		}
		if x1 >= rules.N || x2 >= rules.N || y1 >= rules.N || y2 >= rules.N {
			return layout, ships, fmt.Errorf("failed to parse line %d: points are outside of game field", i+1)
		}
		ships = append(ships, NewShip(x1, y1, x2, y2))
		if x1 != x2 {
			var x_low, x_high uint
			if x1 < x2 {
				x_low, x_high = x1, x2
			} else {
				x_low, x_high = x2, x1
			}
			for x := x_low; x <= x_high; x++ {
				layout[x][y1] = true
			}
		} else {
			var y_low, y_high uint
			if y1 < y2 {
				y_low, y_high = y1, y2
			} else {
				y_low, y_high = y2, y1
			}
			for y := y_low; y <= y_high; y++ {
				layout[x1][y] = true
			}
		}
	}
	return layout, ships, nil
}

func LinkedSquares(x, y uint) []Point {
	arr := make([]Point, 0)
	arr = append(arr, LinkedCornerSquares(x, y)...)
	arr = append(arr, LinkedSideSquares(x, y)...)
	return arr
}

func LinkedSideSquares(x, y uint) []Point {
	arr := make([]Point, 0)
	if x > 0 {
		arr = append(arr, Point{x - 1, y})
	}
	if y > 0 {
		arr = append(arr, Point{x, y - 1})
	}
	if x < rules.N-1 {
		arr = append(arr, Point{x + 1, y})
	}
	if y < rules.N-1 {
		arr = append(arr, Point{x, y + 1})
	}
	return arr
}

func LinkedCornerSquares(x, y uint) []Point {
	arr := make([]Point, 0)
	if x > 0 && y > 0 {
		arr = append(arr, Point{x - 1, y - 1})
	}
	if x < rules.N-1 && y > 0 {
		arr = append(arr, Point{x + 1, y - 1})
	}
	if x < rules.N-1 && y < rules.N-1 {
		arr = append(arr, Point{x + 1, y + 1})
	}
	if x > 0 && y < rules.N-1 {
		arr = append(arr, Point{x - 1, y + 1})
	}
	return arr
}

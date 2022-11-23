package layout

import (
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

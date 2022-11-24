package shots

import (
	"fmt"
	"slava0135/nanoservice/layout"
	"slava0135/nanoservice/rules"
	"strings"
)

type Results struct {
	Untouched, Damaged, Destroyed []layout.Ship 
}

func ParseShots(s string) ([]layout.Point, error) {
	var shots []layout.Point
	for i, line := range strings.Split(s, "\n") {
		var x, y uint
		_, err := fmt.Sscanf(line, "{%d %d}", &x, &y)
		if err != nil {
			return shots, fmt.Errorf("failed to parse line %d: %v", i+1, err)
		}
		if x >= rules.N || y >= rules.N {
			return shots, fmt.Errorf("failed to parse line %d: shots are outside of game field", i+1)
		}
		shots = append(shots, layout.Point{X: x, Y: y})
	}
	return shots, nil
}

func ReplayGame(ships []layout.Ship, shots []layout.Point) Results {
	var res Results
	for _, ship := range ships {
		var segments []layout.Point
		if ship.P1.X != ship.P2.X {
			var x_low, x_high uint
			if ship.P1.X < ship.P2.X {
				x_low, x_high = ship.P1.X, ship.P2.X
			} else {
				x_low, x_high = ship.P2.X, ship.P1.X
			}
			for x := x_low; x <= x_high; x++ {
				segments = append(segments, layout.Point{X: x, Y: ship.P1.Y})
			}
		} else {
			var y_low, y_high uint
			if ship.P1.Y < ship.P2.Y {
				y_low, y_high = ship.P1.Y, ship.P2.Y
			} else {
				y_low, y_high = ship.P1.Y, ship.P2.Y
			}
			for y := y_low; y <= y_high; y++ {
				segments = append(segments, layout.Point{X: ship.P1.X, Y: y})
			}
		}
		length := len(segments)
		for _, shot := range shots {
			for i := 0; i < len(segments); i++ {
				if shot == segments[i] {
					segments = append(segments[:i], segments[i+1:]...)
					break
				}
			}
		}
		if length == len(segments) {
			res.Untouched = append(res.Untouched, ship)
		} else if len(segments) == 0 {
			res.Destroyed = append(res.Destroyed, ship)
		} else {
			res.Damaged = append(res.Damaged, ship)
		}
	}
	return res
} 
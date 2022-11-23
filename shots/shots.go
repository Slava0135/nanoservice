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
	return res
} 
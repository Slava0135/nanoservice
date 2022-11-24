package validate

import (
	"slava0135/nanoservice/layout"
	"slava0135/nanoservice/rules"
)

func Validate(gameLayout layout.Layout) bool {
	for x := uint(0); x < rules.N; x++ {
		for y := uint(0); y < rules.N; y++ {
			if gameLayout[x][y] {
				for _, p := range layout.LinkedCornerSquares(x, y) {
					if gameLayout[p.X][p.Y] {
						return false
					}
				}
			}
		}
	}
	var ships [rules.N]uint
	for x := 0; x < rules.N; x++ {
		length := 0
		for y := 0; y < rules.N; y++ {
			if gameLayout[x][y] {
				length++
			} else {
				ships[length]++
				length = 0
			}
		}
		ships[length]++
	}
	for y := 0; y < rules.N; y++ {
		length := 0
		for x := 0; x < rules.N; x++ {
			if gameLayout[x][y] {
				length++
			} else {
				ships[length]++
				length = 0
			}
		}
		ships[length]++
	}
	// large ships (length > 1) are counted towards 1-sized ships
	var largeShipParts uint
	for i := 2; i < rules.N; i++ {
		largeShipParts += uint(i) * ships[i]
	}
	ships[1] -= largeShipParts
	// 1-sized ships counted twice
	ships[1] /= 2
	for i := 1; i < len(rules.Ships); i++ {
		if ships[i] != rules.Ships[i] {
			return false
		}
	}
	return true
}

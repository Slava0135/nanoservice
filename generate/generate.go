package generate

import (
	"math/rand"
	"slava0135/nanoservice/layout"
	"slava0135/nanoservice/rules"
)

func NewGameLayout() layout.Layout {
	var gameLayout layout.Layout
	for shipLen, shipCount := range rules.Ships {
		for s := 0; s < int(shipCount); s++ {
		retry:
			x := uint(rand.Uint32() % rules.N)
			y := uint(rand.Uint32() % rules.N)
			if gameLayout[x][y] {
				goto retry
			}
			for _, p := range layout.LinkedSquares(x, y) {
				if gameLayout[p.X][p.Y] {
					goto retry
				}
			}
			nextLayout := gameLayout
			nextLayout[x][y] = true
			if rand.Int()%2 == 0 {
				if x+uint(shipLen) >= rules.N {
					goto retry
				}
				for i := uint(1); i < uint(shipLen); i++ {
					for _, p := range layout.LinkedSquares(x+i, y) {
						if gameLayout[p.X][p.Y] {
							goto retry
						}
					}
					nextLayout[x+i][y] = true
				}
			} else {
				if y+uint(shipLen) >= rules.N {
					goto retry
				}
				for i := uint(1); i < uint(shipLen); i++ {
					for _, p := range layout.LinkedSquares(x, y+i) {
						if gameLayout[p.X][p.Y] {
							goto retry
						}
					}
					nextLayout[x][y+i] = true
				}
			}
			gameLayout = nextLayout
		}
	}
	return gameLayout
}

package generate

import (
	"math/rand"
	"slava0135/nanoservice/rules"
)

func NewLayout() [rules.N][rules.N]bool {
	var layout [rules.N][rules.N]bool
	for shipLen, shipCount := range rules.Ships {
		for s := 0; s < int(shipCount); s++ {
		try_place: 
			x, y := rand.Uint32()%rules.N, rand.Uint32()%rules.N
			if layout[x][y] {
				goto try_place
			}
			nextLayout := layout 
			if rand.Int() > 0 {
				if x + uint32(shipLen) >= rules.N {
					goto try_place
				}
				for i := uint32(0); i < uint32(shipLen); i++ {
					nextLayout[x+i][y] = true
				}
			} else {
				if y + uint32(shipLen) >= rules.N {
					goto try_place
				}
				for i := uint32(0); i < uint32(shipLen); i++ {
					nextLayout[x][y+i] = true
				}
			}
			if !rules.Valid(nextLayout) {
				goto try_place
			}
			layout = nextLayout
		}
	}
	return layout
}
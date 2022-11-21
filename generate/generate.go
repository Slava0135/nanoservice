package generate

import (
	"math/rand"
	"slava0135/nanoservice/rules"
)

func NewLayout() [rules.N][rules.N]bool {
	var layout [rules.N][rules.N]bool
	for ship_len, ship_count := range rules.Ships {
		for s := 0; s < int(ship_count); s++ {
		try_place:
			x, y := rand.Uint32()%rules.N, rand.Uint32()%rules.N
			if layout[x][y] {
				goto try_place
			}
			next_layout := layout 
			if rand.Int() > 0 {
				if x + uint32(ship_len) >= rules.N {
					goto try_place
				}
				for i := uint32(0); i < uint32(ship_len); i++ {
					next_layout[x+i][y] = true
				}
			} else {
				if y + uint32(ship_len) >= rules.N {
					goto try_place
				}
				for i := uint32(0); i < uint32(ship_len); i++ {
					next_layout[x][y+i] = true
				}
			}
			if !rules.Valid(next_layout) {
				goto try_place
			}
			layout = next_layout
		}
	}
	return layout
}
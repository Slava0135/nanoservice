package generate

import (
	"slava0135/nanoservice/validate"
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 0; i < 20; i++ {
		l, _ := NewGameLayout()
		if validate.Validate(l) == false {
			t.Fatalf("generated layout is not valid:\n%v", l)
		}
	}
}
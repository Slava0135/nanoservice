package validate

import (
	"slava0135/nanoservice/generate"
	"slava0135/nanoservice/layout"
	"testing"
)

func TestValidate_Empty(t *testing.T) {
	l := layout.Layout{}
	if Validate(l) != false {
		t.Errorf("empty layout should be invalid")
	}
}

func TestValidate_Corners(t *testing.T) {
	l := layout.Layout{}
	l[3][3] = true
	l[4][4] = true
	if Validate(l) != false {
		t.Errorf("meeting corners should be invalid")
	}
}

func TestValidate_Generate(t *testing.T) {
	l, _ := generate.NewGameLayout()
	if Validate(l) != true {
		t.Errorf("generated layout should be invalid")
	}
}
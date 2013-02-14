package hashimg

import (
	"image/color"
	"testing"
)

var colorMaps = map[string]color.Palette{
	"CGA":           CgaPalette,
	"High Contrast": HiContrastPalette,
}

func TestUniqueness(t *testing.T) {
	for pname, palette := range colorMaps {
		m := make(map[color.Color]int)
		for i, c := range palette {
			if oi, contains := m[c]; contains {
				t.Errorf("Palette %v: %d, %d are duplicates (%v)", pname, oi, i, c)
			}
			m[c] = i
		}
	}
}

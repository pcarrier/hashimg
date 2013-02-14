package hashimg

import (
	"testing"
)

var rectSizeTests = []struct {
	surface int
	w       int
	l       int
}{
	{0, 0, 0},
	{3, 3, 1},
	{6, 3, 2},
	{9, 3, 3},
	{12, 4, 3},
	{17, 17, 1},
	{255, 17, 15},
	{257, 257, 1},
}

func TestRectSize(t *testing.T) {
	for _, c := range rectSizeTests {
		w, l := rectSize(c.surface)
		if w*l != c.surface {
			t.Errorf("rectSize incorrect: %d×%d≠%d", w, l, c.surface)
		}
		if w != c.w || l != c.l {
			t.Errorf("rectSize unexpected: (%d,%d)≠(%d,%d)", w, l, c.w, c.l)
		}
	}
}

func BenchmarkRectSize9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectSize(9)
	}
}

func BenchmarkRectSizeGroling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rectSize(i)
	}
}

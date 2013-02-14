package hashimg

import (
	"math"
)

func rectSize(surface int) (int, int) {
	if surface <= 0 {
		return 0, 0
	}
	i := int(math.Sqrt(float64(surface)))
	for surface%i != 0 {
		i--
	}
	return surface / i, i
}

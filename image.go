package hashimg

import (
	"image"
	"image/color"
)

func MakeImage(src []byte, p color.Palette) image.Image {
	pixels := 2 * len(src)
	w, l := rectSize(pixels)
	img := image.NewPaletted(image.Rect(0, 0, w, l), p)

	i := 0
	for x := 0; x < w; x++ {
		for y := 0; y < l; y++ {
			b := src[i/2]
			var q byte
			if i%2 == 0 {
				q = (b & 0xF0) >> 4
			} else {
				q = b & 0xF
			}
			img.SetColorIndex(x, y, q)
			i++
		}
	}
	return img
}

package hashimg

import (
	"image/color"
)

var CgaPalette = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, // black
	color.RGBA{0x00, 0x00, 0xAA, 0xFF}, // blue
	color.RGBA{0x00, 0xAA, 0x00, 0xFF}, // green
	color.RGBA{0x00, 0xAA, 0xAA, 0xFF}, // cyan
	color.RGBA{0xAA, 0x00, 0x00, 0xFF}, // red
	color.RGBA{0xAA, 0x00, 0xAA, 0xFF}, // magenta
	color.RGBA{0xAA, 0x55, 0x00, 0xFF}, // brown
	color.RGBA{0xAA, 0xAA, 0xAA, 0xFF}, // light gray
	color.RGBA{0x55, 0x55, 0x55, 0xFF}, // gray
	color.RGBA{0x55, 0x55, 0xFF, 0xFF}, // light blue
	color.RGBA{0x55, 0xFF, 0x55, 0xFF}, // light green
	color.RGBA{0x55, 0xFF, 0xFF, 0xFF}, // light cyan
	color.RGBA{0xFF, 0x55, 0x55, 0xFF}, // light red
	color.RGBA{0xFF, 0x55, 0xFF, 0xFF}, // light magenta
	color.RGBA{0xFF, 0xFF, 0x55, 0xFF}, // yellow
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, // white (high density)
}

var HiContrastPalette = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	// duplicate (to test tests)
	//color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	color.RGBA{0xAA, 0xAA, 0xAA, 0xFF},
	color.RGBA{0x00, 0x00, 0x88, 0xFF},
	color.RGBA{0x00, 0x88, 0x00, 0xFF},
	color.RGBA{0x00, 0x88, 0x88, 0xFF},
	color.RGBA{0x88, 0x00, 0x00, 0xFF},
	color.RGBA{0x88, 0x00, 0x88, 0xFF},
	color.RGBA{0x88, 0x88, 0x00, 0xFF},
	color.RGBA{0x88, 0x88, 0x88, 0xFF},
}

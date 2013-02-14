package main

import (
	"encoding/hex"
	"flag"
	"github.com/pcarrier/hashimg"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	var src = flag.String("src", "00", "hexadecimal source")
	var dst = flag.String("dst", "out.png", "destination filename")
	var hicontrast = flag.Bool("hi", false, "use high contrast palette")
	flag.Parse()

	s, err := hex.DecodeString(*src)
	if err != nil {
		log.Fatalf("Could not decode hexadecimal (%v)", err)
	}

	var p color.Palette
	if *hicontrast {
		p = hashimg.HiContrastPalette
	} else {
		p = hashimg.CgaPalette
	}

	img := hashimg.MakeImage(s, p)

	file, err := os.Create(*dst)
	if err != nil {
		log.Fatalf("Could not create %v (%v)", file, err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("Could not encode PNG (%v)", err)
	}
}

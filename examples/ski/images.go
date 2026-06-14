package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//
// --- Font ---
//

//go:embed assets/Iceberg-Regular.ttf
var fontBytes []byte

//
// --- Skier ---
//

//go:embed assets/tile_0082.png
var skierFrame1ImageBytes []byte

//go:embed assets/tile_0083.png
var skierFrame2ImageBytes []byte

//go:embed assets/tile_0058.png
var skierSnowTrailFrame1ImageBytes []byte

//go:embed assets/tile_0059.png
var skierSnowTrailFrame2ImageBytes []byte

//
// --- Obstacles ---
//

//go:embed assets/tile_0069.png
var obstacles1ImageBytes []byte

//
// --- World ---
//

//go:embed assets/snowMap2.png
var snowMap1ImageBytes []byte

//go:embed assets/snowMap2.png
var snowMap2ImageBytes []byte

//go:embed assets/rides1.png
var ridesMap1ImageBytes []byte

//go:embed assets/rides2.png
var ridesMap2ImageBytes []byte

//
// --- Utility functions ---
//

// getImageFromFile loads the image from the file and returns an *ebiten.Image
func getImageFromFile(imageBytes []byte) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imageBytes))
	if err != nil {
		log.Fatal("Couldn't find image", err)
		return nil
	}
	return image
}

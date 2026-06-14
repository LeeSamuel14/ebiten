package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Obstacle struct {
	image *ebiten.Image
	x     int
	y     int
}

func getNewObstacle() Obstacle {
	//todo ensure random images of obstacles - have slices and pick random number for obstacle.
	image := getImageFromFile(obstacles1ImageBytes)
	randomX := rng.Intn(ScreenWidth) + 20
	randomY := ScreenHeight + ImageSize
	return Obstacle{
		x:     randomX,
		y:     randomY,
		image: image,
	}
}

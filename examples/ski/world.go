package main

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	snow1Map  *ebiten.Image
	snow1MapX float64
	snow1MapY float64

	snow2Map  *ebiten.Image
	snow2MapX float64
	snow2MapY float64

	rides1Map  *ebiten.Image
	rides1MapX float64
	rides1MapY float64

	rides2Map  *ebiten.Image
	rides2MapX float64
	rides2MapY float64

	obstacles  []Obstacle
	spawnTimer int
	scoreTimer int
}

func (world *World) initWorld() {
	world.snow1Map = getImageFromFile(snowMap1ImageBytes)
	world.snow2Map = getImageFromFile(snowMap2ImageBytes)
	world.snow2MapY = ScreenHeight

	world.rides1Map = getImageFromFile(ridesMap1ImageBytes)
	world.rides2Map = getImageFromFile(ridesMap2ImageBytes)
	world.rides2MapY = ScreenHeight
}

func (world *World) drawSnow(screen *ebiten.Image) {
	drawOptionsSnow1Map := ebiten.DrawImageOptions{}
	drawOptionsSnow1Map.GeoM.Translate(world.snow1MapX, world.snow1MapY)
	screen.DrawImage(world.snow1Map, &drawOptionsSnow1Map)

	drawOptionsSnow2Map := ebiten.DrawImageOptions{}
	drawOptionsSnow2Map.GeoM.Translate(world.snow2MapX, world.snow2MapY)
	screen.DrawImage(world.snow2Map, &drawOptionsSnow2Map)
}

func (world *World) drawRides(screen *ebiten.Image) {
	drawOptionsRides1Map := ebiten.DrawImageOptions{}
	drawOptionsRides1Map.GeoM.Translate(world.rides1MapX, world.rides1MapY)
	screen.DrawImage(world.rides1Map, &drawOptionsRides1Map)

	drawOptionsRides2Map := ebiten.DrawImageOptions{}
	drawOptionsRides2Map.GeoM.Translate(world.rides2MapX, world.rides2MapY)
	screen.DrawImage(world.rides2Map, &drawOptionsRides2Map)
}

func (world *World) drawObstacles(screen *ebiten.Image) {
	for i, _ := range world.obstacles {
		drawOptions := ebiten.DrawImageOptions{}
		drawOptions.GeoM.Translate(float64(world.obstacles[i].x), float64(world.obstacles[i].y))
		screen.DrawImage(world.obstacles[i].image, &drawOptions)
	}
}

func (world *World) moveWorldMap() {
	world.snow1MapY -= WorldMapSpeed
	world.snow2MapY -= WorldMapSpeed
	world.rides1MapY -= WorldMapSpeed
	world.rides2MapY -= WorldMapSpeed

	if world.snow1MapY < -(ScreenHeight) {
		world.snow1MapY = ScreenHeight
	}
	if world.snow2MapY < -(ScreenHeight) {
		world.snow2MapY = ScreenHeight
	}

	if world.rides1MapY < -ScreenHeight {
		world.rides1MapY = ScreenHeight
	}
	if world.rides2MapY < -ScreenHeight {
		world.rides2MapY = ScreenHeight
	}
}

func (world *World) moveObstacles() {
	world.spawnTimer++
	if world.spawnTimer > 20 {
		obstacle := getNewObstacle()
		world.obstacles = append(world.obstacles, obstacle)
		world.spawnTimer = 0
	}
	for i := len(world.obstacles) - 1; i >= 0; i-- {
		world.obstacles[i].y -= WorldMapSpeed
		if world.obstacles[i].y < -ImageSize {
			world.obstacles = slices.Delete(world.obstacles, i, i+1)
		}
	}
}

package main

import "github.com/hajimehoshi/ebiten/v2"

type Animation struct {
	frames       []*ebiten.Image
	currentFrame int
	timer        int
	speed        int
}

func (a *Animation) Update() {
	a.timer++
	if a.timer > a.speed {
		a.timer = 0
		a.currentFrame++

		if a.currentFrame > len(a.frames)-1 {
			a.currentFrame = 0
		}
	}
}

func (a *Animation) CurrentFrame() *ebiten.Image {
	return a.frames[a.currentFrame]
}

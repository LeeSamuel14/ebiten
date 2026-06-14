package main

import "github.com/hajimehoshi/ebiten/v2"

type Skier struct {
	x                  int
	y                  int
	animation          Animation
	snowTrailAnimation Animation
	speed              int
}

func (skier *Skier) initSkierAnimation() {
	skierFrame1 := getImageFromFile(skierFrame1ImageBytes)
	skierFrame2 := getImageFromFile(skierFrame2ImageBytes)
	skier.animation = Animation{
		frames: []*ebiten.Image{
			skierFrame1,
			skierFrame2,
		},
		speed: 9,
	}
}

func (skier *Skier) initSkierSnowTrailAnimation() {
	skierSnowTrailFrame1 := getImageFromFile(skierSnowTrailFrame1ImageBytes)
	skierSnowTrailFrame2 := getImageFromFile(skierSnowTrailFrame2ImageBytes)
	skier.snowTrailAnimation = Animation{
		frames: []*ebiten.Image{
			skierSnowTrailFrame1,
			skierSnowTrailFrame2,
		},
		speed: 5,
	}
}

func (skier *Skier) initSkier() {
	skier.initSkierAnimation()
	skier.initSkierSnowTrailAnimation()
	skier.speed = SkierSpeed
	skier.x = ScreenWidth / 2  //position in middle on screen
	skier.y = ScreenHeight / 4 //position lower on screen
}

func (skier *Skier) moveSkier() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && skier.x > 0 {
		skier.x -= skier.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && skier.x < ScreenWidth-(ImageSize*1.5) {
		skier.x += skier.speed
	}
}

func (skier *Skier) drawSkier(screen *ebiten.Image) {
	// Draw skier animation
	drawOptionsSkier := ebiten.DrawImageOptions{}
	drawOptionsSkier.GeoM.Translate(float64(skier.x), float64(skier.y))
	//drawOptionsSkier.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(skier.animation.CurrentFrame(), &drawOptionsSkier)

	// Draw skier snow trail animation
	drawOptionsSnowTrail := ebiten.DrawImageOptions{}
	drawOptionsSnowTrail.GeoM.Translate(float64(skier.x), float64(skier.y-ImageSize))
	//drawOptionsSnowTrail.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(skier.snowTrailAnimation.CurrentFrame(), &drawOptionsSnowTrail)
}

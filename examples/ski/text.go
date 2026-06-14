package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	_ "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	_ "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	_ "golang.org/x/image/font/opentype"
)

var gameFont font.Face

func initGameFont() {
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	gameFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (game *Game) drawGameTextCentre(screen *ebiten.Image, msg string, color color.Color) {
	bounds := text.BoundString(gameFont, msg)

	x := (ScreenWidth - bounds.Dx()) / 2
	y := ScreenHeight / 2

	text.Draw(screen, msg, gameFont, x, y, color)
}

func (game *Game) drawGameTextScore(screen *ebiten.Image, msg string, color color.Color) {
	text.Draw(screen, msg, gameFont, ScreenWidth/2-20, 20, color)
}

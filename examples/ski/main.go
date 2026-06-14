package main

import (
	_ "embed"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth    = 480
	ScreenHeight   = 480
	SkierSpeed     = 5
	ImageSize      = 16
	ScoreIncreaser = 10
	WorldMapSpeed  = 5
)

type GameState int

const (
	Playing GameState = iota
	GameOver
	Start
)

type Game struct {
	skier Skier
	world World
	state GameState
	score int
}

// rng used to generate random numbers in the game
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func (game *Game) Update() error {
	switch game.state {
	case Start:
		game.startGame()

	case Playing:
		game.skier.moveSkier()
		game.skier.animation.Update()
		game.skier.snowTrailAnimation.Update()
		game.world.moveWorldMap()
		game.world.moveObstacles()
		game.checkCollisions()
		game.updateScore()

	case GameOver:
		game.gameOver()
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	switch game.state {
	case Start:
		screen.Fill(color.RGBA{R: 200, G: 225, B: 255, A: 255})
		game.drawGameTextCentre(screen, "Press space bar to play", color.Black)

	case Playing:
		screen.Fill(color.RGBA{R: 200, G: 225, B: 255, A: 255})

		//order matters to show player underneath and above the world items
		game.world.drawSnow(screen)
		game.world.drawObstacles(screen)
		game.skier.drawSkier(screen)
		game.world.drawRides(screen)
		game.drawGameTextScore(screen, strconv.Itoa(game.score), color.Black)

	case GameOver:
		screen.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})
		game.drawGameTextCentre(screen, "Game Over.\nScore: "+strconv.Itoa(game.score)+"\nPress space bar to restart.", color.RGBA{R: 220, G: 40, B: 40, A: 255})
	}

}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (game *Game) startGame() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.state = Playing
	}
}

func (game *Game) gameOver() {
	game.world.obstacles = game.world.obstacles[:0]
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		game.score = 0
		game.state = Playing
	}
}

func (game *Game) checkCollisions() {
	obstacles := game.world.obstacles
	skier := game.skier
	for i, _ := range obstacles {
		if collides(skier, obstacles[i]) {
			game.state = GameOver
		}
	}
}

func collides(skier Skier, obstacle Obstacle) bool {
	return skier.x < obstacle.x+(ImageSize/2) &&
		skier.x+(ImageSize/2) > obstacle.x &&
		skier.y < obstacle.y+(ImageSize/2) &&
		skier.y+(ImageSize/2) > obstacle.y
}

func (game *Game) updateScore() {
	game.world.scoreTimer++
	if game.world.scoreTimer > 30 {
		game.world.scoreTimer = 0
		game.score += ScoreIncreaser
	}
}

func initGameWindow() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Ski")
}

func main() {
	initGameFont()
	initGameWindow()

	world := World{}
	skier := Skier{}

	world.initWorld()
	skier.initSkier()

	game := &Game{
		skier: skier,
		world: world,
		state: Start,
	}

	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}

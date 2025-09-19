package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"snake-game/internal"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")
	ebiten.SetWindowResizable(false)

	game := internal.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		// Handle termination gracefully
		if err == ebiten.Termination {
			// Exit normally
			return
		}
		log.Fatal(err)
	}
}
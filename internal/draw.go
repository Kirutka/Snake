package internal

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw gradient background
	g.drawGradientBackground(screen)

	// Draw background particles
	for _, particle := range g.particles {
		vector.DrawFilledCircle(screen, float32(particle.x), float32(particle.y), 2, particle.color, false)
	}

	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying:
		g.drawGame(screen)
	case StatePaused:
		g.drawGame(screen)
		g.drawPause(screen)
	case StateGameOver:
		g.drawGame(screen)
		g.drawGameOver(screen)
	}
}

func (g *Game) drawGradientBackground(screen *ebiten.Image) {
	// Draw dark gradient background
	for y := 0; y < ScreenHeight; y++ {
		alpha := uint8(20 + (y * 20 / ScreenHeight))
		c := color.RGBA{10, 10, 30, alpha}
		vector.DrawFilledRect(screen, 0, float32(y), float32(ScreenWidth), 1, c, false)
	}
	
	// Draw grid pattern
	for x := 0; x < ScreenWidth; x += 40 {
		for y := 0; y < ScreenHeight; y += 40 {
			if (x/40+y/40)%2 == 0 {
				vector.DrawFilledRect(screen, float32(x), float32(y), 20, 20, color.RGBA{20, 20, 40, 50}, false)
			}
		}
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	// Draw title with glow effect
	title := "SNAKE GAME"
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			if i != 0 || j != 0 {
				text.Draw(screen, title, basicfont.Face7x13, ScreenWidth/2-len(title)*7/2+i, 100+j, color.RGBA{0, 100, 200, 100})
			}
		}
	}
	text.Draw(screen, title, basicfont.Face7x13, ScreenWidth/2-len(title)*7/2, 100, color.RGBA{100, 200, 255, 255})
	
	// Draw menu items
	instructions := []string{
		"CONTROLS:",
		"Arrow Keys - Move snake",
		"Space - Pause game",
		"Enter - Start game",
		"Escape - Exit game",
		"",
		fmt.Sprintf("Speed: %d (Arrow Keys to change)", g.speedLevel),
	}
	
	for i, instruction := range instructions {
		text.Draw(screen, instruction, basicfont.Face7x13, 50, 150+i*20, color.White)
	}
	
	// Draw animated border
	t := float64(ebiten.GamepadAxisValue(0, 0)) * 0.01
	borderColor := color.RGBA{
		uint8(100 + 100*math.Sin(t)),
		uint8(100 + 100*math.Cos(t)),
		uint8(200 + 55*math.Sin(t)),
		255,
	}
	
	vector.StrokeRect(screen, 10, 10, ScreenWidth-20, ScreenHeight-20, 2, borderColor, false)
	
	// Draw start prompt
	startText := "Press ENTER to start game"
	text.Draw(screen, startText, basicfont.Face7x13, ScreenWidth/2-len(startText)*7/2, ScreenHeight-100, color.RGBA{255, 255, 100, 255})
}

func (g *Game) drawGame(screen *ebiten.Image) {
	// Draw game border
	vector.StrokeRect(screen, 0, 0, ScreenWidth, ScreenHeight, 2, color.RGBA{50, 150, 200, 100}, false)
	
	// Draw snake
	for i, segment := range g.snake {
		c := color.RGBA{50, 255, 50, 255} // Green for body
		if i == 0 {
			c = color.RGBA{100, 255, 100, 255} // Lighter green for head
		}
		vector.DrawFilledRect(screen, 
			float32(segment.X*GridSize), 
			float32(segment.Y*GridSize), 
			GridSize, GridSize, c, false)
		
		// Add border to snake segments
		vector.StrokeRect(screen,
			float32(segment.X*GridSize),
			float32(segment.Y*GridSize),
			GridSize, GridSize, 1, color.RGBA{0, 150, 0, 255}, false)
	}

	// Draw food with pulsing effect
	t := float64(g.frameCount) * 0.2
	pulse := 1 + 0.2*math.Sin(t)
	size := float32(GridSize * pulse)

	
	vector.DrawFilledCircle(screen, 
		float32(g.food.X*GridSize+GridSize/2), 
		float32(g.food.Y*GridSize+GridSize/2), 
		size/2, color.RGBA{255, 50, 50, 255}, false)
	
	// Draw food core
	vector.DrawFilledCircle(screen, 
		float32(g.food.X*GridSize+GridSize/2), 
		float32(g.food.Y*GridSize+GridSize/2), 
		size/4, color.RGBA{255, 200, 200, 255}, false)

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, basicfont.Face7x13, 10, 20, color.White)
	
	// Draw speed
	speedText := fmt.Sprintf("Speed: %d", g.speedLevel)
	text.Draw(screen, speedText, basicfont.Face7x13, ScreenWidth-100, 20, color.White)
}

func (g *Game) drawPause(screen *ebiten.Image) {
	// Semi-transparent overlay
	vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{0, 0, 30, 180}, false)
	
	// Draw pause menu
	pauseText := "PAUSED"
	text.Draw(screen, pauseText, basicfont.Face7x13, ScreenWidth/2-len(pauseText)*7/2, ScreenHeight/2-40, color.RGBA{255, 255, 100, 255})
	
	options := []string{
		"Enter/Space - Resume",
		"Escape - Main Menu",
	}
	
	for i, option := range options {
		text.Draw(screen, option, basicfont.Face7x13, ScreenWidth/2-len(option)*7/2, ScreenHeight/2+i*20, color.White)
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	// Semi-transparent overlay
	vector.DrawFilledRect(screen, 0, 0, ScreenWidth, ScreenHeight, color.RGBA{30, 0, 0, 180}, false)
	
	// Draw game over text
	gameOverText := "GAME OVER"
	text.Draw(screen, gameOverText, basicfont.Face7x13, ScreenWidth/2-len(gameOverText)*7/2, ScreenHeight/2-60, color.RGBA{255, 100, 100, 255})
	
	scoreText := fmt.Sprintf("Final Score: %d", g.score)
	text.Draw(screen, scoreText, basicfont.Face7x13, ScreenWidth/2-len(scoreText)*7/2, ScreenHeight/2-30, color.White)
	
	options := []string{
		"Enter - New Game",
		"Escape - Main Menu",
	}
	
	for i, option := range options {
		text.Draw(screen, option, basicfont.Face7x13, ScreenWidth/2-len(option)*7/2, ScreenHeight/2+i*20, color.White)
	}
}
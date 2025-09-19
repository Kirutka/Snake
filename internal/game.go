package internal

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	GridSize     = 20
)

type Position struct {
	X, Y int
}

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
)

type Particle struct {
	x, y     float64
	vx, vy   float64
	lifetime int
	color    color.RGBA
}

type Game struct {
	state      GameState
	snake      []Position
	direction  Direction
	nextDir    Direction
	food       Position
	score      int
	speed      int
	speedLevel int
	frameCount int
	particles  []Particle
	shouldExit bool
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		state:      StateMenu,
		snake:      []Position{{X: 10, Y: 10}, {X: 9, Y: 10}, {X: 8, Y: 10}},
		direction:  DirRight,
		nextDir:    DirRight,
		food:       Position{X: 15, Y: 15},
		score:      0,
		speed:      10,
		speedLevel: 2,
		frameCount: 0,
		particles:  make([]Particle, 0),
		shouldExit: false,
	}
	game.generateFood()
	game.generateBackgroundParticles()
	return game
}

func (g *Game) generateBackgroundParticles() {
	g.particles = make([]Particle, 0)
	for i := 0; i < 50; i++ {
		particle := Particle{
			x:        rand.Float64() * ScreenWidth,
			y:        rand.Float64() * ScreenHeight,
			vx:       (rand.Float64() - 0.5) * 0.5,
			vy:       (rand.Float64() - 0.5) * 0.5,
			lifetime: rand.Intn(300) + 100,
			color:    color.RGBA{50, 50, 100, 100},
		}
		g.particles = append(g.particles, particle)
	}
}

func (g *Game) Update() error {
	// Handle exit request
	if g.shouldExit {
		return ebiten.Termination
	}

	// Update background particles
	for i := range g.particles {
		g.particles[i].x += g.particles[i].vx
		g.particles[i].y += g.particles[i].vy
		g.particles[i].lifetime--
		
		if g.particles[i].x < 0 || g.particles[i].x > ScreenWidth {
			g.particles[i].vx = -g.particles[i].vx
		}
		if g.particles[i].y < 0 || g.particles[i].y > ScreenHeight {
			g.particles[i].vy = -g.particles[i].vy
		}
		
		if g.particles[i].lifetime <= 0 {
			g.particles[i] = Particle{
				x:        rand.Float64() * ScreenWidth,
				y:        rand.Float64() * ScreenHeight,
				vx:       (rand.Float64() - 0.5) * 0.5,
				vy:       (rand.Float64() - 0.5) * 0.5,
				lifetime: rand.Intn(300) + 100,
				color:    color.RGBA{50, 50, 100, 100},
			}
		}
	}

	switch g.state {
	case StateMenu:
		g.updateMenu()
	case StatePlaying:
		g.updateGame()
	case StatePaused:
		g.updatePause()
	case StateGameOver:
		g.updateGameOver()
	}
	return nil
}

func (g *Game) updateMenu() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.state = StatePlaying
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.speedLevel > 1 {
		g.speedLevel--
		g.speed = g.speedLevel * 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.speedLevel < 5 {
		g.speedLevel++
		g.speed = g.speedLevel * 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// Exit game from menu
		g.shouldExit = true
	}
}

func (g *Game) updatePause() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.state = StatePlaying
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.state = StateMenu
		g.resetGame()
	}
}

func (g *Game) updateGameOver() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.resetGame()
		g.state = StatePlaying
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.state = StateMenu
		g.resetGame()
	}
}

func (g *Game) resetGame() {
	g.snake = []Position{{X: 10, Y: 10}, {X: 9, Y: 10}, {X: 8, Y: 10}}
	g.direction = DirRight
	g.nextDir = DirRight
	g.food = Position{X: 15, Y: 15}
	g.score = 0
	g.frameCount = 0
	g.generateFood()
}

func (g *Game) updateGame() {
	// Handle pause
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.state = StatePaused
		return
	}

	// Handle direction
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.direction != DirDown {
		g.nextDir = DirUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.direction != DirUp {
		g.nextDir = DirDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && g.direction != DirRight {
		g.nextDir = DirLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && g.direction != DirLeft {
		g.nextDir = DirRight
	}

	g.frameCount++
	if g.frameCount < 60/g.speed {
		return
	}
	g.frameCount = 0

	g.direction = g.nextDir

	// Move snake
	head := g.snake[0]
	newHead := Position{X: head.X, Y: head.Y}

	switch g.direction {
	case DirUp:
		newHead.Y--
	case DirDown:
		newHead.Y++
	case DirLeft:
		newHead.X--
	case DirRight:
		newHead.X++
	}

	// Check wall collision
	if newHead.X < 0 || newHead.X >= ScreenWidth/GridSize || newHead.Y < 0 || newHead.Y >= ScreenHeight/GridSize {
		g.state = StateGameOver
		return
	}

	// Check self collision
	for _, segment := range g.snake {
		if newHead.X == segment.X && newHead.Y == segment.Y {
			g.state = StateGameOver
			return
		}
	}

	// Add new head
	g.snake = append([]Position{newHead}, g.snake...)

	// Check food collision
	if newHead.X == g.food.X && newHead.Y == g.food.Y {
		g.score += 10
		g.generateFood()
		// Add celebration particles
		g.addFoodParticles(float64(g.food.X*GridSize+GridSize/2), float64(g.food.Y*GridSize+GridSize/2))
	} else {
		// Remove tail if no food eaten
		g.snake = g.snake[:len(g.snake)-1]
	}
}

func (g *Game) addFoodParticles(x, y float64) {
	for i := 0; i < 20; i++ {
		particle := Particle{
			x:        x,
			y:        y,
			vx:       (rand.Float64() - 0.5) * 4,
			vy:       (rand.Float64() - 0.5) * 4,
			lifetime: rand.Intn(30) + 20,
			color:    color.RGBA{255, 255, 0, 200},
		}
		g.particles = append(g.particles, particle)
	}
}

func (g *Game) generateFood() {
	// Generate food at random position, avoiding snake body
	maxX := ScreenWidth/GridSize - 1
	maxY := ScreenHeight/GridSize - 1
	
	for {
		g.food = Position{
			X: rand.Intn(maxX-1) + 1,
			Y: rand.Intn(maxY-1) + 1,
		}
		
		// Check that food is not on snake body
		onSnake := false
		for _, segment := range g.snake {
			if g.food.X == segment.X && g.food.Y == segment.Y {
				onSnake = true
				break
			}
		}
		if !onSnake {
			break
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
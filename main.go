package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20
)

type Point struct { // Структура для хранения координат
	X, Y int
}

type Game struct { // Структура для хранения состояния игры
	snake      []Point   // Змейка (массив точек)
	direction  Point     // Направление движения змейки
	food       Point     // Координаты еды
	gameOver   bool      // Флаг завершения игры
	lastUpdate time.Time // Время последнего обновления
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) { // Проверка нажатия клавиши R
			g.reset()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && g.direction.Y == 0 {
		g.direction = Point{0, -1} // Движение вверх
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.direction.Y == 0 {
		g.direction = Point{0, 1} // Движение вниз
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && g.direction.X == 0 {
		g.direction = Point{-1, 0} // Движение влево
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && g.direction.X == 0 {
		g.direction = Point{1, 0} // Движение вправо
	}

	if time.Since(g.lastUpdate).Milliseconds() > 100 { // Обновление позиции змейки каждые 100 миллисекунд
		head := g.snake[0]
		newHead := Point{head.X + g.direction.X, head.Y + g.direction.Y}

		if newHead.X < 0 || newHead.X >= screenWidth/gridSize || newHead.Y < 0 || newHead.Y >= screenHeight/gridSize { // Проверка на столкновение с границами экрана
			g.gameOver = true
			return nil
		}

		for _, p := range g.snake { // Проверка на столкновение с собой
			if p == newHead {
				g.gameOver = true
				return nil
			}
		}

		if newHead == g.food { // Проверка на съедание еды
			g.snake = append([]Point{newHead}, g.snake...) // Увеличиваем змейку
			g.placeFood()                                  // Размещаем новую еду
		} else {
			g.snake = append([]Point{newHead}, g.snake[:len(g.snake)-1]...) // Двигаем змейку
		}

		g.lastUpdate = time.Now()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{245, 245, 220, 255}) // Цвет фона

	for _, p := range g.snake { // Отрисовка змейки
		ebitenutil.DrawRect(screen, float64(p.X*gridSize), float64(p.Y*gridSize), gridSize, gridSize, color.RGBA{255, 165, 0, 255})
	}

	ebitenutil.DrawRect(screen, float64(g.food.X*gridSize), float64(g.food.Y*gridSize), gridSize, gridSize, color.RGBA{0, 128, 0, 255}) // Отрисовка еды

	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press R to restart.")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) placeFood() { // Размещает еду в случайной позиции
	rand.Seed(time.Now().UnixNano())
	g.food = Point{rand.Intn(screenWidth / gridSize), rand.Intn(screenHeight / gridSize)}

	for _, p := range g.snake { // Проверка, чтобы еда не появилась на змейке
		if p == g.food {
			g.placeFood()
			return
		}
	}
}

func (g *Game) reset() {
	g.snake = []Point{{5, 5}, {4, 5}, {3, 5}} // Начальная позиция змейки
	g.direction = Point{1, 0}                 // Начальное направление (вправо)
	g.gameOver = false
	g.placeFood() // Размещаем еду
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{ // Инициализация игры
		snake:     []Point{{5, 5}, {4, 5}, {3, 5}},
		direction: Point{1, 0},
		gameOver:  false,
	}
	game.placeFood()

	ebiten.SetWindowSize(screenWidth, screenHeight) // Размеры окна
	ebiten.SetWindowTitle("Змейка")                 // Название окна

	if err := ebiten.RunGame(game); err != nil { // Запуск игры
		log.Fatal(err)
	}
}

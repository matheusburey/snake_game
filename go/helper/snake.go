package helper

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var snake_color = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func DrawSnake(screen *ebiten.Image, snake [][2]float64) {
	snake_segment := ebiten.NewImage(10, 10)
	snake_segment.Fill(snake_color)
	for _, point := range snake {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(point[0]), float64(point[1]))
		screen.DrawImage(snake_segment, op)
	}
}

func MoveSnake(snake [][2]float64, direction string) [][2]float64 {
	head := (snake)[0]
	var newHead [2]float64

	switch direction {
	case "UP":
		newHead = [2]float64{head[0], head[1] - 10}
	case "DOWN":
		newHead = [2]float64{head[0], head[1] + 10}
	case "RIGHT":
		newHead = [2]float64{head[0] + 10, head[1]}
	case "LEFT":
		newHead = [2]float64{head[0] - 10, head[1]}
	}
	newSnake := append([][2]float64{newHead}, snake[:len(snake)-1]...)
	return newSnake
}

func ChangeDirection(keys []ebiten.Key, curDirection string) string {
	direction := curDirection

	for _, p := range keys {
		switch p {
		case ebiten.KeyArrowUp, ebiten.KeyW:
			if curDirection != "DOWN" {
				direction = "UP"
			}
		case ebiten.KeyArrowDown, ebiten.KeyS:
			if curDirection != "UP" {
				direction = "DOWN"
			}
		case ebiten.KeyArrowRight, ebiten.KeyD:
			if curDirection != "LEFT" {
				direction = "RIGHT"
			}
		case ebiten.KeyArrowLeft, ebiten.KeyA:
			if curDirection != "RIGHT" {
				direction = "LEFT"
			}
		}
	}

	return direction
}

func CheckCollisionWithSelf(snake [][2]float64) bool {
	for i := 1; i < len(snake); i++ {
		if snake[0][0] == snake[i][0] && snake[0][1] == snake[i][1] {
			return true
		}
	}
	return false
}

func CheckCollisionWithWall(snake [][2]float64) bool {
	if snake[0][0] < 0 || snake[0][0] > 590 || snake[0][1] < 0 || snake[0][1] > 590 {
		return true
	}
	return false
}

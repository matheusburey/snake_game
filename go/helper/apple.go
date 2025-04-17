package helper

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var apple_color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

func DrawApple(screen *ebiten.Image, pos [2]float64) {
	apple := ebiten.NewImage(10, 10)
	apple.Fill(apple_color)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos[0], pos[1])
	screen.DrawImage(apple, op)
}

func CheckCollisionWithApple(snake [][2]float64, apple [2]float64) bool {
	return snake[0][0] == apple[0] && snake[0][1] == apple[1]
}

func GenerateNewApple() [2]float64 {
	return [2]float64{float64(rand.Intn(59) * 10), float64(rand.Intn(59) * 10)}
}

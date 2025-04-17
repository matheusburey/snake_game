package main

import (
	"bytes"
	"image/color"
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/matheusburey/snake_game/helper"
)

const (
	screenWidth   = 600
	screenHeight  = 600
	fontSize      = 24
	titleFontSize = fontSize * 1.5
)

var (
	snakeStart = [][2]float64{
		{120, 50},
		{130, 50},
		{140, 50},
		{150, 50},
	}
	direction       = "LEFT"
	snake           [][2]float64
	apple           = [2]float64{200, 200}
	score           = 0
	mode            = "init"
	mplusFaceSource *text.GoTextFaceSource
)

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	time.Sleep(time.Second / 10)
	switch mode {
	case "init":
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			snake = make([][2]float64, len(snakeStart))
			copy(snake, snakeStart)
			score = 0
			mode = "play"
		}
	case "play":

		snake = helper.MoveSnake(snake, direction)
		g.keys = inpututil.AppendPressedKeys(g.keys[:0])
		direction = helper.ChangeDirection(g.keys, direction)

		if helper.CheckCollisionWithApple(snake, apple) {
			score += 10
			apple = helper.GenerateNewApple()
			snake = append(snake, snake[len(snake)-1])
		}

		if helper.CheckCollisionWithSelf(snake) || helper.CheckCollisionWithWall(snake) {
			mode = "gameover"
		}
	case "gameover":
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			mode = "init"
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch mode {
	case "init":
		titleTexts := "SNAKE\n\n press space to start"
		op := &text.DrawOptions{}
		op.GeoM.Translate(screenWidth/2, 3*titleFontSize)
		op.ColorScale.ScaleWithColor(color.White)
		op.LineSpacing = titleFontSize
		op.PrimaryAlign = text.AlignCenter
		text.Draw(screen, titleTexts, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   titleFontSize,
		}, op)
	case "play":
		helper.DrawApple(screen, apple)
		helper.DrawSnake(screen, snake)
	case "gameover":
		titleTexts := "Game Over\n\n press space to restart"
		op := &text.DrawOptions{}
		op.GeoM.Translate(screenWidth/2, 3*titleFontSize)
		op.ColorScale.ScaleWithColor(color.White)
		op.LineSpacing = titleFontSize
		op.PrimaryAlign = text.AlignCenter
		// converte score para string
		text.Draw(screen, titleTexts, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   titleFontSize,
		}, op)
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(500, titleFontSize*0.2)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	// converte score para string
	scoreString := "Score: " + strconv.Itoa(score)
	text.Draw(screen, scoreString, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   titleFontSize,
	}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

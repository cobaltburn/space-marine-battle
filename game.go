package main

import (
	"image/color"
	"slices"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth = 640
	ScreenHeight

	frameOX = 0
	frameOY = 32
	frameWidth
	frameHeight
	frameCount = 8
	TextOffSet = 4
)

type Game struct {
	marines  []Unit
	tyranids []Unit
	counter  int
}

func (game *Game) Update() error {
	if len(game.tyranids) != 0 && len(game.marines) != 0 {
		game.playTurn()
		time.Sleep(time.Second / 8)
	}
	game.counter++
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.drawBoard(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (game *Game) drawBoard(screen *ebiten.Image) {
	if len(game.marines) == 0 {
		winner := "tyranids win"
		ebitenutil.DebugPrintAt(screen, winner, ScreenWidth/2-len(winner)*TextOffSet, ScreenHeight/2)
		return
	} else if len(game.tyranids) == 0 {
		winner := "SpaceMarines win"
		ebitenutil.DebugPrintAt(screen, winner, ScreenWidth/2-len(winner)*TextOffSet, ScreenHeight/2)
		return
	}

	tileHeight := ScreenHeight / HEIGHT
	tileWidth := ScreenWidth / WIDTH

	for y := 0; HEIGHT > y; y++ {
		for x := 0; WIDTH > x; x++ {
			xPos := x * tileWidth
			yPos := y * tileHeight
			vector.DrawFilledRect(screen, float32(xPos), float32(yPos), float32(tileWidth), float32(tileHeight), game.checkColor(x, y), true)
		}
	}
}

func (game *Game) checkColor(x, y int) color.Color {
	for _, m := range game.marines {
		mx, my := m.GetPoint()
		if mx == x && my == y {
			return m.GetColor()
		}
	}
	for _, m := range game.tyranids {
		mx, my := m.GetPoint()
		if mx == x && my == y {
			return m.GetColor()
		}
	}
	return color.White
}

func (game *Game) playTurn() {
	switch i := game.counter; {
	case i%4 == 0:
		for i := 0; i < len(game.tyranids); i++ {
			game.tyranids[i].Attack(game.marines)
		}
		game.marines = slices.DeleteFunc(game.marines, func(e Unit) bool { return e.hp <= 0 })
	case i%3 == 0:
		for i := 0; i < len(game.tyranids); i++ {
			game.tyranids[i].Move(game.marines, game.tyranids)
		}
	case i%2 == 0:
		for i := 0; i < len(game.marines); i++ {
			game.marines[i].Attack(game.tyranids)
		}
		game.tyranids = slices.DeleteFunc(game.tyranids, func(e Unit) bool { return e.hp <= 0 })
	default:
		for i := 0; i < len(game.marines); i++ {
			game.marines[i].Move(game.tyranids, game.marines)
		}
	}
}

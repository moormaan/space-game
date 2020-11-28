package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	lastUpdateTime time.Time
	angle          float64
}

var game Game
var shipImg *ebiten.Image

func init() {
	var err error
	shipImg, _, err = ebitenutil.NewImageFromFile("resources/fighter.png")
	if err != nil {
		log.Fatal(err)
	}

	game.lastUpdateTime = time.Now()
}

func (g *Game) Update() error {
	newTime := time.Now()
	elapsed := float64(newTime.Sub(g.lastUpdateTime)) / 3000000000
	g.lastUpdateTime = newTime
	g.angle = g.angle + elapsed
	return nil
}

func (g *Game) DrawCursor(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	var path vector.Path
	fillOpts := vector.FillOptions{Color: color.White}
	path.MoveTo(float32(x-1), float32(y-1))
	path.LineTo(float32(x-1), float32(y+1))
	path.LineTo(float32(x+1), float32(y+1))
	path.LineTo(float32(x+1), float32(y-1))
	path.Fill(screen, &fillOpts)
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f\n"+
		"x:%d y:%d\n"+
		"Enter:%t", ebiten.CurrentFPS(), x, y, ebiten.IsKeyPressed(ebiten.KeyEnter)))

	g.DrawCursor(screen)
	shipGeom := ebiten.GeoM{}
	shipGeom.Rotate(g.angle)
	shipGeom.Scale(0.2, 0.2)
	shipGeom.Translate(float64(x), float64(y))
	shipOpts := &ebiten.DrawImageOptions{GeoM: shipGeom}
	screen.DrawImage(shipImg, shipOpts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	// ebiten.SetWindowSize(640, 480)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

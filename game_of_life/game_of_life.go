package main

import (
	"log"
	// "image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth = 400
	screenHeight = 400
	pixelSize = 4
)

type Game struct{
	state  []bool
	pixels []byte

	width  int
	height int
}

func (g *Game) Update() error {
	newState := make([]bool, len(g.state))
	for i := range g.state {
		row := i / g.width
		col := i - (row * g.width)
		
		dx := [8]int{0, 0, 1, 1, 1, -1, -1, -1}
		dy := [8]int{1, -1, 0, 1, -1, 0, 1, -1}

		neighs := 0
		for i := 0; i < 8; i++ {
			nrow := row + dx[i]
			ncol := col + dy[i]

			row_ok := nrow >= 0 && nrow < g.height
			col_ok := ncol >= 0 && ncol < g.width
			if row_ok && col_ok && g.state[nrow * g.width + ncol] {
				neighs++
			}
		}

		if g.state[i] == true && (neighs >= 4 || neighs <= 1) {
			newState[i] = false
		} else if g.state[i] == false && (neighs == 3) {
			newState[i] = true
		} else {
			newState[i] = g.state[i]
		}
	}
	g.state = newState
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < g.width * g.height; i++ {
		if g.state[i] == true {
			g.pixels[4 * i + 0] = 0xFF;
			g.pixels[4 * i + 1] = 0xFF;
			g.pixels[4 * i + 2] = 0xFF;
			g.pixels[4 * i + 3] = 0xFF;
		} else {
			g.pixels[4 * i + 0] = 0x00;
			g.pixels[4 * i + 1] = 0x00;
			g.pixels[4 * i + 2] = 0x00;
			g.pixels[4 * i + 3] = 0x00;
		}
	}
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")
	ebiten.SetTPS(15)

	state := make([]bool, screenWidth * screenHeight)

	for i := range state {
		if rand.Int() % 2 == 1 {
			state[i] = true
		} else {
			state[i] = false
		}
	}

	scale := 20
	width := screenWidth / scale
	height := screenHeight / scale
	
	g := Game{
		state: state,
		pixels: make([]byte, width * height * 4),
		width: width,
		height: height,
	}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"tetris/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.New()
	ebiten.SetWindowSize(240, 400)
	ebiten.SetWindowTitle("Tetris - Go + Ebiten")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

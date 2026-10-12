package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"cronometro-futsal/internal/ui"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Cronometro de Futsal")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(ui.NewGame()); err != nil {
		log.Fatal(err)
	}
}
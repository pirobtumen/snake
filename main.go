package main

import (
	"log"
	"snake/core"
	"snake/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	settings := core.NewSettings()
	game := ui.NewGame(settings)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

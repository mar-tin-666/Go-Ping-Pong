package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(titleText)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

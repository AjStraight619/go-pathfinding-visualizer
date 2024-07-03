package main

import (
	"log"

	"github.com/AJStraight619/go-pathfinding-visualizer/internal/pathfinding"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	p := pathfinding.InitPathfinding()

	ebiten.SetWindowSize(1200, 900)
	ebiten.SetWindowTitle("Pathfinding Visualizer")
	if err := ebiten.RunGame(p); err != nil {
		log.Fatal(err)
	}

}

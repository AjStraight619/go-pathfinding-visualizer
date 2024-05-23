package main

import (
	"fmt"

	"github.com/AjStraight619/pathfinding-go/internal/grid"
)

func findStartAndFinishNodes(g *grid.Grid) (*grid.Node, *grid.Node) {
	var startNode, finishNode *grid.Node
	for i := range g.Nodes {
		for j := range g.Nodes[i] {
			if g.Nodes[i][j].IsStart {
				startNode = &g.Nodes[i][j]
			}
			if g.Nodes[i][j].IsFinish {
				finishNode = &g.Nodes[i][j]
			}
		}
	}
	return startNode, finishNode
}

func main() {
	grid := grid.NewGrid(30, 30)

	startNode, finishNode := findStartAndFinishNodes(grid)
	if startNode == nil || finishNode == nil {
		fmt.Println("could not find start or finish node")
		return
	}

	fmt.Printf("Start node: %+v\n", *startNode)
	fmt.Printf("Finish node: %+v\n", *finishNode)

}

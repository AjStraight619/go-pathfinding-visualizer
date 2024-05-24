package algorthims

import (
	"github.com/AjStraight619/pathfinding-go/internal/grid"
	"github.com/AjStraight619/pathfinding-go/internal/pq"
)

func Astar(grid *grid.Grid, startNode *grid.Node, finishNode *grid.Node) {
	// Your A* implementation here
	openSet := pq.New()
	openSet.Push(startNode)

}

func heuristic(startNode *grid.Node, finishNode *grid.Node) int {
	dx := abs(startNode.Row - finishNode.Row)
	dy := abs(startNode.Col - finishNode.Col)
	return dx + dy
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

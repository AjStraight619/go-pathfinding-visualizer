package grid

// Represents individual nodes in the grid
type Node struct {
	Row, Col  int
	Walkable  bool
	G, H, F   int
	Parent    *Node
	IsStart   bool
	IsFinish  bool
	isVisited bool
}

type Grid struct {
	Cols, Rows int
	Nodes      [][]Node
}

// Initializing grid

func NewGrid(cols, rows int) *Grid {
	nodes := make([][]Node, rows)
	for i := range nodes {
		nodes[i] = make([]Node, cols)
		for j := range nodes[i] {
			nodes[i][j] = Node{
				Row:       i,
				Col:       j,
				Walkable:  true,
				Parent:    nil,
				IsStart:   i == 0 && j == 0,
				IsFinish:  i == rows-1 && j == cols-1,
				isVisited: false,
			}

			if nodes[i][j].IsStart {
				nodes[i][j].Walkable = false
			}
		}
	}

	return &Grid{
		Cols:  cols,
		Rows:  rows,
		Nodes: nodes,
	}
}

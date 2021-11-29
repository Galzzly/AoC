package main

import "github.com/Galzzly/AoC/2020/utils"

func (t *Tile) allEdge() []string {
	edges := []string{
		/*top*/ t.top(),
		/*bottom*/ t.bottom(),
		/*left*/ t.left(),
		/*right*/ t.right(),
	}

	return []string{
		edges[0],
		utils.Reverse(edges[0]),
		edges[1],
		utils.Reverse(edges[1]),
		edges[2],
		utils.Reverse(edges[2]),
		edges[3],
		utils.Reverse(edges[3]),
	}
}

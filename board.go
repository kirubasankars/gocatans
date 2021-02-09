package main

import (
	"fmt"
	"strings"
)

type Board struct {
	grid *Grid
}

func (board Board) GetNeighborIntersections1(intersection int) []int {
	thisIntersection := board.grid.intersections[intersection]
	neighborIntersections := thisIntersection.neighbors
	var output = make([]int, len(thisIntersection.neighbors))
	for _, ins := range neighborIntersections {
		t := ins.index
		output = append(output, t)
	}
	return output
}

func (board Board) GetNeighborIntersections2(intersection int) [][2]int {
	thisIntersection := board.grid.intersections[intersection]
	neighborIntersections := thisIntersection.neighbors
	var output = make([][2]int, len(thisIntersection.neighbors))
	for _, ins := range neighborIntersections {
		if ins.index < intersection {
			output = append(output, [2]int{ins.index, intersection})
		} else {
			output = append(output, [2]int{intersection, ins.index})
		}
	}
	return output
}

func (board Board) GetAvailableIntersections(occupied []int) []int {
	intersections := board.grid.intersections
	l := len(occupied)
	for i := 0; i < l; i++ {
		occupiedNeighbors := intersections[occupied[i]].neighbors
		for _, nins := range occupiedNeighbors {
			occupied = append(occupied, nins.index)
		}
	}
	keys := make([]int, 0, len(intersections))
	for k := range intersections {
		if !Contains(occupied, k) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (board Board) GetTileIndices(intersection int) []int {
	coordinator := board.grid.intersections[intersection]
	var indices = make([]int, len(coordinator.nodes))
	for idx, n := range coordinator.nodes {
		indices[idx] = n.index
	}
	return indices
}

func (board Board) GetTiles() [][2]int {
	var tiles = make([][2]int, len(board.grid.nodes))
	var tIndex = 0
	for idx, n := range board.grid.nodes {
		var rt = -1
		var token = -1
		switch n.resource {
		case "t":
			rt = 0
		case "h":
			rt = 1
		case "p":
			rt = 2
		case "f":
			rt = 3
		case "m":
			rt = 4
		default:
			rt = -1
		}
		if rt != -1 {
			token = n.token
			tIndex++
		}
		tiles[idx] = [2]int{rt, token}
	}
	return tiles
}

func (board Board) JSON() string {
	//tiles := board.GetTiles()
	//makeIntersections := func(l []*Intersection) string {
	//	var nodes []string
	//	for _, h := range l {
	//		nodes = append(nodes, fmt.Sprintf("{index:%d,x:%f,y:%f, hasPort:%t, portResource:'%s'}", h.index, h.x, h.y, h.hasPort, h.portResource))
	//	}
	//	return "[" + strings.Join(nodes, ",") + "]"
	//}
	var nodes []string
	for _, h := range board.grid.nodes {
		if h.resource == "-" || h.resource == "s" {
			continue
		}
		nodes = append(nodes, fmt.Sprintf("{index:%d,x:%f,y:%f,r:%f}", h.index, h.x, h.y,h.r))
	}
	return "{'nodes': [" + strings.Join(nodes, ",") + "]}"
}

func NewBoard(ID int) Board {
	var grid = new(Grid)
	if ID == 0 {
		m := DefaultMap{}
		grid.Build(m)
	}
	if ID == 1 {
		m := SmallMap{}
		grid.Build(m)
	}
	board := new(Board)
	board.grid = grid
	fmt.Println(board.JSON())
	return *board
}

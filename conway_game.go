package main

import (
	"conway-v2/game_utils"
	"conway-v2/terminal_utils"
)

type ConwayGame struct {
	height      int
	width       int
	cells       [][]int
	gridOffsetH int
	gridOffsetW int
	color       terminal_utils.TerminalColor
}

func NewConwayGame(height int, width int, gridOffsetH int, gridOffsetW int, color terminal_utils.TerminalColor) *ConwayGame {
	cells := make([][]int, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]int, width)
	}

	return &ConwayGame{
		height:      height,
		width:       width,
		cells:       cells,
		gridOffsetH: gridOffsetH,
		gridOffsetW: gridOffsetW,
		color:       color,
	}
}

func (game *ConwayGame) PlayCycle() {
	terminal_utils.TerminalSetColor(game.color)

	cells_cp := make([][]int, game.height)
	for i := 0; i < game.height; i++ {
		cells_cp[i] = make([]int, game.width)
	}
	for r := 0; r < game.height; r++ {
		for c := 0; c < game.width; c++ {
			neighbors := game_utils.CalcNumNeighbors(r, c, game.height, game.width, &game.cells)
			state := game_utils.CalcCurrentState(neighbors, game.cells[r][c])

			if state != game.cells[r][c] {
				row := r + (game.height * game.gridOffsetH)
				col := c + (game.width * game.gridOffsetW)
				if state == 1 {
					terminal_utils.TerminalWrite(row, col, "█")
				} else {

					terminal_utils.TerminalWrite(row, col, " ")
				}
			}

			cells_cp[r][c] = state
		}
	}

	game.cells = cells_cp
}

func (game *ConwayGame) RandomlyPopulate() {
	game_utils.RandomlyPopulateCells(game.height, game.width, &game.cells)
}

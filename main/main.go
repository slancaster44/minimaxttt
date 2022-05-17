package main

import (
	"tictactoe"
)

func main() {
	tictactoe.DoGame(tictactoe.MinimaxPlayer, tictactoe.HumanPlayer)
}

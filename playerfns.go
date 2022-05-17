package tictactoe

import (
	"fmt"
	"math"
)

type mkMoveFn func(Board, byte) Board

func HumanPlayer(b Board, playerPiece byte) Board {
	var val int

	fmt.Print("Enter a number 0-8: ")
	fmt.Scanf("%d", &val)

	if b[val] == EMPTY {
		b[val] = playerPiece
	} else {
		fmt.Println("Try Again!")
		b = HumanPlayer(b, playerPiece)
	}
	return b
}

func MinimaxPlayer(b Board, playerPiece byte) Board {

	playPosition := 0
	value := int(math.Inf(-1))

	for iter, item := range b {
		if item == EMPTY {

			b[iter] = playerPiece
			score := Minimax(b, InversePlayer(playerPiece))
			if score > value {
				playPosition = iter
				value = score
			}

			b[iter] = EMPTY
		}
	}

	b[playPosition] = playerPiece
	return b
}

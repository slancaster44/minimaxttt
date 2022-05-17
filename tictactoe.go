package tictactoe

import "fmt"

type Board [9]byte

const (
	EMPTY byte = iota
	O
	X
)

func InversePlayer(player byte) byte {
	if player == X {
		return O
	}
	return X
}

func NewBoard() Board {
	b := Board{}
	for iter := range b {
		b[iter] = EMPTY
	}

	return b
}

func (b Board) Print() {
	fmt.Println("\n+---+---+---+")

	for iter, item := range b {
		if iter%3 == 2 {
			fmt.Print("| ")
			printItem(item)
			fmt.Print(" |")
			fmt.Println("\n+---+---+---+")
		} else {
			fmt.Print("| ")
			printItem(item)
			fmt.Print(" ")
		}
	}
}

func printItem(item byte) {
	switch item {
	case X:
		fmt.Print("X")
	case O:
		fmt.Print("O")
	default:
		fmt.Print(" ")
	}
}

package tictactoe

import "math"

func Minimax(b Board, playerPiece byte) int {
	analysis := AnalyzeBoard(b)

	if analysis.WhoIsWinner == InversePlayer(playerPiece) {
		return analysis.NumMovesFor(InversePlayer(playerPiece))
	}

	value := int(math.Inf(-1))

	for iter, item := range b {
		if item == EMPTY {

			b[iter] = playerPiece

			//Scores lowest on shortest combination that causes other player to win.
			score := -Minimax(b, InversePlayer(playerPiece))

			if score > value {
				value = score
			}
			b[iter] = EMPTY
		}
	}

	return value
}

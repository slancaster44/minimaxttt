package tictactoe

type BoardAnalysis struct {
	IsThereWinner bool
	IsGameOver    bool
	WhoIsWinner   byte
	EmptySpaces   int
	Xspaces       int
	Ospaces       int
}

func AnalyzeBoard(b Board) BoardAnalysis {
	analysis := BoardAnalysis{}

	winner := getWinner(b)
	analysis.IsThereWinner = winner != EMPTY
	analysis.WhoIsWinner = winner

	analysis.IsGameOver = boardFull(b) || analysis.IsThereWinner

	analysis.EmptySpaces = 0
	analysis.Xspaces = 0
	analysis.Ospaces = 0

	for _, item := range b {
		switch item {
		case X:
			analysis.Xspaces++
		case O:
			analysis.Xspaces++
		case EMPTY:
			analysis.EmptySpaces++
		}
	}

	return analysis
}

func getWinner(b Board) byte {

	//This list contains values that represent the indicies that must match
	//in order for a win to happen
	tests := [8][3]int{{0, 4, 8}, {6, 4, 2}, {0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}}

	for _, test := range tests {
		if b[test[0]] == b[test[1]] && b[test[1]] == b[test[2]] && b[test[0]] != EMPTY {
			return b[test[0]]
		}
	}

	return EMPTY
}

func (a BoardAnalysis) NumMovesFor(playerPiece byte) int {
	if playerPiece == X {
		return a.Xspaces
	}
	return a.Ospaces
}

func boardFull(b Board) bool {
	for _, item := range b {
		if item == EMPTY {
			return false
		}
	}
	return true
}

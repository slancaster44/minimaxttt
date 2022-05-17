package tictactoe

func DoGame(xPlayerfn, oPlayerfn mkMoveFn) {
	b := NewBoard()
	analysis := AnalyzeBoard(b)

	curPlayer := X
	for !analysis.IsGameOver {
		b.Print()

		if curPlayer == X {
			b = xPlayerfn(b, curPlayer)
		} else {
			b = oPlayerfn(b, curPlayer)
		}

		curPlayer = InversePlayer(curPlayer)
		analysis = AnalyzeBoard(b)
	}
	b.Print()
}

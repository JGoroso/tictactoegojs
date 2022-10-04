package board

import (
	"fmt"
)

var (
	GAMECHOOSE   = [9]string{"", "", "", "", "", "", "", "", ""}
	POSIBILITIES = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	gameActive   = true
	isDraw       = false
)

type Round struct {
	GameOn bool   `json:"gameon"`
	Player string `json:"player"`
	Draw   bool   `json:"isdraw"`
}

func GameFunc(player string, positionIndex int, restart bool) Round {

	addChoose(player, positionIndex)

	GameWin()
	isRestart(restart)
	lookForDraw()

	var result Round
	result.Player = player
	result.GameOn = gameActive
	result.Draw = isDraw
	return result
}

func GameWin() {
	haveWinner := false
	for x := range POSIBILITIES {
		winningPosibilities := POSIBILITIES[x]
		firstPosition := GAMECHOOSE[winningPosibilities[0]]
		secondPosition := GAMECHOOSE[winningPosibilities[1]]
		thirdPosition := GAMECHOOSE[winningPosibilities[2]]

		if firstPosition == "" || secondPosition == "" || thirdPosition == "" {
			continue
		}
		if firstPosition == secondPosition && secondPosition == thirdPosition {
			haveWinner = true
			break
		}
	}

	if haveWinner {
		gameActive = false
		GAMECHOOSE = [9]string{"", "", "", "", "", "", "", "", ""}
	}
}

func addChoose(player string, index int) {

	if GAMECHOOSE[index] == "" {
		GAMECHOOSE[index] = player
	} else if GAMECHOOSE[index] != "" {
		fmt.Println("elige otro lugar")
	}
}

func isRestart(restart bool) {
	if restart {
		gameActive = true
		isDraw = false
		GAMECHOOSE = [9]string{"", "", "", "", "", "", "", "", ""}
	}
}

func lookForDraw() {
	if GAMECHOOSE[0] != "" && GAMECHOOSE[1] != "" && GAMECHOOSE[2] != "" && GAMECHOOSE[3] != "" && GAMECHOOSE[4] != "" && GAMECHOOSE[5] != "" && GAMECHOOSE[6] != "" && GAMECHOOSE[7] != "" && GAMECHOOSE[8] != "" {
		isDraw = true
	}

}

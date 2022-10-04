package board

import (
	"fmt"
)

var (
	//Guardamos los movimientos
	GAMECHOOSE = [9]string{"", "", "", "", "", "", "", "", ""}
	//Array de posibilidades
	POSIBILITIES = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	gameActive   = true
	isDraw       = false
)

// Creamos estructura para luego retornarla y consumirla como json.
type Round struct {
	GameOn bool   `json:"gameon"`
	Player string `json:"player"`
	Draw   bool   `json:"isdraw"`
}

func GameFunc(player string, positionIndex int, restart bool) Round {
	// Asignar posiciones
	addChoose(player, positionIndex)
	// Realiza el chequeo de si existe un ganador
	GameWin()
	// Permite resetear el juego
	isRestart(restart)
	// Busca un posible empate
	lookForDraw()

	//Instanciamos variable con tipo de dato Round y asignamos los valores
	var result Round
	result.Player = player
	result.GameOn = gameActive
	result.Draw = isDraw
	//Retornamos info para consumir desde el frontend
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

package game

import (
	"fmt"
	"math/rand"
)

func setSeed(seed string, tablero *Juego) {

	defer func() {
		err := recover()
		if err != nil {
			panic(fmt.Sprintf("the minimum size of the board must be 30x40 %s", err))
		}
	}()

	switch seed {
	case "BIRD":
		tablero.tablero[6][12].value = viva
		tablero.tablero[7][10].value = viva
		tablero.tablero[7][14].value = viva

		tablero.tablero[8][15].value = viva

		tablero.tablero[9][10].value = viva
		tablero.tablero[9][15].value = viva

		tablero.tablero[10][11].value = viva
		tablero.tablero[10][12].value = viva
		tablero.tablero[10][13].value = viva
		tablero.tablero[10][14].value = viva
		tablero.tablero[10][15].value = viva

	case "INITIAL":
		tablero.tablero[15][30].value = viva
		tablero.tablero[15][31].value = viva
		tablero.tablero[16][31].value = viva
		tablero.tablero[16][32].value = viva
		tablero.tablero[17][31].value = viva

	case "MATUSALENES":
		tablero.tablero[5][25].value = viva

		tablero.tablero[6][23].value = viva
		tablero.tablero[6][25].value = viva

		tablero.tablero[7][13].value = viva
		tablero.tablero[7][14].value = viva
		tablero.tablero[7][21].value = viva
		tablero.tablero[7][22].value = viva
		tablero.tablero[7][35].value = viva
		tablero.tablero[7][36].value = viva

		tablero.tablero[8][12].value = viva
		tablero.tablero[8][16].value = viva
		tablero.tablero[8][21].value = viva
		tablero.tablero[8][22].value = viva
		tablero.tablero[8][35].value = viva
		tablero.tablero[8][36].value = viva

		tablero.tablero[9][1].value = viva
		tablero.tablero[9][2].value = viva
		tablero.tablero[9][11].value = viva
		tablero.tablero[9][17].value = viva
		tablero.tablero[9][21].value = viva
		tablero.tablero[9][22].value = viva

		tablero.tablero[10][1].value = viva
		tablero.tablero[10][2].value = viva
		tablero.tablero[10][11].value = viva
		tablero.tablero[10][15].value = viva
		tablero.tablero[10][17].value = viva
		tablero.tablero[10][18].value = viva
		tablero.tablero[10][23].value = viva
		tablero.tablero[10][25].value = viva

		tablero.tablero[11][11].value = viva
		tablero.tablero[11][17].value = viva
		tablero.tablero[11][25].value = viva

		tablero.tablero[12][12].value = viva
		tablero.tablero[12][16].value = viva

		tablero.tablero[13][13].value = viva
		tablero.tablero[13][14].value = viva

	case "RANDOM":

		for x := range tablero.tablero {
			for y := range tablero.tablero[x] {
				//rand.Seed(253)
				alive := rand.Intn(2)
				if alive == 1 {
					tablero.tablero[x][y].value = viva
				} else {
					tablero.tablero[x][y].value = muerta
				}
			}
		}
	}
}

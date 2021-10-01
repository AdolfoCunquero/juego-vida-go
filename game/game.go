package game

import (
	"fmt"
	"time"
)

type cell struct {
	value     string
	nextValue string
}

type tablero [][]cell
type Juego struct {
	tablero
}

var (
	viva    string
	muerta  string
	refresh time.Duration
	nrows   int
	ncols   int
)

var generation int
var population int

func (j *Juego) Print() {
	fmt.Println(j)
}

func (j *Juego) Fill(pRows int, pCols int, pRefresh time.Duration, pViva string, pMuerta string) {
	nrows = pRows
	ncols = pCols
	refresh = pRefresh
	viva = pViva
	muerta = pMuerta

	j.tablero = make([][]cell, nrows)
	for i := range j.tablero {
		j.tablero[i] = make([]cell, ncols)
	}

	for row, i := range j.tablero {
		for col := range i {
			j.tablero[row][col].value = muerta
		}
	}

}

func (j *Juego) Initialize(seed string) {
	setSeed(seed, j)
}

func (j *Juego) String() string {
	str := "\n"
	for _, v := range j.tablero {
		str += "["
		for _, w := range v {
			str += fmt.Sprintf("%v", w.value)
		}
		str += "]\n"
	}
	str += fmt.Sprintf("Generation: %d\n", generation)
	str += fmt.Sprintf("Population: %d\n", population)
	return str
}

func (j *Juego) Run() {
	for {
		j.Print()
		j.ValidateAlive()
		time.Sleep(refresh)
	}
}

func (j *Juego) ValidateAlive() {

	population = 0
	generation++

	for x, row := range j.tablero {
		for y := range row {
			aliveN := j.aliveNeighbors(x, y)
			flag := j.tablero[x][y].value

			if flag == muerta && aliveN == 3 {
				j.tablero[x][y].nextValue = viva
				population++
			} else if flag == viva && (aliveN < 2 || aliveN > 3) {
				j.tablero[x][y].nextValue = muerta
			} else {
				j.tablero[x][y].nextValue = j.tablero[x][y].value
				if j.tablero[x][y].value == viva {
					population++
				}
			}
		}
	}

	for x, row := range j.tablero {
		for y := range row {
			j.tablero[x][y].value = j.tablero[x][y].nextValue
		}
	}
}

func (j *Juego) aliveNeighbors(x int, y int) int {
	var aliveN int = 0

	if x == 0 && y == 0 {
		aliveN += j.aliveNRight(x, y)
		aliveN += j.aliveNBottom(x, y)
		aliveN += j.aliveNBottomRight(x, y)

	} else if x == 0 && y > 0 && y < (ncols-1) {
		aliveN += j.aliveNLeft(x, y)
		aliveN += j.aliveNBottom(x, y)
		aliveN += j.aliveNBottomLeft(x, y)
		aliveN += j.aliveNBottomRight(x, y)
		aliveN += j.aliveNRight(x, y)

	} else if x == 0 && y == (ncols-1) {
		aliveN += j.aliveNLeft(x, y)
		aliveN += j.aliveNBottomLeft(x, y)
		aliveN += j.aliveNBottom(x, y)

	} else if x > 0 && x < (nrows-1) && y == 0 {
		aliveN += j.aliveNTop(x, y)
		aliveN += j.aliveNTopRight(x, y)
		aliveN += j.aliveNRight(x, y)
		aliveN += j.aliveNBottomRight(x, y)
		aliveN += j.aliveNBottom(x, y)

	} else if x > 0 && x < (nrows-1) && y < (ncols-1) {
		aliveN += j.aliveNTop(x, y)
		aliveN += j.aliveNTopRight(x, y)
		aliveN += j.aliveNRight(x, y)
		aliveN += j.aliveNBottomRight(x, y)
		aliveN += j.aliveNBottom(x, y)
		aliveN += j.aliveNBottomLeft(x, y)
		aliveN += j.aliveNLeft(x, y)
		aliveN += j.aliveNTopLeft(x, y)

	} else if x > 0 && x < (nrows-1) && y == (ncols-1) {
		aliveN += j.aliveNTop(x, y)
		aliveN += j.aliveNTopLeft(x, y)
		aliveN += j.aliveNLeft(x, y)
		aliveN += j.aliveNBottomLeft(x, y)
		aliveN += j.aliveNBottom(x, y)

	} else if x == (nrows-1) && y == 0 {
		aliveN += j.aliveNTop(x, y)
		aliveN += j.aliveNTopRight(x, y)
		aliveN += j.aliveNRight(x, y)

	} else if x == (nrows-1) && y > 0 && y < (ncols-1) {
		aliveN += j.aliveNLeft(x, y)
		aliveN += j.aliveNTopLeft(x, y)
		aliveN += j.aliveNTop(x, y)
		aliveN += j.aliveNTopRight(x, y)
		aliveN += j.aliveNRight(x, y)

	} else {
		aliveN += j.aliveNLeft(x, y)
		aliveN += j.aliveNTopLeft(x, y)
		aliveN += j.aliveNTop(x, y)

	}

	return aliveN
}

func (j *Juego) aliveNRight(x int, y int) int {
	if j.tablero[x][y+1].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNLeft(x int, y int) int {
	if j.tablero[x][y-1].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNBottom(x int, y int) int {
	if j.tablero[x+1][y].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNTop(x int, y int) int {
	if j.tablero[x-1][y].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNTopRight(x int, y int) int {
	if j.tablero[x-1][y+1].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNTopLeft(x int, y int) int {
	if j.tablero[x-1][y-1].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNBottomRight(x int, y int) int {
	if j.tablero[x+1][y+1].value == viva {
		return 1
	}
	return 0
}

func (j *Juego) aliveNBottomLeft(x int, y int) int {
	if j.tablero[x+1][y-1].value == viva {
		return 1
	}
	return 0
}

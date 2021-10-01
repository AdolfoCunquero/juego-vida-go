package main

import (
	"fmt"
	"time"

	gm "github.com/AdolfoCunquero/juego-vida-go/game"
)

//AVIABLE SEEDS
//	BIRD
//	INITIAL
//	MATUSALENES
//	RANDOM

const (
	VIVA         = "▄ "
	MUERTA       = "  " //□
	REFRESH_MLLS = 50
	NROWS        = 30
	NCOLS        = 60
	SEED         = "MATUSALENES"
)

func main() {
	j := new(gm.Juego)

	refresh := time.Millisecond * time.Duration(REFRESH_MLLS)
	j.Fill(NROWS, NCOLS, refresh, VIVA, MUERTA)

	j.Initialize(SEED)

	go j.Run()

	var s string
	fmt.Scanf("%s", s)
}

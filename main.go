package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type World struct {
	Grid   [][]string
	Player struct {
		X int
		Y int
	}
}

const FieldSymbol = "#"
const PlayerSymbol = "P"

// Метод инициализации мира
func (w *World) Init(width int, height int) {
	w.Grid = make([][]string, width)
	for i := range w.Grid {
		w.Grid[i] = make([]string, height)
		for j := range w.Grid[i] {
			w.Grid[i][j] = FieldSymbol
		}
	}
	w.Player.X, w.Player.Y = 0, 0 // Начальная позиция игрока
	w.Grid[w.Player.X][w.Player.Y] = PlayerSymbol
}

func (w World) Print() {
	for _, row := range w.Grid {
		fmt.Println(row)
	}
}

func (w *World) MovePlayer(dx, dy int) {
	newX, newY := w.Player.X+dx, w.Player.Y+dy
	if newX >= 0 && newX < len(w.Grid) && newY >= 0 && newY < len(w.Grid[0]) {
		w.Grid[w.Player.X][w.Player.Y] = FieldSymbol
		w.Player.X, w.Player.Y = newX, newY
		w.Grid[w.Player.X][w.Player.Y] = PlayerSymbol
	}
}

// Игровой цикл
func GameLoop(w *World) {
	reader := bufio.NewReader(os.Stdin)
	for {
		w.Print()
		fmt.Print("Enter a command (UP, DOWN, LEFT, RIGHT, EXIT) ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		switch input {
		case "UP":
			w.MovePlayer(-1, 0)
		case "DOWN":
			w.MovePlayer(1, 0)
		case "LEFT":
			w.MovePlayer(0, -1)
		case "RIGHT":
			w.MovePlayer(0, 1)
		case "EXIT":
			fmt.Println("Quit a game.")
			return
		default:
			fmt.Println("Unknown command. Try again.")
		}
	}
}

func main() {
	var world World
	world.Init(5, 5)
	GameLoop(&world)
}

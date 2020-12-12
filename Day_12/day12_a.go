package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	//	--- STANDARD FILE OPENING ---
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

	// --- START ---
	sc := bufio.NewScanner(input)

	var (
		northSouth 	int
		eastWest	int
		direction	int
	)

	const (
		E = 0
		S = 1
		W = 2
		N = 3
	)

	for sc.Scan() {
		units, _ := strconv.Atoi(sc.Text()[1:])
		if sc.Text()[0] == 'S' || sc.Text()[0] == 'W' || sc.Text()[0] == 'F' && (direction == S || direction == W){
			units *= -1
		}

		switch sc.Text()[0] {
		case 'N', 'S':
			northSouth += units
		case 'E', 'W':
			eastWest += units
		case 'F':
			switch direction  {
			case E, W:
				eastWest += units
			case N, S:
				northSouth += units
			}
		case 'R':
			direction = (direction + units/90)%4
		case 'L':
			direction = (direction - units/90 + 4)%4
		}
	}
	if northSouth < 0{
		northSouth *= -1
	}
	if eastWest < 0{
		eastWest *= -1
	}
	fmt.Println(northSouth+eastWest)
}
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
		northSouth 		int
		eastWest		int
		WpNorthSouth	int	= 1
		WpEastWest		int	= 10
	)

	for sc.Scan() {
		units, _ := strconv.Atoi(sc.Text()[1:])
		if sc.Text()[0] == 'S' || sc.Text()[0] == 'W'{
			units *= -1
		}

		switch sc.Text()[0] {
		case 'N', 'S':
			WpNorthSouth += units
		case 'E', 'W':
			WpEastWest 	+= units
		case 'F':
			northSouth 	+= units * WpNorthSouth
			eastWest 	+= units * WpEastWest
		case 'R':
			for i:=0; i<units/90; i++{
				WpEastWest, WpNorthSouth = WpNorthSouth, -WpEastWest
			}
		case 'L':
			for i:=0; i<units/90; i++{
				WpEastWest, WpNorthSouth = -WpNorthSouth, WpEastWest
			}
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
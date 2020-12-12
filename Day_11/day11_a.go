package main

import (
	"fmt"
	"os"
	"bufio"
)

type seat struct{
	state		rune
	adjacents	int
}

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
	var waitingArea [][]seat
	
	for sc.Scan(){
		var waitingLine []seat
		for _, state := range sc.Text(){
			waitingLine = append(waitingLine, seat{state, 0})
		}
		waitingArea = append(waitingArea, waitingLine)
	}
	for{
	//	show(waitingArea)
		waitingArea = countAdjacents(waitingArea)
		change, occupied := rule(waitingArea)
		if !change{
			fmt.Println(occupied)
			break
		}
	}
}

func countAdjacents(waitingArea [][]seat)[][]seat{
	for i, waitingLine := range waitingArea{
		for j := range waitingLine{
			waitingArea[i][j].adjacents = 0
			for x:=-1; x<=1; x++{
				for y:=-1; y<=1; y++{
					if i+y >= len(waitingArea) || j+x >= len(waitingLine) || i+y < 0 || j+x < 0{
						continue
					}
					if (x!=0 || y!=0) && waitingArea[i+y][j+x].state == '#' {
						waitingArea[i][j].adjacents++
					}
				}
			}
		}
	}
	return waitingArea
}

func rule(waitingArea [][]seat)(change bool, occupied int){
	for i, waitingLine := range waitingArea{
		for j, seat := range waitingLine{
			if seat.state == 'L' && seat.adjacents == 0{
				waitingArea[i][j].state = '#'
				change = true
			}else if seat.state == '#' && seat.adjacents>3{
				waitingArea[i][j].state = 'L'
				change = true
			}
			if waitingArea[i][j].state == '#'{
				occupied++
			}
		}
	}
	return
}
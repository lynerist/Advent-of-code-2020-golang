package main

import (
	"fmt"
	"os"
	"bufio"
)

type cube struct{
 	active		bool
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

	poketDimension := make(map[string]cube)
	
	var countStartY int
	for sc.Scan(){
		for x, active := range sc.Text(){
			key := fmt.Sprintf("%d,%d,0,0",x,countStartY)
			if active == '#'{
				poketDimension[key] = cube{true, 0}
			}
		}
		countStartY++
	}

	for cycle :=0; cycle<6; cycle++{
		countNeighbours(poketDimension)
		changeState(poketDimension)
	}
	fmt.Println(len(poketDimension))
}

func countNeighbours(poketDimension map[string]cube){
	for coordinates, cube := range poketDimension{
		if !cube.active{
			continue
		}
		var x,y,z,w int
		fmt.Sscanf(coordinates,"%d,%d,%d,%d",&x,&y,&z,&w)
		
		for i:=-1; i<=1; i++{
			for j:=-1; j<=1; j++{
				for k:=-1; k<=1; k++{
					for l:=-1; l<=1; l++{
						adjacent := fmt.Sprintf("%d,%d,%d,%d",x+i,y+j,z+k, w+l)
						adjacentCube := poketDimension[adjacent]
						if adjacent != coordinates {
							adjacentCube.adjacents++
							poketDimension[adjacent] = adjacentCube
						}
					}
				}
			}
		}
	}
}

func changeState(poketDimension map[string]cube){
	for coordinates, currentCube := range poketDimension{
		if currentCube.active && (currentCube.adjacents < 2 || currentCube.adjacents > 3) || !currentCube.active && currentCube.adjacents != 3{
			delete(poketDimension, coordinates)
		}else{
			poketDimension[coordinates] = cube{true, 0}
		}
	}
}
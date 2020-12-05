package main

import(
	"os"
	"bufio"
	"fmt"
)

func main(){
//	--- STANDARD FILE OPENING ---
	input, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

// --- START --- 
	sc := bufio.NewScanner(input)

	var (
		rightSteps, trees [5]int
		line int
	)

	for sc.Scan(){
		for slope := 0; slope<5; slope++{
			if slope==4 && line%2 != 0{ //This is used to let the last slope skip odd lines
				continue
			}
			if sc.Text()[rightSteps[slope] % len(sc.Text())] == '#'{ //I use module operator to prevent overflow
				trees[slope]++
			}
			rightSteps[slope] += (2*slope+1)%8 //This full of magic numbers expression gives the sequence [1 3 5 7 1]
		}
		line++ 
	}
	fmt.Println(trees[0]*trees[1]*trees[2]*trees[3]*trees[4])
}
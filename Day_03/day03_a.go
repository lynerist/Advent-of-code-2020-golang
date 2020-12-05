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
	var rightSteps, trees int
	for sc.Scan(){
		if sc.Text()[rightSteps % len(sc.Text())] == '#'{ //I use module operator to prevent overflow
			trees++
		}
		rightSteps += 3
	}
	fmt.Println(trees)
}
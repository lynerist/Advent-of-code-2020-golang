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

//	--- START --
	sc:= bufio.NewScanner(input)
	Groupforms := make(map[rune]bool)		//Simulation of a Set to save the answers of a group
	var countAnswers int

	for {
		end := !sc.Scan()
		
		if sc.Text() == "" || end {			//I have to do this to process the last input
			countAnswers += len(Groupforms)
			Groupforms = make(map[rune]bool)//Clear the set for the next group
		}else{
			for _, answer := range sc.Text(){
				Groupforms[answer] = true	//add something to a set is an idempotent operation
			}
		}
		
		if end {
			break
		}
	}
	fmt.Println(countAnswers)
}
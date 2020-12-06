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
	Groupforms := make(map[rune]int)		//Simulation of a Set to save the answers of a group
	var countAnswers int
	var lenGroup rune						//I save the len of the group in the set but I need a rune to do this

	for {
		end := !sc.Scan()
		
		if sc.Text() == "" || end {			//I have to do this to process the last input
			for question, numberOfYes := range Groupforms{
				if numberOfYes == Groupforms[lenGroup] && question != lenGroup{
					countAnswers++			
				}
			}
			Groupforms = make(map[rune]int)	//Clear the set for the next group
		}else{
			Groupforms[lenGroup]++			//I count the len of each group
			for _, answer := range sc.Text(){
				Groupforms[answer]++
			}
		}
		
		if end {
			break
		}
	}
	fmt.Println(countAnswers)
}
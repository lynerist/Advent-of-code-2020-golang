package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){

//	--- STANDARD FILE OPENING ---
	input, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory")
		return
	}
	defer input.Close()

//	--- START ---
	sc := bufio.NewScanner(input)
	IDs := make(map[int]bool)

	for sc.Scan(){
		currentID := findPlace(sc.Text()[:7], 0, 127) * 8 + findPlace(sc.Text()[7:], 0, 7)
		IDs[currentID] = true
	}

	for lr:=0; lr<8; lr++{
		for fb:=0; fb<127; fb++{
			if !IDs[fb*8+lr] && IDs[(fb+1)*8+lr] && IDs[(fb-1)*8+lr]{
				fmt.Println(fb*8+lr)
				return
			}
		}
	}
}

func findPlace(boardingPass string, bottom, top int)(int){ //recursive function that cuts the boarding pass from the left
	if boardingPass == "F" || boardingPass == "L"{
		return bottom
	}
	if boardingPass == "B" || boardingPass == "R"{
		return top
	}
	if boardingPass[0] == 'F' || boardingPass[0] == 'L'{
		return findPlace(boardingPass[1:], bottom, (top-bottom)/2+bottom)
	}
	return findPlace(boardingPass[1:], (top-bottom+1)/2+bottom, top)
}
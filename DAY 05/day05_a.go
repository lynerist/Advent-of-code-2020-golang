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
	var highestID int

	for sc.Scan(){
		currentID := findPlace(sc.Text()[:7], 0, 127) * 8 + findPlace(sc.Text()[7:], 0, 7)
		if currentID > highestID{
			highestID = currentID
		}
	}

	fmt.Println(highestID)
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
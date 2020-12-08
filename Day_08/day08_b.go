package main

import(
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type instruction struct{
	name	string
	value	int
	visited	bool
}

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
	var loop []instruction
	
	// --- INPUT READING ---
	for sc.Scan(){
		line := strings.Fields(sc.Text())
		number, _ := strconv.Atoi(line[1])
		loop = append(loop, instruction{line[0], number, false})
	}

	for ProgramCounter := range loop{		//I try to change every nop and jmp, one for each iteration
		if loop[ProgramCounter].name == "acc"{
			continue
		}

		copyLoop := make([]instruction, len(loop))	
		copy(copyLoop, loop)				//I need a copy for each simulation because slices are passed by reference
		copyLoop[ProgramCounter].name = change(copyLoop[ProgramCounter].name)

		done, accumulator := boot(copyLoop)	//this is a try of boot
		
		if done {
			fmt.Println(accumulator)
			return
		}
	}
}

func boot(loop []instruction )(bool, int){
	var ProgramCounter, accumulator int
	for{
		if ProgramCounter == len(loop){
			return true, accumulator
		}

		if loop[ProgramCounter].visited {
			return false, accumulator
		}
		loop[ProgramCounter].visited = true

		switch loop[ProgramCounter].name{
		case "jmp":
			ProgramCounter += loop[ProgramCounter].value - 1
		case "acc":
			accumulator += loop[ProgramCounter].value
		}
		ProgramCounter++
	}
}

func change(value string)string{
	if value == "nop"{
		return "jmp"
	}
	return "nop"
}
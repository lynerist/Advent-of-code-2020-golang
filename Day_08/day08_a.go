package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name    string
	value   int
	visited bool
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
	var (
		loop        	[]instruction
		accumulator 	int
		ProgramCounter	int
	)

	// --- INPUT READING ---
	for sc.Scan() {
		line := strings.Fields(sc.Text())
		number, _ := strconv.Atoi(line[1])
		loop = append(loop, instruction{line[0], number, false})
	}

	// --- BOOT EXECUTION ---
	for {
		if loop[ProgramCounter].visited {
			fmt.Println(accumulator)
			return
		}
		loop[ProgramCounter].visited = true

		switch loop[ProgramCounter].name {
		case "jmp":
			ProgramCounter += loop[ProgramCounter].value - 1	// -1 to compense the last instruction
		case "acc":
			accumulator += loop[ProgramCounter].value
		}
		ProgramCounter++
	}
}

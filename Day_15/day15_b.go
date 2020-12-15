package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

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
	sc.Scan()
	memory := make(map[int]int)

	for turn, number := range strings.Split(sc.Text(), ","){
		num, _ := strconv.Atoi(number)
		memory[num] = turn
	}
	
	var spoken int

	for turn := len(memory)+1; turn<30000000; turn++{
		lastSpoken, old := memory[spoken]
		if old{
			memory[spoken] = turn -1
			spoken = turn -1 - lastSpoken
		}else{
			memory[spoken] = turn -1
			spoken = 0
		}
	}
	fmt.Println(spoken)
	fmt.Println(len(memory))
}
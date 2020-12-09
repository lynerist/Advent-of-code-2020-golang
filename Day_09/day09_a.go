package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func mainn() {
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
		numbers		[]int
		index		int
	)

	for sc.Scan(){
		number, _ := strconv.Atoi(sc.Text())
		numbers = append(numbers, number)
		if len(numbers) <= 25{
			continue							//preamble
		}
		
		if !inPrevSum(numbers[index:index+25], number){
			fmt.Println(number)
			return
		}
		index++
	}
}

func inPrevSumm(prev []int, number int)bool{
	for i, n1 := range prev{
		for _, n2 := range prev[i+1:]{
			if n1 + n2 == number{
				return true
			}
		} 
	}
	return false
}

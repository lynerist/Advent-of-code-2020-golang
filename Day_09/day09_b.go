package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
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
			fmt.Println(adiacentSumMaxMin(numbers, number))
			return
		}
		index++
	}
}

func inPrevSum(prev []int, number int)bool{
	for i, n1 := range prev{
		for _, n2 := range prev[i+1:]{
			if n1 + n2 == number{
				return true
			}
		} 
	}
	return false
}

func adiacentSumMaxMin(list []int, number int)int{
	for numberAdiacent := 2; numberAdiacent<= len(list); numberAdiacent++{
		for firstAdiacent := 0; firstAdiacent<len(list)-numberAdiacent+1; firstAdiacent++{
			min, max, sum := list[firstAdiacent], 0, 0
			for _, n := range list[firstAdiacent:firstAdiacent+numberAdiacent]{
				sum += n
				if n < min{
					min = n
				}
				if n > max{
					max = n
				}
			}
			if sum == number{
				return min + max
			}
		}
	}
	return 0
}
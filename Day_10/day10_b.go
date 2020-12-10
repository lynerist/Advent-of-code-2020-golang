package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
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
	adapters := []int{0}
		
	for sc.Scan(){
		adapter, _ := strconv.Atoi(sc.Text())
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)

	var (
		previousAdapter		int
		step				int
	)
	combinations := 1
	cache := map[int]int{0:0, 1:1, 2:1}

	for i, adapter := range adapters{
		if i == len(adapters)-1{
			step--
		}
		if adapter-previousAdapter == 3 || i == len(adapters)-1{
			combinations *= conta(i-step, cache)
			step = i
		}
		previousAdapter = adapter
	}
	fmt.Println(combinations)	
}

func conta(num int, combinations map[int]int)int{
	n, ok := combinations[num]
	if ok{
		return n
	} 
	combinations[num] = conta(num-1, combinations) + conta(num-2, combinations) + conta(num-3, combinations)
	return combinations[num]
}
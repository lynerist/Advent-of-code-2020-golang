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
	var adapters []int
		
	for sc.Scan(){
		adapter, _ := strconv.Atoi(sc.Text())
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)

	var (
		previousAdapter	int
		differences		[3]int
	)

	for _, adapter := range adapters{
		differences[adapter-previousAdapter-1]++
		previousAdapter = adapter
	}
	differences[2]++
	fmt.Println(differences[0]*differences[2])
}
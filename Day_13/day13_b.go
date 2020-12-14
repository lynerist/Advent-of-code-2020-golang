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
	var ( 
		buses 		[]int
		timestamp 	int = 100000000000000
		step		int = 1
	)
	sc.Scan()
	sc.Scan()
	for _, bus := range strings.Split(sc.Text(), ","){
		timestamp, _ := strconv.Atoi(bus)
		buses = append(buses, timestamp)
	}

	for i, bus := range buses{
		if bus == 0 {
			continue
		}
		for (timestamp+i)%bus != 0{
			timestamp += step
		}
		step *= bus
	}
	fmt.Println(timestamp)
}

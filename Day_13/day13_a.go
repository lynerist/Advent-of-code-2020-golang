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
		departTime int
		buses []int
	)

	sc.Scan()
	departTime, _ = strconv.Atoi(sc.Text())

	sc.Scan()
	for _, bus := range strings.Split(sc.Text(), ","){
		timestamp, ok := strconv.Atoi(bus)
		if ok == nil {
			buses = append(buses, timestamp)
		}
	}

	var (
		early	int
		wait	int
	)

	for i, bus := range buses {
		if bus - departTime % bus < wait || i == 0{
			wait = bus - departTime % bus
			early = bus
		}
	}

	fmt.Println(early * wait)
}
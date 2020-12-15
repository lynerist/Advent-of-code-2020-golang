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
	sc 		:= bufio.NewScanner(input)
	mem 	:= make(map[string]int64)
	mask 	:= ""

	for sc.Scan(){
		line := strings.Fields(sc.Text())
		if line[0] == "mask"{
			mask = line[2]
		}else{
			number, _ := strconv.Atoi(line[2])
			value := fmt.Sprintf("%b", number)
			for len(value) != 36 {
				value = "0" + value
			}
			maskedValue :=""
			for i, bit := range mask {
				if bit != 'X' {
					maskedValue += string(bit)
				}else{
					maskedValue += string(value[i])
				}
			}
			mem[line[0][4:len(line[0])-1]], _ = strconv.ParseInt(maskedValue, 2, 64)
		}
	}
	var sum int
	for _, value := range mem{
		sum += int(value)
	}
	fmt.Println(sum)
}
package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
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
	mem 	:= make(map[string]int)
	mask 	:= ""

	for sc.Scan(){
		line := strings.Fields(sc.Text())
		if line[0] == "mask"{
			mask = line[2]
		}else{
			number, _ := strconv.Atoi(line[0][4:len(line[0])-1])
			address := fmt.Sprintf("%b", number)
			for len(address) != 36 {
				address = "0" + address
			}
			maskedAddress :=""
			countX := 0
			for i, bit := range mask {
				switch bit {
					case '1':
						maskedAddress += "1"
					case '0':
						maskedAddress += string(address[i])
					case 'X':
						maskedAddress += "X"
						countX++
				}
			}

			for i:=0; i<int(math.Pow(2, float64(countX))); i++{
				combination := fmt.Sprintf("%b", i)
				for len(combination) < countX{
					combination = "0" + combination
				}
				tempMasked := maskedAddress
				for _, bit := range combination{
					tempMasked = strings.Replace(tempMasked, "X", string(bit),1)
				}
				mem[tempMasked], _ = strconv.Atoi(line[2])
			}
		}
	}
	var sum int
	for _, address := range mem{
		sum += address
	}
	fmt.Println(sum)
}
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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

	sc := bufio.NewScanner(input)
	rules := make(map[string]func(int)bool)
	var inputPart, errorRate int


	for sc.Scan(){
		if sc.Text() == ""{
			inputPart++
		}
		if sc.Text() == "" || sc.Text() == "your ticket:" || sc.Text() == "nearby tickets:" {
			continue
		}

		switch inputPart {
		case 0:
			line := strings.Split(sc.Text(), ":")
			var a,b,c,d int
			fmt.Sscanf(line[1]," %d-%d or %d-%d",&a,&b,&c,&d)
			rules[line[0]] = func(num int)bool{
				return num >= a && num <= b || num >= c && num <= d
			}
		case 2:
			for _, num := range strings.Split(sc.Text(), ","){
				number, _ := strconv.Atoi(num)
				for _, rule := range rules{
					if rule(number){
						errorRate -= number
						break
					}
				}
				errorRate += number
			}
		}
	}
	fmt.Println(errorRate)
}
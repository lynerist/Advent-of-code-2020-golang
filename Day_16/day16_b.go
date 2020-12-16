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
	var fields []string
	var inputPart int
	var myticket []int
	possibleRules := make(map[string][]bool)
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
			fields = append(fields, line[0])
			possibleRules[line[0]] = []bool{}
			
		case 1:
			for _, num := range strings.Split(sc.Text(), ","){
				number, _ := strconv.Atoi(num)
				myticket = append(myticket, number)
			}
			for name := range possibleRules{
				possibleRules[name] = make([]bool, len(myticket))
				for i := range possibleRules[name]{possibleRules[name][i] = true}
			}
		case 2:
			var notPossibleRulesline [][]string
			var valid bool
			for _, num := range strings.Split(sc.Text(), ","){
				valid = false			
				var notPossibleRules []string
				number, _ := strconv.Atoi(num)
				for name, rule := range rules{
					if rule(number){
						valid = true
					}else{
						notPossibleRules = append(notPossibleRules, name)
					}
				}	
				notPossibleRulesline = append(notPossibleRulesline, notPossibleRules)
				if !valid{
					break
				}
			}
			if valid{
				for i, notPossibleRules := range notPossibleRulesline{
					for _, name := range notPossibleRules{
						possibleRules[name][i] = false
					}
				}
			}
		}		
	}
	positionRules := make([]string, len(rules))

	for i:=0; i<len(positionRules); i++ {
		for name, possibilities := range possibleRules{
			var position []int
			for i, possible := range possibilities{
				if possible && positionRules[i] == ""{
					position = append(position, i)
				}
			}
			if len(position) == 1 {
				positionRules[position[0]] = name
			}
		}
	}

	product := 1

	for i, name := range positionRules {
		if strings.Fields(name)[0] == "departure"{
			product *= myticket[i]
		}
	}
	
	fmt.Println(product)
}
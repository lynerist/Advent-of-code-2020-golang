package main

import(
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type bagLabel struct{
	bag		string
	number	int
}

func main(){
//	--- STANDARD FILE OPENING ---
	input, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

//	--- START ---
	sc := bufio.NewScanner(input)
	rules := make(map[string][]bagLabel)	//each bag is mapped to the list of bags that contains with the number of istances of each bag

	for sc.Scan(){
		parsedInput := strings.ReplaceAll(sc.Text(), " bags", "")
		parsedInput = strings.ReplaceAll(parsedInput, " bag", "")
		parsedInput = strings.ReplaceAll(parsedInput, ".", "")
		rule := strings.Split(parsedInput, " contain ")

		for _, bag := range strings.Split(rule[1], ", "){
			numberOfBags, _ := strconv.Atoi(bag[0:1])
			rules[rule[0]] = append(rules[rule[0]], bagLabel{bag[2:], numberOfBags})
		}
	}
	fmt.Println(countIn("shiny gold", &rules) - 1)	// -1 because I don't have to count the Shiny gold bag
}

func countIn(lookedbag string, rules *map[string][]bagLabel)int{	
	countBags := 1
	for _, bag := range (*rules)[lookedbag]{		//I countIn each bag contained in the current bag recursively
		countBags += bag.number * countIn(bag.bag, rules)
	}
	return countBags
}
package main

import(
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main(){
//	--- STANDARD FILE OPENING ---
	input, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer input.Close()

// --- START --
	sc := bufio.NewScanner(input)
	rules := make(map[string][]string)		//each bag is mapped to the list of bags that contains
	valids := make(map[string]bool)			//each bag that contains a shiny gold bag is a valid bag
	valids["shiny gold"] = true

	for sc.Scan(){
		parsedInput := strings.ReplaceAll(sc.Text(), " bags", "")
		parsedInput = strings.ReplaceAll(parsedInput, " bag", "")
		parsedInput = strings.ReplaceAll(parsedInput, ".", "")
		rule := strings.Split(parsedInput, " contain ")

		for _, bag := range strings.Split(rule[1], ", "){
			rules[rule[0]] = append(rules[rule[0]], bag[2:])
		}
	}

	for bag := range rules{
		if lookIn(bag, &rules, &valids){	//lookIn says if the bag contains a shiny gold bag
			valids[bag] = true
		}
	}
	fmt.Println(len(valids) - 1)			// -1 because I don't have to count the Shiny gold bag
}

func lookIn(lookedBag string, rules *map[string][]string, valids *map[string]bool)bool{
	if (*valids)[lookedBag]{
		return true
	}
	for _, bag := range (*rules)[lookedBag]{	//I lookIn each bag contained in the current bag recursively
		if lookIn(bag, rules, valids){
			(*valids)[bag] = true
			return true
		}
	}
	return false
}
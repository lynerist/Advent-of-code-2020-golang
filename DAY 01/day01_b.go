package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func main(){

//	--- STANDARD FILE OPENING ---
	file, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory")
		return
	}
	defer file.Close()

// 	--- LET'S READ THE INPUT AND SAVE ALL IN A SLICE OF INT ---

	sc := bufio.NewScanner(file)
	var values []int
	for(sc.Scan()){
		num, _ := strconv.Atoi(sc.Text())
		values = append(values, num)
	}

//	--- NOW I SORT THE INPUT ---
	sort.Ints(values)

// --- I COMPARE EVERY VALUE WITH EACH OTHER
	for i, first := range (values){
		for j, second := range (values[i+1:]){
			for _, third := range (values[j+1:]){
				if first + second + third == 2020{
					fmt.Println(first * second * third)
					return
				}else if first + second > 2020{ 
					break	//if we are over 2020 there is no reason to compare the other values with the current second
				}
			}
		}
	}
}
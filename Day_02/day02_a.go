package main

import(
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main(){

//	--- STANDARD FILE OPENING ---
	file, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err, "\nYou need a file called input.txt in this directory")
		return
	}
	defer file.Close()

	var validPasswords int

	sc := bufio.NewScanner(file)
	for(sc.Scan()){
		var (
			lowerLimit, upperLimit int
			char rune
			password string
		)
		fmt.Sscanf(sc.Text(), "%d-%d %c: %s", &lowerLimit, &upperLimit, &char, &password) //Formatted scan of the input string
		
		if strings.Count(password, string(char)) >= lowerLimit && strings.Count(password, string(char)) <= upperLimit{
			validPasswords++ 
		}
	}
	fmt.Println(validPasswords)	
}
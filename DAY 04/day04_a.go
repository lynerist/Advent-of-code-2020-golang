package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func mainn(){

//	--- STANDARD FILE OPENING ---
	input, err := os.Open("input.txt") 
	if err != nil{
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory")
		return
	}
	defer input.Close()

	sc := bufio.NewScanner(input)
	var (
		passaport string
		valids int
	)

// 	--- START ---	

	for {
		end := !sc.Scan()

		if sc.Text() == "" || end {
			passaportFields := make(map[string]string)		// map nameField -> valueField
			for _, field := range strings.Fields(passaport){
				passaportFields[field[:3]] = field[4:]
			}
			
			_, hasCid := passaportFields["cid"]				//optional value
			if len(passaportFields) == 8 || len(passaportFields) == 7 && !hasCid{
				valids++
			}
			passaport = ""									//clear passaport
		}else{
			passaport += sc.Text()+" "						//append passaport information
		}
		
		if end{
			break											//it was the last line
		}
	}
	fmt.Println(valids)
}
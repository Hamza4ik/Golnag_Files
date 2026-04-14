package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for {
		fmt.Println("Choice the menu")
		fmt.Println("1.Create File")
		fmt.Println("0. Out")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			CreateFile()
		case 0:
			return
		default:
			fmt.Println("Choice from 1 to 4")
			return
		}
	}
}

func CreateFile() {
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal("Can't create file: ", err)
		return
	}

	defer file.Close()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("Choice the menu")
		fmt.Println("1.Create File")
		fmt.Println("2.Write to file")
		fmt.Println("3.Read from file")
		fmt.Println("0. Out")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			CreateFile()
		case 2:
			WriteFile()
		case 3:
			ReadFromFile()
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

func WriteFile() {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter text (type'exit' to stop):")

		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)

		if text == "exit" {
			break
		}

		if text == "" {
			continue
		}

		file.WriteString("User: " + text + "\n")
	}

}

func ReadFromFile() {
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal("Can't Read file", err)
	}
	fmt.Println(string(file))
}

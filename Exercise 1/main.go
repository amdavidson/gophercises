package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv delimited set of questions and solutions")
	flag.Parse()

	fmt.Print("Opening question file\n")
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Errorf("Could not open filename: %s \n %s", *csvFilename, err)
	}
	defer file.Close()

	fmt.Print("Parsing questions...\n")
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Errorf("Could not parse %s as CSV", *csvFilename)
	}

	rand.Seed(time.Now().Unix())

	var correct int
	var wrong int

	fmt.Printf("Enter the answer to each question.\nWhen you are finished type the word `done`\n")

	for {
		r := rand.Intn(len(lines))

		fmt.Print("What is the answer to \n", lines[r][0], "\n")

		var input string
		fmt.Scanln(&input)

		if input == "done" {
			break
		}

		if input == lines[r][1] {
			correct++
			fmt.Print("Correct!\n")

		} else {
			wrong++
			fmt.Print("Boo incorrect!\n")
		}
	}

	fmt.Print("You got ", correct, " correct.\n")
	fmt.Print("You got ", wrong, " wrong.\n")
	fmt.Print("That is a ", 100*float64(correct)/(float64(correct)+float64(wrong)), "% hit rate.\n")

	os.Exit(0)
}

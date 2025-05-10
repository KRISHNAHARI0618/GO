package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'questions and answers' ")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed To Open CSV File %s \n", *csvFileName))
	}

	r := csv.NewReader(file) //reading the csv file 
	lines, err := r.ReadAll() // taking each lines
	if err != nil { //errors lekunteyy kndha exit run avutundhi 
		exit("Failed to parse the provided csv file")
	}
	fmt.Println(lines)

	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You have answered %v out of  %v", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

/*
1: Parse the command line flags using the flag commands
2: flag.String will take three arguments one with -csv or --csv
3: second is file name if we are passing it will take the default file what ever we pass
4: flag.parse will actually look at the command line arguments provided when the program was run
5:
*/

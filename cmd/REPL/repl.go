package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// String that get printed on every new line

const cliDBINTRO = "\n    ######################\n    #  Welcome to alSQL  #\n    ######################\n"
const cliDBPREFIX = "> "

// Possible keywords inputed into CLI

type metaCommandStatus int

const (
	metaSuccess metaCommandStatus = iota
	metaUnrecognized
	metaFailure
)

const (
	exitMetaCommand string = ".exit"
)

type statementType int

const (
	statementInsert statementType = iota
	statementDelete
)

func printIntro() {
	fmt.Println(cliDBINTRO)
}

func cleanInput(input *string) {
	*input = strings.TrimSpace(*input)
}

func isMetaCommand(input string) bool {
	return input[:1] == "."
}

func doMetaCommand(input string) error {
	if input == exitMetaCommand {
		os.Exit(1)
	}
	return errors.New("Unrecognized command")
}

// type statement struct {
// 	type statementType
// }

func isStatementCommand(input string) bool {
	return input[:1] == "@"
}

// If statement starts with a ".", this is a METASTATEMENT
// If a value starts with @ that means its a keyword like "SELECT", "JOIN", etc.

func main() {
	printIntro()
	for {
		var input string

		// check if its a METASTATEMENT
		fmt.Print(cliDBPREFIX)
		fmt.Scanln(&input)
		cleanInput(&input)

		if isMetaCommand(input) {
			err := doMetaCommand(input)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		} else if isStatementCommand(input) {
			fmt.Printf("Can't do that yet\n")
		}

	}
}

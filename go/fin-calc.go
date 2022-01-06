package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
	"strconv"
	"strings"
)

func main() {
	monthly := ask("Monthly addition ($/month)", "Monthly addition should be a number")
	years := ask("Number of years", "Number of years should be a positive number")
	cap := ask("Capitalization (%)", "Capitalization should be a number")
	initial := ask("Initial amount ($)", "Initial amount should be a number")

	total := initial
	totalCap := initial

	for i := 1; i <= years*12; i++ {
		total += monthly
		totalCap += monthly + totalCap*cap/1200
	}

	printer := message.NewPrinter(language.English)
	printer.Printf("\033[32mTotal:\033[0m $%d\n", total)
	printer.Printf("\033[32mCapitalization:\033[0m $%d\n", totalCap-total)
	printer.Printf("\033[32mTotal with capitalization:\033[0m $%d\n", totalCap)
}

func ask(question string, inputError string) int {
	fmt.Printf("\033[34m%s:\033[0m ", question)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	answers := strings.Fields(scanner.Text())

	var answer int
	var err error

	if len(answers) == 1 {
		answer, err = strconv.Atoi(answers[0])
	} else {
		err = errors.New("More than one thing answers")
	}

	if err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", inputError)

		return ask(question, inputError)
	}

	return answer
}

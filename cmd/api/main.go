package main

import (
	"bufio"
	"calculator/pkg/calc"
	"calculator/pkg/calc/command"
	"calculator/pkg/calc/expression"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout

	storage := expression.NewStorage()
	parser := command.NewParser(storage, out)

	calculator := calc.NewCalculator(storage, parser)

	for {
		var str string
		str, _ = in.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		calculator.RunCommand(str)
	}
}

package main

import (
	"bufio"
	"os"
	"strings"

	"calculator/pkg/calc"
	"calculator/pkg/calc/cache"
	"calculator/pkg/calc/command"
	"calculator/pkg/calc/expression"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout

	cacheStorage := cache.NewCacheStorage()

	storage := expression.NewStorage()
	cStorage := cache.NewCacheForStorage(storage, cacheStorage)

	cmdBuilder := command.NewCmdBuilder(cStorage)
	cacheCmdBuilder := cache.NewCacheCmdBuilder(cmdBuilder, cacheStorage)

	parser := command.NewParser(cStorage, out, cacheCmdBuilder)

	calculator := calc.NewCalculator(parser, out)

	for {
		var str string
		str, _ = in.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		calculator.RunCommand(str)
	}
}

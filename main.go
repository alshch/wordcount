package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print(WordCount(input))
}

func WordCount(value string) int {
	results := regexp.MustCompile(`[\S]+`).FindAllString(value, -1)
	return len(results)
}

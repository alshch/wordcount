// This is a template for the wordcount CLI tool. It should work as follows:
// $ ./wordcount 'go is awesome'
// 3
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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("convertToBin result:")
	fmt.Println(convertToBin(5), convertToBin(13), convertToBin(548612), convertToBin(0))

	fmt.Println("abc.txt contents:")
	printFile("../branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d
aa
bb

123
`
	printFileContents(strings.NewReader(s))
}

package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
			fmt.Println(lastI)
			fmt.Println(start)
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
		fmt.Println(lastOccurred)
	}

	return maxLength
}

func main()  {
	fmt.Println(lengthOfNonRepeatingSubStr("abcbacbb"))
}
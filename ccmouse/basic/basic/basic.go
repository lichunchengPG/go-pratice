// 注释注释
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "abc"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func trianger() {
	var a, b int = 3, 4
	fmt.Println(calTriangle(a, b))
}

func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	trianger()
	consts()
	enums()
}

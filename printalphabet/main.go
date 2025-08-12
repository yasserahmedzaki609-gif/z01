package main

import "github.com/01-edu/z01"

func main() {
	// Print lowercase alphabet from 'a' to 'z'
	for c := 'a'; c <= 'z'; c++ {
		z01.PrintRune(c)
	}
	// Print newline at the end
	z01.PrintRune('\n')
}

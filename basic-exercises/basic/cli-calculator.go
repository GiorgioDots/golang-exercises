package basic

import (
	"flag"
	"fmt"
	"os"
)

func CliCalc() {
	// Define flags
	op := flag.String("op", "", "Operation to perform: add, sub, mul, div")
	a := flag.Float64("a", 0, "First operand")
	b := flag.Float64("b", 0, "Second operand")

	// Parse the flags
	flag.Parse()

	switch *op {
	case "add":
		fmt.Println("Add:", *a+*b)
	case "sub":
		fmt.Println("Sub:", *a-*b)
	case "div":
		if *b == 0 {
			fmt.Println("Error: cannot divide by zero")
			os.Exit(1)
		}
		fmt.Println("Div:", *a / *b)
	case "mul":
		fmt.Println("Mul:", *a**b)
	default:
		fmt.Println("Unsupported operation. Use --op add|sub|mul|div")
		os.Exit(1)
	}
}

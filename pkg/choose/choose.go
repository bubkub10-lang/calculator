package choose

import (
	"calculator/pkg/calc"
	"fmt"
	"os"
)

func Numbers(first, second *float64) (*float64, *float64) {
	fmt.Printf("> Calculator \n> (need 2 numbers) \n> 1:")
	fmt.Scan(first)
	fmt.Printf("> 2:")
	fmt.Scan(second)
	return first, second
}

func Choose() {
	var first, second float64
	Numbers(&first, &second)
	for {
		var choose int
		fmt.Printf("> Choose Action \n  > [0] Exit \n  > [1] Addition \n  > [2] Substraction \n  > [3] Multiplication \n  > [4] Division \n  > [5] Pow \n  > [6] New Numbers")
		fmt.Printf("> ")
		fmt.Scan(&choose)
		switch choose {
		case 0:
			os.Exit(0)
		case 1:
			fmt.Println(calc.Add(first, second))
		case 2:
			fmt.Println(calc.Sub(first, second))
		case 3:
			fmt.Println(calc.Mult(first, second))
		case 4:
			fmt.Println(calc.Dev(first, second))
		case 5:
			fmt.Println(calc.Pow(first, second))
		case 6:
			first = 0
			second = 0
			Numbers(&first, &second)
		default:
			fmt.Println("error: doesn't have this option")
			continue
		}
	}
}

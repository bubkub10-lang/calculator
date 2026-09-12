package choose

import (
	"calculator/pkg/calc"
	"fmt"
	"os"
)

func mathNumbers(first, second *float64) (*float64, *float64) {
	fmt.Printf("> Math Calculator \n> (need 2 numbers) \n> 1:")
	fmt.Scan(first)
	fmt.Printf("> 2:")
	fmt.Scan(second)
	return first, second
}

func trigoNum(num *float64) *float64 {
	fmt.Printf("> Trigonometric Calculator \n>(need 1 number) \n> number:")
	fmt.Scan(num)
	return num
}

func logNum(num *float64) *float64 {
	fmt.Printf("> Logarithmetic and Root Calculator \n>(need 1 number) \n> number:")
	fmt.Scan(num)
	return num
}

func DefChoose() {
	var first, second float64
	mathNumbers(&first, &second)
	for {
		var choose int
		fmt.Printf("> Choose Action \n  > [0] Exit \n  > [1] Addition \n  > [2] Substraction \n  > [3] Multiplication \n  > [4] Division \n  > [5] Pow \n > [6] New Numbers \n > [7] Next Page \n")
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
			mathNumbers(&first, &second)
		case 7:
			trigoChoose()
		default:
			fmt.Println("error: doesn't have this option")
			continue
		}
	}
}

func trigoChoose() {
	var num float64
	trigoNum(&num)
	for {
		var choose int
		fmt.Printf("> Choose Action \n > [0] Exit \n > [1] Sinus \n > [2] Cosinus \n > [3] Tangens \n > [4] Asinus \n > [5] Acosinus \n > [6] Atangens \n > [7] Degrees to Radians \n > [8] Radians to Degrees \n > [9] New Number \n > [10] Next Page \n")
		fmt.Printf("> ")
		fmt.Scan(&choose)
		switch choose {
		case 0:
			os.Exit(0)
		case 1:
			fmt.Println(calc.Sin(num))
		case 2:
			fmt.Println(calc.Cos(num))
		case 3:
			fmt.Println(calc.Tan(num))
		case 4:
			fmt.Println(calc.Asin(num))
		case 5:
			fmt.Println(calc.Acos(num))
		case 6:
			fmt.Println(calc.Atan(num))
		case 7:
			fmt.Println(calc.DegtoRad(num))
		case 8:
			fmt.Println(calc.RadtoDeg(num))
		case 9:
			num = 0
			trigoNum(&num)
		case 10:
			logAndRootsChoose()
		}
	}
}

func logAndRootsChoose() {
	var num float64
	logNum(&num)
	for {
		var choose int
		fmt.Printf("> Choose Action \n > [0] Exit \n > [1] Exponentiation \n > [2] Logarithm \n > [3] Logarithm from 2 \n > [4] Logarithm from 10 \n > [5] Square Root \n > [6] Cube Root \n > [7] New Number \n > [8] Page №1 \n")
		fmt.Printf("> ")
		fmt.Scan(&choose)
		switch choose {
		case 0:
			os.Exit(0)
		case 1:
			fmt.Println(calc.Exp(num))
		case 2:
			fmt.Println(calc.Log(num))
		case 3:
			fmt.Println(calc.Log2(num))
		case 4:
			fmt.Println(calc.Log10(num))
		case 5:
			fmt.Println(calc.Sqrt(num))
		case 6:
			fmt.Println(calc.Cbrt(num))
		case 7:
			num = 0
			logNum(&num)
		case 8:
			DefChoose()
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcRome(a, operator, b string) (result string) {
	a1, b1 := convertRomeToint(a), convertRomeToint(b)
	res := calculation(operator, a1, b1)

	if res <= 0 {
		fmt.Println("err: not negative numbers in Roman numerals")
		os.Exit(0)
	}

	for k, v := range numbInt {
		if res == k {
			result = v
			break
		}
	}

	fmt.Println(result)
	return result
}

func convertRomeToint(str string) (num int) {
	for k, v := range numbRome {
		if str == k {
			num = v
		}
	}
	return num
}

func calcInt(a, operator, b string) {
	a1, err := strconv.Atoi(a)
	b1, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
	}

	if (a1 <= 0 || a1 > 10) || (b1 <= 0 || b1 > 10) {
		fmt.Println("number is out of range")
		os.Exit(0)
	}
	result := calculation(operator, a1, b1)
	fmt.Println(result)
}

func calculation(operator string, a, b int) (result int) {
	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("unknown operator")
		os.Exit(0)
	}
	return result
}

func isDigit(str string) bool {
	var label bool
	atoi, _ := strconv.Atoi(str)
	if atoi >= 1 {
		label = true
	} else {
		label = false
	}
	return label
}

func isRome(str string) bool {
	var label bool
	for k := range numbRome {
		if k == str {
			label = true
			break
		}
	}
	return label
}

func selectCalc(a, operator, b string) {
	if isDigit(a) && isDigit(b) {
		calcInt(a, operator, b)
	} else if isRome(a) && isRome(b) {
		calcRome(a, operator, b)
	} else {
		fmt.Println("err mismatch of the number system")
		os.Exit(0)
	}
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

var numbRome = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4,
	"V": 5, "VI": 6, "VII": 7, "VIII": 8,
	"IX": 9, "X": 10,
}

var numbInt = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
	21: "XXI", 24: "XXIV", 25: "XXV", 27: "XXVII", 28: "XXVIII", 30: "XXX", 32: "XXXII", 35: "XXXV", 36: "XXXVI", 40: "XL",
	42: "XLII", 45: "XLV", 48: "XLVIII", 49: "XLIX", 50: "L", 56: "LVI", 60: "LX", 63: "LXIII", 64: "LXIV", 70: "LXX",
	72: "LXXII", 80: "LXXX", 81: "LXXXI", 90: "XC", 100: "C",
}

func main() {
	fmt.Print("Enter the expression: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	exp := strings.Fields(str)

	if len(exp) < 3 {
		fmt.Println("string is not a math operation")
		return
	}

	firstNum := exp[0]
	operator := exp[1]
	secondNum := exp[2]

	selectCalc(firstNum, operator, secondNum)

}

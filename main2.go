package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]string{
	"I":    "1",
	"II":   "2",
	"III":  "3",
	"IV":   "4",
	"V":    "5",
	"VI":   "6",
	"VII":  "7",
	"VIII": "8",
	"IX":   "9",
	"X":    "10",
}

type Roam struct {
	num1 string
	num2 string
	zn   string
}

func main() {
	sliceMath := TextInput()
	structMath := Roam{sliceMath[0], sliceMath[2], sliceMath[1]}

	structMath.Cek()
}

func TextInput() []string {
	textInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textInput = strings.TrimSpace(textInput)

	// Разделение строки на слайс строк
	sliceMath := strings.Split(textInput, " ")

	if len(sliceMath) != 3 {
		panic("вывод ошибки, так как формат математической операции не удовлетворяет заданию - 2 числа и знак операции")

	}
	return sliceMath
}

func (r *Roam) Cek() {
	n1, ok1 := roman[r.num1]
	n2, ok2 := roman[r.num2]

	if ok1 && ok2 {
		fmt.Println(r.GenerationRim(r.Calculator(n1, n2)))

	} else if strings.ContainsAny(r.num1, "0123456789") && strings.ContainsAny(r.num2, "0123456789") {
		fmt.Println(r.Calculator(r.num1, r.num2))

	} else {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
}

func (r *Roam) Calculator(n1, n2 string) int {
	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)

	switch {
	case r.zn == "+":
		return num1 + num2
	case r.zn == "-":
		return num1 - num2
	case r.zn == "*":
		return num1 * num2
	case r.zn == "/":
		return num1 / num2
	default:
		panic("вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")

	}
}

func (r *Roam) GenerationRim(num int) string {
	if num <= 0 {
		panic("вывод ошибки, так как в римской системе нет отрицательных чисел")
	}

	var result string
	var roman = [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var arab = [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(arab); i++ {
		for num >= arab[i] {
			num -= arab[i]
			result += roman[i]
		}
	}
	return result
}

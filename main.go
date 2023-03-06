package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Roman римские цифры от 1 до 10 в мапе
var Roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func main() {
	// вызов функции калькулятор
	if sliceMath, err := textInput(); err != nil {
		fmt.Println(err)
	} else {
		Calculator(sliceMath)
	}
}

// textInput Функция ввода пирвой строки и возврата слайса строк и ошибки
func textInput() ([]string, error) {
	// Запрос на ввод математической операции одгой строкой
	textInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return nil, errors.New("вывод ошибки, так как ввод не удался")
	}
	textInput = strings.TrimSpace(textInput)

	// Разделение строки на слайс строк
	sliceMath := strings.Split(textInput, " ")

	// проверка на количество символов и наличие знака математической операции
	if len(sliceMath) != 3 || !strings.ContainsAny(sliceMath[1], "+-*/") {
		return nil, errors.New("вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	return sliceMath, nil
}

// Calculator Функция с основной логикой калькулятора
func Calculator(sliceMath []string) {
	// проверка на вхождение цыфры от 1 до 10 и вызов функции для расчета
	if strings.ContainsAny(sliceMath[0], "0123456789") && strings.ContainsAny(sliceMath[2], "0123456789") {
		calculateNum(sliceMath)
		return

		// проверка на вхождение римских цыфр от I до X через функцию romanToRim
	} else if romanToRim(sliceMath[0]) && romanToRim(sliceMath[2]) {
		calculateRoman(sliceMath)
		return

		// провекра на разные системы счисления
	} else {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}

}

// calculateNum преобразование нескольких строк в целые числа спомощью функции strToInt
func calculateNum(str1 []string) {

	num1 := strToInt(str1[0])
	num2 := strToInt(str1[2])

	// проверка на вхождение знака + - * / в switch
	switch str1[1] {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	}

}

// strToInt преобразование строки в целое число
func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic("Ошибка преобразования строки в целое число")
	}
	return num
}

// romanToRim проверка на вхождение римских цыфр от I до X
func romanToRim(str string) bool {
	if _, ok := Roman[str]; ok {
		return true
	}
	return false
}

// calculateRoman математические операции и преобразования решения в римскую цифру
func calculateRoman(str1 []string) {
	switch {
	case str1[1] == "+":
		fmt.Println(generationRoman(Roman[str1[0]] + Roman[str1[2]]))

	case str1[1] == "-":
		if Roman[str1[0]]-Roman[str1[2]] <= 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		} else {
			fmt.Println(generationRoman(Roman[str1[0]] - Roman[str1[2]]))
		}

	case str1[1] == "*":
		fmt.Println(generationRoman(Roman[str1[0]] * Roman[str1[2]]))

	case str1[1] == "/":
		if Roman[str1[0]]/Roman[str1[2]] <= 0 {
			fmt.Println("Отрицательное число")
		} else {
			fmt.Println(generationRoman(Roman[str1[0]] / Roman[str1[2]]))
		}

	}
}

// функция преобразования арабских цифр в римские от 1 до 100
func generationRoman(num int) string {
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

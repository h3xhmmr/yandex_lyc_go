package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Solve(num1_str, num2_str, op string) string { // создаем функцию, которая будет считать нужные элементы строки
	num1, _ := strconv.Atoi(num1_str)
	num2, _ := strconv.Atoi(num2_str)
	num1_f := float64(num1)
	num2_f := float64(num2)
	switch op {
	case "/":
		return strconv.FormatFloat(num1_f/num2_f, 'g', -1, 64)
	case "*":
		return strconv.FormatFloat(num1_f*num2_f, 'g', -1, 64)
	case "+":
		return strconv.FormatFloat(num1_f+num2_f, 'g', -1, 64)
	case "-":
		return strconv.FormatFloat(num1_f-num2_f, 'g', -1, 64)
	}
	return ""
}

func Calc(expression string) (float64, error) {
	error_unknown := errors.New("Error: unknown sign or forgotten comma") // проверяем строку на наличие ненужных символов или незакрытых скобок
	error_divide := errors.New("Error: divide by 0")
	expression = strings.ReplaceAll(expression, " ", "")
	exp_nums := expression
	exp_signs := expression
	for _, i := range "+-/*()" {
		exp_nums = strings.ReplaceAll(exp_nums, string(i), " ") // создаем список всех чисел строки и строчку с операторами
	}
	exp_nums = strings.TrimRight(exp_nums, " ")
	exp_nums_slice := strings.Split(exp_nums, " ")
	for _, i := range "1234567890" {
		exp_signs = strings.ReplaceAll(exp_signs, string(i), "")
	}
	for j := 0; j < len(expression)-1; j++ {
		if strings.Contains("+-*/(", string(expression[j])) && strings.Contains("+-*/)", string(expression[j+1])) {
			return 0, error_unknown
		}
	}
	if strings.Count(expression, "(") != strings.Count(expression, ")") {
		return 0, error_unknown
	}
	if expression == "" || expression == " " {
		return 0, error_unknown
	}
	if strings.Contains("*(/+-", string(expression[len(expression)-1])) {
		return 0, error_unknown
	}
	if strings.Count(expression, "(") == 0 { // если в строке нет скобок, просто считаем каждый элемент строки(первое число + знак операции + второе число) и переписываем его на результат этой операции, делаем это пока не закончатся нужные символы
		for strings.Count(expression, "*") != 0 {
			var str_in_comma string
			num1_str := exp_nums_slice[strings.Index(exp_signs, "*")]
			num2_str := exp_nums_slice[strings.Index(exp_signs, "*")+1]
			str_in_comma = num1_str + "*" + num2_str
			expression = strings.ReplaceAll(expression, str_in_comma, Solve(num1_str, num2_str, "*"))
			exp_nums = expression
			exp_signs = expression
			for _, i := range "+-/*()" {
				exp_nums = strings.ReplaceAll(exp_nums, string(i), " ")
			}
			exp_nums_slice = strings.Split(exp_nums, " ")
			for _, i := range "1234567890" {
				exp_signs = strings.ReplaceAll(exp_signs, string(i), "")
			}
		}
		for strings.Count(expression, "/") != 0 {
			var str_in_comma string
			num1_str := exp_nums_slice[strings.Index(exp_signs, "/")]
			num2_str := exp_nums_slice[strings.Index(exp_signs, "/")+1]
			if num2_str == "0" { // если есть деление на ноль - выводим ошибку
				return 0, error_divide
			}
			str_in_comma = num1_str + "/" + num2_str
			expression = strings.ReplaceAll(expression, str_in_comma, Solve(num1_str, num2_str, "/"))
			exp_nums = expression
			exp_signs = expression
			for _, i := range "+-/*()" {
				exp_nums = strings.ReplaceAll(exp_nums, string(i), " ")
			}
			exp_nums_slice = strings.Split(exp_nums, " ")
			for _, i := range exp_nums_slice {
				exp_signs = strings.ReplaceAll(exp_signs, i, "")
			}
		}
		for strings.Count(expression, "+") != 0 {
			var str_in_comma string
			num1_str := exp_nums_slice[strings.Index(exp_signs, "+")]
			num2_str := exp_nums_slice[strings.Index(exp_signs, "+")+1]
			str_in_comma = num1_str + "+" + num2_str
			expression = strings.ReplaceAll(expression, str_in_comma, Solve(num1_str, num2_str, "+"))
			exp_nums = expression
			exp_signs = expression
			for _, i := range "+-/*()" {
				exp_nums = strings.ReplaceAll(exp_nums, string(i), " ")
			}
			exp_nums_slice = strings.Split(exp_nums, " ")
			for _, i := range exp_nums_slice {
				exp_signs = strings.ReplaceAll(exp_signs, i, "")
			}
		}
		for strings.Count(expression, "-") != 0 {
			var str_in_comma string
			num1_str := exp_nums_slice[strings.Index(exp_signs, "-")]
			num2_str := exp_nums_slice[strings.Index(exp_signs, "-")+1]
			str_in_comma = num1_str + "-" + num2_str
			expression = strings.ReplaceAll(expression, str_in_comma, Solve(num1_str, num2_str, "-"))
			exp_nums = expression
			exp_signs = expression
			for _, i := range "+-/*()" {
				exp_nums = strings.ReplaceAll(exp_nums, string(i), " ")
			}
			exp_nums_slice = strings.Split(exp_nums, " ")
			for _, i := range exp_nums_slice {
				exp_signs = strings.ReplaceAll(exp_signs, i, "")
			}
		}
		result, _ := strconv.ParseFloat(expression, 64) // возвращаем результат решенной строки
		return result, nil
	} else { // если скобки все-таки есть, то берем выражение из этих скобок и с помощью рекурсии решаем его, заменяем на результат операции, после чего программа начинает основные вычисления
		for strings.Count(expression, "(") != 0 {
			var str_comma string
			var str_in_comma string
			for i := strings.Index(expression, "("); i < strings.LastIndex(expression, ")")+1; i++ {
				str_comma = str_comma + string(expression[i])
			}
			for i := strings.Index(expression, "(") + 1; i < strings.LastIndex(expression, ")"); i++ {
				str_in_comma = str_in_comma + string(expression[i])
			}
			solve, _ := Calc(str_in_comma)
			solved_str := strconv.FormatFloat(solve, 'g', -1, 64)
			expression = strings.ReplaceAll(expression, str_comma, solved_str)
		}
	}
	return Calc(expression) // возвращаем значение функции с измененной строкой в ней
}
func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(Calc(a))
}

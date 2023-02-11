package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var priority = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

func InFixToPostFix(expression string) []string {
	var postfix []string
	var operator []string
	for i := 0; i < len(expression); i++ {
		char := string(expression[i])
		switch char {
		case " ":
			continue
		// 左括号优先级最高, 直接入栈
		case "(":
			operator = append(operator, char)
		// 遇到右括号, 弹出操作符, 直到左括号
		case ")":
			for len(operator) > 0 {
				op := operator[len(operator)-1]
				if op == "(" {
					operator = operator[:len(operator)-1]
					break
				}
				postfix = append(postfix, op)
				operator = operator[:len(operator)-1]
			}
		// 操作数情况下
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < len(expression) && unicode.IsDigit(rune(expression[j])); j++ {
				digit += string(expression[j])
			}
			postfix = append(postfix, digit)
			// 注意 j 判断是下一个字符了不为数字情况, 所以需要-1
			i = j - 1
		// + - * / 操作符情况下
		default:
			for len(operator) > 0 {
				op := operator[len(operator)-1]
				// char高优先级直接入栈, 注意此时栈中的 "(" 只有当右括号才处理
				if op == "(" || (priority[char] > priority[op]) {
					break
				}
				postfix = append(postfix, op)
				operator = operator[:len(operator)-1]
			}
			operator = append(operator, char)
		}
	}
	for len(operator) > 0 {
		op := operator[len(operator)-1]
		postfix = append(postfix, op)
		operator = operator[:len(operator)-1]
	}
	return postfix
}

func calculate(postfix []string) int {
	var values []int
	for _, s := range postfix {
		if !strings.ContainsAny(s, "+-*/") {
			d, _ := strconv.Atoi(s)
			values = append(values, d)
		} else {
			op1 := values[len(values)-1]
			op2 := values[len(values)-2]
			values = values[:len(values)-2]
			switch s {
			case "+":
				values = append(values, op2+op1)
			case "-":
				values = append(values, op2-op1)
			case "*":
				values = append(values, op2*op1)
			case "/":
				values = append(values, op2/op1)
			}
		}
	}
	return values[len(values)-1]
}

func main() {
	fmt.Println(InFixToPostFix("2*(9+6/3-5)+4"))
	fmt.Println(calculate(InFixToPostFix("2*(9+6/3-5)+4")))
}

package postfix

import (
	"fmt"
	"github.com/samarec1812/calc-roman-numbers/check"
	"math"
	"strconv"
	"strings"
)

func PriorityOperation(operation string) int {
	switch operation {
	case "+":
		return 1
	case "-":
		return 1
	case "*":
		return 2
	case "/":
		return 2
	case "!":
		return 3
	}
	return 0
}

func PerformToPostfix(expr []string) (string, []string) {
	ToPostfix := make([]string, 0, 0)
	if len(expr) == 1 {
		return expr[0], expr
	}
	operation := make([]string, 0, 0)

	// fmt.Println(expr)
	for i := 0; i < len(expr); i++ {

		if check.IsNumber(expr[i]) {
			chInStr := string(expr[i])
			ToPostfix = append(ToPostfix, chInStr)

		} else if expr[i] == "-" && (i == 0 || expr[i-1] == "(" && (check.IsNumber(expr[i+1]))) {
			//operation = append(operation, expr[i])
			str := expr[i]
			str += expr[i+1]
			ToPostfix = append(ToPostfix, str)
			i++
		} else if check.IsOperationS(expr[i]) {
			if len(operation) == 0 {
				operation = append(operation, expr[i])

			} else if PriorityOperation(operation[len(operation)-1]) >= PriorityOperation(expr[i]) {
				for len(operation) > 0 && (PriorityOperation(operation[len(operation)-1]) >= PriorityOperation(expr[i])) {
					chInStr := string(operation[len(operation)-1])
					operation = operation[:len(operation)-1]
					ToPostfix = append(ToPostfix, chInStr)
				}
				operation = append(operation, expr[i])
			} else {
				operation = append(operation, expr[i])
			}
		} else if expr[i] == "(" {
			operation = append(operation, expr[i])

		} else if expr[i] == ")" {

			for operation[len(operation)-1] != "(" {
				chInStr := operation[len(operation)-1]
				operation = operation[:len(operation)-1]
				ToPostfix = append(ToPostfix, chInStr)
			}
			operation = operation[:len(operation)-1]
		}
	}
	for len(operation) > 0 {
		chInStr := operation[len(operation)-1]
		operation = operation[:len(operation)-1]
		ToPostfix = append(ToPostfix, chInStr)
	}
	strToPostfix := strings.Join(ToPostfix, "")
	return strToPostfix, ToPostfix
}

func calculate(operation string, a int64, b int64) (int64, error) {

	if operation == "+" {
		if a+b < math.MaxInt64 {
			return a + b, nil
		} else {
			return 0, fmt.Errorf("error: overflow int number")
		}
	}
	if operation == "-" {
		if a-b > math.MinInt64 {
			return a - b, nil
		} else {
			return 0, fmt.Errorf("error: overflow int number")
		}
	}
	if operation == "*" {
		if a*b < math.MaxInt64 {
			return a * b, nil
		} else {
			return 0, fmt.Errorf("error: overflow int number")
		}
	}
	if operation == "/" {
		if b == 0 {

			return 0, fmt.Errorf("Error: Division by zero")
		}
		return a / b, nil
	}
	// fmt.Println("Error: operation is not find")
	return 0, fmt.Errorf("Error: operation is not find")
}

func Evalation(ToPostfix []string) (int64, error) {
	result := make([]int64, 0, 0)

	for index := 0; index < len(ToPostfix); index++ {
		//  fmt.Println(result, ToPostfix)
		/* else if ToPostfix[index] == "!" {
			a := result[len(result)-1]
			result = result[:len(result)-1]
			result = append(result, -a)
		} else*/
		if ToPostfix[index] == "+" ||
			ToPostfix[index] == "*" || ToPostfix[index] == "/" || (ToPostfix[index] == "-") {

			a, b := result[len(result)-1], result[len(result)-2]
			result = result[:len(result)-2]
			c, err := calculate(ToPostfix[index], b, a)
			if err != nil {
				return 0, err
			}
			result = append(result, c)
		//	fmt.Println("calculate: ", result, c)
			/*
				else if ToPostfix[index] == "-" && index != len(ToPostfix)-1 {
					fmt.Println(ToPostfix[index], index)

					// if check.IsNumber(ToPostfix[index+1]) && (index == 0) { /* || (index > 0 && (check.IsDigit2(ToPostfix[index-1])||
					//	check.IsLetter2(ToPostfix[index-1]) && check.IsOperation2(ToPostfix[index-2])))) {
						number, _ := strconv.ParseInt(ToPostfix[index+1], 10, 64)
						result = append(result, -number)
						index += 1
					}

				}*/
		} else {
			num, err := strconv.ParseInt(ToPostfix[index], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, num)
		}
	}

	return result[len(result)-1], nil
}

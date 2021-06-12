package check

import (
	"fmt"
	"github.com/samarec1812/calc-roman-numbers/roman"
	"strconv"
)

func CorrectSymbolInString(expr string) bool {
	correctFlag := true
	for index, _ := range expr {
		if !IsRomeNumber(expr[index]) && !IsOperation(expr[index]) && !IsBracket(expr[index]) {
			correctFlag = false
			break
		}
	}
	return correctFlag
}

func IsBracket(ch byte) bool {
	if ch == '(' || ch == ')' {
		return true
	}
	return false
}

func CorrectString(expr string) bool {
	if expr == "" {
		fmt.Println("error: Empty input string")
		return false
	}
	if expr[0] == '+' || expr[0] == '*' || expr[0] == '/' {
		fmt.Println("error: first symbol is operation")
		return false
	}
	if len(expr) == 1 && IsOperation(expr[0]) {
		fmt.Println("error: string len equal 1 and first symbol is operation")
		return false
	}
	balance := 0
	if expr[0] == '(' {
		balance += 1
	}
	if expr[0] == ')' {
		fmt.Println("error: first symbol is right bracket")
		return false
	}
	if !CorrectBracket(FormBracketStr(expr)) {
		fmt.Println("error: no equal left and right brackets")
		return false
	}
	for i := 1; i < len(expr)-1; i++ {
		switch expr[i] {
		case '(':
			balance += 1
			if !IsOperation(expr[i-1]) && expr[i-1] == ')' {
				fmt.Println("error 1")
				return false
			}
			if !IsRomeNumber(expr[i+1]) && expr[i+1] == ')' {
				fmt.Println("error: 1.1")
				return false
			}
		case ')':
			balance--

			if (!IsRomeNumber(expr[i-1])) && (expr[i-1] == '(' || IsOperation(expr[i-1])) {
				fmt.Println("error 2.2")
				return false
			}
			if !IsOperation(expr[i+1]) && expr[i+1] == '(' {
				fmt.Println("error 2.3")
				return false
			}
		case '+':
			if !IsRomeNumber(expr[i-1]) && (expr[i-1] == '(' || IsOperation(expr[i-1])) {
				fmt.Println("error 3")
				return false
			}
			if !IsRomeNumber(expr[i+1]) && (expr[i+1] == ')' || IsOperation(expr[i+1])) {
				fmt.Println("error 4")
				return false
			}
		case '-':
			if IsOperation(expr[i-1]) {
				// !IsLetter(expr[i-1]) && !IsDigit(expr[i-1]) && expr[i-1] != '('  {
				fmt.Println("error 5")
				return false
			}
			if !IsRomeNumber(expr[i+1]) && expr[i+1] == ')' {
				fmt.Println("error 6")
				return false
			}
		case '*':
			if !IsRomeNumber(expr[i-1]) && expr[i-1] == '(' {
				fmt.Println("error 7")
				return false
			}
			if !IsRomeNumber(expr[i+1]) && expr[i+1] == ')' {
				fmt.Println("error 8")
				return false
			}
		case '/':
			if !IsRomeNumber(expr[i-1]) && (expr[i-1] == '(' || IsOperation(expr[i-1])) {
				fmt.Println("error 9")
				return false
			}
			if !IsRomeNumber(expr[i+1]) && (expr[i+1] == ')' || IsOperation(expr[i+1])) {
				fmt.Println("error 10")
				return false
			}
		}
	}
	if !IsRomeNumber(expr[len(expr)-1]) && expr[len(expr)-1] != ')' {
		fmt.Println("error 11 " + string(expr[len(expr)-1]))
		return false
	}

	if expr[len(expr)-1] == ')' {
		balance -= 1
	}
	if balance != 0 {
		fmt.Println("error 12")
		return false
	}
	return true
}

func IsRomeNumber(ch byte) bool {
	if _, ok := roman.Dict[string(ch)]; ok {
		return true
	}
	return false
}

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func IsOperationS(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func IsOperation(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func CreateToken(inputStr string) []string {
	tokenStr := make([]string, 0, 0)
	romeStr := ""
	for idx, value := range inputStr {

		if value == '(' || value == ')' || value == '+' || value == '-' || value == '*' || value == '/' || idx == len(inputStr)-1 {
			if romeStr != "" {
				tokenStr = append(tokenStr, romeStr)
			}

			tokenStr = append(tokenStr, string(value))
			romeStr = ""

		} else {
			romeStr += string(value)
		}

	}
	return tokenStr
}
func FormBracketStr(s string) string {
	outStr := ""
	for _, value := range s {
		if value == '(' || value == ')' {
			outStr += string(value)
		}
	}
	return outStr
}

func CorrectBracket(s string) bool {
	stack := make([]byte, 0, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		}
		if len(stack) == 0 {
			return false
		}
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func CheckRomanToken(tokens []string) bool {
	flag := true
	for _, value := range tokens {
		if value == "(" || value == ")" || value == "+" || value == "-" || value == "*" || value == "/" {
			continue
		} else {

			if !roman.IsRoman(value) {
				flag = false
				fmt.Println("This value is ", value)
				break
			}
		}
	}
	return flag
}

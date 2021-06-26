package main

import (
	"bufio"
	"fmt"
	"github.com/samarec1812/calc-roman-numbers/check"
	"github.com/samarec1812/calc-roman-numbers/postfix"
	"github.com/samarec1812/calc-roman-numbers/roman"
	"os"
	"strings"
)

func Processing(expr string) string {
	s := strings.ReplaceAll(expr, " ", "")
	return s
}

func main() {

	for {
		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		inputStr := myscanner.Text()
		// end input stream
		if inputStr == "" {
			break
		}
		inputStr = Processing(inputStr)

		if check.CorrectSymbolInString(inputStr) {
			if check.CorrectString(inputStr) {
				tokens := check.CreateToken(inputStr)

				if check.CheckRomanToken(tokens) {
					for index, val := range tokens {
						if roman.IsRoman(val) {
							tokens[index] = roman.ConvertRoman(val)
						}
					}

					_, postfixForm := postfix.PerformToPostfix(tokens)
					arabic, err := postfix.Evalation(postfixForm)
					if err != nil {
						fmt.Println(err)

					} else {
						fmt.Println(roman.ConvertArabic(arabic))
					}
				} else {
					fmt.Println("error: incorrect roman number")

				}
			} else {

			}
		} else {
			fmt.Println("error: incorrect symbol in string")

		}
	}

}

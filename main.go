package main

import (
	"./check"
	"./postfix"
	"./roman"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Processing(expr string) string {
	s := strings.ReplaceAll(expr, " ", "")
	return s
}

func main() {
	fmt.Println(roman.ConvertRoman("XXIX"), roman.ConvertRoman("XLIX"))
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	inputStr := myscanner.Text()

	inputStr = Processing(inputStr)
	fmt.Println("_________________")
	fmt.Println(inputStr)


	if check.CorrectSymbolInString(inputStr) {
		fmt.Println(check.CorrectString(inputStr))
		tokens := check.CreateToken(inputStr)
		fmt.Println(tokens)

		if check.CheckRomanToken(tokens) {
			fmt.Println("GOOD")
			for index, val := range tokens {
				if roman.IsRoman(val) {
					tokens[index] = roman.ConvertRoman(val)
				}
			}
			fmt.Println(tokens)
			_, postfixForm := postfix.PerformToPostfix(tokens)
			fmt.Println(postfixForm)
			arabic, err := postfix.Evalation(postfixForm)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(roman.ConvertArabic(arabic))

			// fmt.Println(math.MaxInt64)

		} else {
			fmt.Println("Incorrect roman number")
		}
	} else {
		fmt.Println("Incorrect str")

	}


	/*for _, val := range tokens {
		fmt.Println(val)
	}*/



	// fmt.Println(roman.IsRoman(inputStr))
}

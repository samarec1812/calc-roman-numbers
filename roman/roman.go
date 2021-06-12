package roman

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Dict = map[string]int{
	"Z": 0,
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var StructNum = []struct {
	numInt int64
	roman  string
}{

	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "IX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

func IsRoman(roman string) bool {
	if roman == "" {
		return false
	}

	romanByte := []byte(strings.ToUpper(roman))
	check, _ := regexp.Match("^Z{1}$|^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", romanByte)
	return check
}

func ConvertRoman(roman string) string {
	result := 0
	for i := 0; i < len(roman)-1; i++ {
		if Dict[string(roman[i])] < Dict[string(roman[i+1])] {
			result = result - Dict[string(roman[i])]
		} else if Dict[string(roman[i])] >= Dict[string(roman[i+1])] {
			result += Dict[string(roman[i])]
		}
	}
	result += Dict[string(roman[len(roman)-1])]
	return strconv.Itoa(result)
}

func ConvertArabic(num int64) string {
	// fmt.Println(num)
	sort.Slice(StructNum, func(i, j int) bool {

		return StructNum[i].numInt > StructNum[j].numInt
	})

	roman := ""
	if num < 0 {
		num = (-num)
		roman += "-"
	} else if num == 0 {
		return "Z"
	}
	// fmt.Println(num)
	for _, value := range StructNum {
		for num >= value.numInt {
			roman += value.roman
			num -= value.numInt
		}

	}
	return roman
}

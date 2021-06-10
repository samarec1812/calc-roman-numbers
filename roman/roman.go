package roman

import (
	"regexp"
	"strings"
)

func IsRoman(roman string) bool {
	if roman == "" {
		return false
	}

	romanByte := []byte(strings.ToUpper(roman))
	check, _ := regexp.Match("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", romanByte)
	return check
}


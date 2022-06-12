package word

import (
	"strings"
	"unicode"
)

const (
	ModeUpper = iota
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamleCaseToUnderScore
)

//字符串转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//字符串转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线转空格，小写转大写,如test_case_go转换为TestCaseGO,大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s) //首字母换成大写
	return strings.Replace(s, " ", "", -1)
}

//下test_case_go转换为testCaseGO,小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//将头字符大写改小写，大写前面加_
func CamleCaseToUnderScore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

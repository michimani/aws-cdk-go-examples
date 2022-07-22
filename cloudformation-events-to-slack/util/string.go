package util

import (
	"strings"
)

// toUpperCamelCase converts a string to an Upper Camel Case.
// e.g) "hello-aws-cdk-golang" -> "HelloAwsCdkGolang"
func ToUpperCamelCase(base string) string {
	base = strings.ReplaceAll(base, "-", "_")

	sp := strings.Split(base, "_")
	res := ""
	for _, s := range sp {
		res += strings.ToUpper(s[0:1]) + s[1:]
	}

	return res
}

// ToKebabCase converts a string to an Kebab Case.
// e.g) "HelloAwsCdkGolang" -> "hello-aws-cdk-golang"
func ToKebabCase(base string) string {
	tr := rune(' ')
	hp := rune('-')
	uA := rune('A')
	uZ := rune('Z')

	resRune := []rune{}
	for _, r := range base {
		if r >= uA && r <= uZ {
			if len(resRune) != 0 {
				resRune = append(resRune, hp)
			}
			resRune = append(resRune, r+tr)
		} else {
			resRune = append(resRune, r)
		}
	}

	return string(resRune)
}

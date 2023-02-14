package libs

import "strings"

// ToUpperCamelCase converts a string to an Upper Camel Case.
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

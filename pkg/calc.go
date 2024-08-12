package pkg

import "strings"

func Test(x int) int {
	return x * x
}

func MoreTest(str string) string {
	return strings.ToUpper(str)
}

func SayHello(name string) string {
	return "Hello " + name
}

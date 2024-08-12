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

func TwoSum(a, b int) int {
	return a + b
}

func ArrSum(arr []int) int {
	result := 0
	for _, value := range arr {
		result += value
	}
	return result
}

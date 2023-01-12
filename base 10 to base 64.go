package main

import (
	"fmt"
)

func base12ToBase64(num int) string {
	
	base64Chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	result := ""

	
	for num > 0 {
		remainder := num % 64
		result = string(base64Chars[remainder]) + result
		num = num / 64
	}

	return result
}

func main() {
	
	num := 12345
	fmt.Println(base12ToBase64(num))
}

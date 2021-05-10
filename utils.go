package utils

import (
	"fmt"
	"strings"
)

func ParseParamsFromUrl(url string) []string {
	var urlParams []string
	parts := strings.Split(url, "?")[1]
	params := strings.Split(parts, "&")
	for _, param := range params {
		urlParams = append(urlParams, strings.Split(param, "=")[1])
	}
	return urlParams
}

func GetMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func privateUtil() {
	fmt.Println("I am private")
}

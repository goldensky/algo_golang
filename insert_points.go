package main

import (
	"fmt"
	"strings"
)

func InsertPoints(s string) string {
	isNegative := false
	fmt.Println(isNegative)
	if strings.HasPrefix(s, "-") {
		isNegative = true
		s = s[1:]
	}
	parts := len(s) / 3
	index := len(s) % 3
	builder := strings.Builder{}
	if isNegative {
		builder.WriteString("-")
	}
	if index > 0 {
		builder.WriteString(s[:index])

	}
	if index > 0 && len(s) > 3 {
		builder.WriteString(".")
	}

	for i := 0; i < parts; i++ {
		builder.WriteString(s[index : index+3])
		index += 3
		if i < parts-1 {
			builder.WriteString(".")
		}

	}
	return builder.String()

}

//s := "12"
//s := "1234567"
////s := "123456"
//s = "-1223456"
//fmt.Println(InsertPoints(s))

func NumberWithSeparators(s string) string {
	builder := strings.Builder{}

	if strings.HasPrefix(s, "-") {
		builder.WriteString("-")
		s = s[1:]
	}

	index := len(s) % 3
	parts := len(s) / 3

	if index > 0 {
		builder.WriteString(s[:index])
	}
	if index > 0 && len(s) > 3 {

		builder.WriteString(".")
	}

	for i := 0; i < parts; i++ {
		builder.WriteString(s[index : index+3])
		index += 3
		if i < parts-1 {
			builder.WriteString(".")
		}
	}
	return builder.String()
}

/*
Напишите функцию numberWithSeparators, которая преобразует строку s с целым числом в строку с разделителями
тысяч в виде точек. Используй strings.Builder.

Пример:

Ввод: s = "1234567"
Вывод: "1.234.567"
Пример:

Ввод: s = "-12"
Вывод: "-12"

*/

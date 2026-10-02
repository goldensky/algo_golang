package main

import (
	//"fmt"
	"unicode"
)

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	if len(runes) <= 1 {
		return true
	}

	lo, hi := 0, len(runes)-1
	for lo < hi {
		for lo < hi && !isAlnum(runes[lo]) {
			lo++
		}

		for lo < hi && !isAlnum(runes[hi]) {
			hi--
		}
		if unicode.ToUpper(runes[lo]) != unicode.ToUpper(runes[hi]) {
			return false
		}
		lo++
		hi--
	}
	return true
}

//var testStr = "A b C, пп  c, B a, "
//fmt.Println(isPalindrome(testStr))

/*
Дана строка s. Необходимо реализовать функцию IsPalindrome, которая возвращает true,
если строка является палиндромом, и false в противном случае.

Строка является палиндромом, если она читается одинаково в обоих направлениях
(слева направо и справа налево), игнорируя регистр букв и не учитывая пробелы и знаки пунктуации.

Пример 1:

Ввод: s = ""
Вывод: true
Пример 2:

Ввод: s = "abA"
Вывод: true
Пример 3:

Ввод: s = "Was it a car or a cat I saw?"
Вывод: true
Пример 4:

Ввод: s = "привет"
Вывод: false


*/

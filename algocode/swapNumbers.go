package main

func swapNumbers(val1 *int, val2 *int) {
	temp := *val1
	*val1 = *val2
	*val2 = temp

}

//a := 12
//b := 45
//fmt.Println(a, b)
//swapNumbers(&a, &b)
//fmt.Println(a, b)

/*
Даны две ссылки val1 и val2 на целые числа. Нужно реализовать функцию Swap,
которая меняет их значения местами.

Считать, что указатель не может быть nil.

Пример:

Ввод: val1 = 20, val2 = 10; Swap(&val1, &val2)
Вывод: val1 = 10, val2 = 20

*/

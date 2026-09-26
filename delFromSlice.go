package main

func delFromSlice(arr []int, pos int) []int {
	if pos >= len(arr) || pos < 0 {
		new_arr := make([]int, len(arr))
		copy(new_arr, arr)
		return new_arr
	}
	new_arr := make([]int, 0, len(arr)-1)
	new_arr = append(new_arr, arr[:pos]...)
	new_arr = append(new_arr, arr[pos+1:]...)

	return new_arr
}

//arr := []int{1, 2, 3, 4, 5, 6}
//fmt.Println(delFromSlice(arr, 20))

/*

Дан слайс целых чисел values и целое число pos. Необходимо написать функцию removeElement,
которая возвращает новый слайс с удаленным элементом, стоявщим на позиции pos исходного слайса.

Если pos выходит за границы слайса (отрицательный или ≥ длины), необходимо вернуть копию исходного слайса.

Пример 1:

Ввод: values = [1, 2, 3, 4, 5], pos = 1
Вывод: [1, 3, 4, 5]
Пример 2:

Ввод: values = [1, 2, 3, 4, 5], pos = 5
Вывод: [1, 2, 3, 4, 5]

*/

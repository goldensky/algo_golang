package main

/*
Необходимо реализовать структуру Set для хранения уникальных целых чисел с методами:

NewSet(capacity int) — создание с начальной емкостью
Add(values...) — добавление элементов (игнорирует дубли)
Remove(value) — удаление элемента
Contains(value) — проверка наличия
Len() — текущее количество
ToSlice() — возврат элементов в виде слайса
При реализации можно использовать встроенный тип map.

Пример:
c
s := NewSet(10)
s.Add(1, 2, 2) // {1, 2}
s.Contains(1)  // true
s.ToSlice()    // [1, 2] (порядок не фиксирован)

type Set struct{}

	func NewSet(length int) *Set {
		// Создаем новый set с заданной капасити
	}

	func (s *Set) Add(value ...int) {
		// Добавляем новые значения
	}

	func (s *Set) Remove(value int) {
		// Удаляем значение
	}

	func (s *Set) Contains(value int) bool {
		// Проверяем наличие элемента
	}

	func (s *Set) Len() int {
		// Количество элементов
	}

	func (s *Set) ToSlice() []int {
		// Получить список элементов (порядок не важен)
	}
*/
type void struct{}

type Set struct {
	data map[int]void
}

func NewSet(length int) *Set {
	// Создаем новый set с заданной капасити
	return &Set{
		data: make(map[int]void, length),
	}
}

func (s *Set) Add(value ...int) {
	// Добавляем новые значения

	for _, v := range value {
		s.data[v] = void{}
	}
}

func (s *Set) Remove(value int) {
	// Удаляем значение
	delete(s.data, value)
}

func (s *Set) Contains(value int) bool {
	// Проверяем наличие элемента
	_, ok := s.data[value]
	return ok
}

func (s *Set) Len() int {
	// Количество элементов
	return len(s.data)
}

func (s *Set) ToSlice() []int {
	// Получить список элементов (порядок не важен)
	values := make([]int, 0, len(s.data))

	for k, _ := range s.data {
		values = append(values, k)
	}
	return values
}

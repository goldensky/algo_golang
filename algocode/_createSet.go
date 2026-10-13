package main

/*


В языке Go выражение struct{}{} используется для создания пустого значения (экземпляра) пустой структуры.
Две пары фигурных скобок {}{} нужны потому, что первая пара определяет тип данных,
а вторая — инициализирует значение этого типа.

Подробный разбор: почему скобок две?

1. struct{} — это тип данных (чертеж)
Это объявление структуры, в которой нет ни одного поля. Она занимает 0 байт в памяти.
2. {} — это создание значения (заполнение чертежа)
В Go для создания любого объекта (структуры, массива, слайса) используются фигурные скобки,
внутри которых перечисляются начальные значения полей.
Так как у нашей структуры полей нет, скобки остаются пустыми.

Аналогия с обычной структурой:
go
// 1. Сначала определяем тип (чертеж)
type User struct { Name string }

// 2. Потом создаем значение этого типа
var u = User{Name: "Alice"}

// А теперь делаем то же самое, но "на лету" (без имени типа):
var anonymous = struct{ Name string }{Name: "Alice"}
//              |--- ТИП ДАННЫХ ---||-- ЗНАЧЕНИЕ --|

Если мы уберем из этого примера поле Name, то получим:
go
var empty = struct{}{"Здесь ничего нет, поэтому скобки пустые"} // Ошибка, если писать текст
var data  = struct{}{}
//          | ТИП ||ЗНАЧЕНИЕ|


Зачем так делать в map?

Комбинация map[int]struct{} часто используется в Go для реализации множеств (Set),
где важны только уникальные ключи, а значения не имеют значения.
Сравнение с map[int]bool:
• data[4] = true — тип bool занимает в памяти 1 байт для каждого элемента.
• data[1] = struct{}{} — пустая структура struct{} занимает 0 байт в памяти.
Если в вашей карте будут миллионы элементов, использование struct{} вместо bool
сэкономит оперативную память, так как Go не будет выделять место под значения.



*/

func createSet() {
	data := make(map[int]struct{}, 5)
	data[1] = struct{}{}
	data[2] = struct{}{} // почему для структуры пишем struct{}{} -  две пары {}?
	data[3] = struct{}{}
	// data[4] = true // для map[int]bool
	fmt.Println(data)

}

/*


Вот как сделать код более читаемым с помощью собственного типа и как правильно проверять,
есть ли элемент в таком множестве.

1. Делаем код красивым через type

Писать каждый раз struct{}{} неудобно, и это перегружает код.
В Go принято создавать псевдоним типа (alias) или новый тип для пустого значения.
Это делает код чистым и понятным с первого взгляда.
go
package main

import "fmt"

// Ограничиваем область видимости: создаем понятный тип-пустышку
type void struct{}

func main() {
    // Теперь код читается как "карта чисел, внутри которой пустота"
    data := make(map[int]void, 5)

    // Вместо struct{}{} пишем понятное void{}
    data[1] = void{}
    data[2] = void{}
    data[3] = void{}

    fmt.Println(data) // Выведет: map[1:{} 2:{} 3:{}]
}
Use code with caution.

2. Как проверять наличие элемента (Идиома value, ok)

В случае с map[int]bool вы могли бы проверить элемент простым условием if data[1] { ... }.
С пустой структурой так сделать нельзя, потому что она не возвращает логическое true или false.
В Go для этого используют специальную конструкцию value, ok,
где первая переменная — это само значение (в нашем случае struct{}), а вторая (ok) — это bool,
который показывает, существует ли ключ в карте.
go
package main

import "fmt"

type void struct{}

func main() {
    data := make(map[int]void)
    data[42] = void{}

    // Проверяем, есть ли число 42 в нашем множестве
    // Переменную со значением игнорируем с помощью нижнего подчеркивания (_)
    if _, exists := data[42]; exists {
        fmt.Println("Число 42 есть в множестве!")
    } else {
        fmt.Println("Числа 42 там нет.")
    }

    // Проверка отсутствующего элемента
    if _, exists := data[99]; !exists {
        fmt.Println("Числа 99 точно нет в карте.")
    }
}
Use code with caution.

Главные фишки этого подхода:

• Читаемость: Код void{} сразу говорит разработчику, что значение не играет роли.
• Эффективность: Проверка _, exists := data[key] работает максимально быстро (за O(1)),
так как Go не нужно копировать данные из памяти — их просто нет.


*/

//======================================================================================================
//======================================================================================================
//======================================================================================================

/*

можем разобрать:
• Как правильно удалять элементы из такого множества
• Как написать удобные функции-помощники Add, Has и Remove, чтобы работать с картой как с полноценным объектом Set
• В каких редких случаях map[int]bool все-таки выигрывает у пустой структуры


Вот детальный разбор всех трех тем:
удаление элементов,
создание полноценного типа Set (Множество)
и сравнение с map[int]bool.

1. Как правильно удалять элементы

Для удаления элементов из карты (неважно, с пустой структурой или любым другим типом)
в Go используется встроенная функция delete.
go
package main

import "fmt"

type void struct{}

func main() {
    data := make(map[int]void)
    data[1] = void{}
    data[2] = void{}

    // Удаляем ключ 1
    delete(data, 1)

    // Если ключа нет в карте, delete ничего не сделает и не вызовет ошибку
    delete(data, 99)

    fmt.Println(data) // Выведет только: map[2:{}]
}


2. Создаем полноценный объект Set (Множество)

Чтобы не писать логику карт вручную по всему проекту, в Go принято инкапсулировать её в отдельный тип.
Давайте напишем чистый, поддерживаемый тип IntSet с методами Add, Has (проверка) и Remove.
go
package main

import "fmt"

type void struct{}

// IntSet оборачивает карту для работы как с множеством
type IntSet struct {
    elements map[int]void
}

// NewIntSet — конструктор для инициализации множества
func NewIntSet(size int) *IntSet {
    return &IntSet{
        elements: make(map[int]void, size),
    }
}

// Add добавляет элемент в множество
func (s *IntSet) Add(value int) {
    s.elements[value] = void{}
}

// Has проверяет, существует ли элемент в множестве
func (s *IntSet) Has(value int) bool {
    _, exists := s.elements[value]
    return exists
}

// Remove удаляет элемент из множества
func (s *IntSet) Remove(value int) {
    delete(s.elements, value)
}

func main() {
    // Использование:
    set := NewIntSet(5)

    set.Add(10)
    set.Add(20)

    fmt.Println("Есть 10?:", set.Has(10)) // true
    fmt.Println("Есть 30?:", set.Has(30)) // false

    set.Remove(10)
    fmt.Println("Есть 10 после удаления?:", set.Has(10)) // false
}
Use code with caution.

3. Когда map[int]bool все-таки выигрывает у пустой структуры?

Несмотря на то, что struct{} экономит память (0 байт против 1 байта у bool),
в коммерческой разработке вы часто будете встречать именно map[int]bool.
Вот ситуации, когда bool оказывается лучше:

А. Простота и лаконичность кода (Меньше «мусора»)

Проверка существования ключа в map[int]bool пишется в одну строчку без использования конструкции value, ok:
go
// С использованием bool:
if data[42] {
    // Код выполнится, если ключ 42 существует и равен true
}

// С использованием struct{}:
if _, exists := data[42]; exists {
    // Код длиннее и содержит «заглушку» в виде нижнего подчеркивания
}
Use code with caution.
Важное примечание: Этот трюк с if data[42] работает безопасно только в том случае,
если вы никогда не записываете в карту значение false, а используете её строго как множество (
записывая только true или удаляя ключ через delete). Если в карте окажется ключ со значением false,
проверка if data[42] вернет false, хотя сам ключ в карте присутствует.

Б. Переиспользование стандартных пакетов

Некоторые встроенные механизмы (например, JSON-сериализация через json.Marshal) гораздо «охотнее»
и понятнее для фронтенда или внешних систем превращают map[string]bool в обычный объект с флагами true/false,
чем карту со странными пустыми объектами {} внутри.

В. Когда экономия памяти незначительна

Если в вашей карте планируется хранить до 10 000 – 50 000 элементов, разница в памяти между struct{} и bool
составит всего несколько десятков килобайт. В масштабах современных серверов эта экономия ничтожна.
В таких случаях разработчики часто выбирают bool ради более чистого и привычного синтаксиса.



*/

/*

можем углубиться дальше. Например, я могу показать:
• Как сделать это множество потокобезопасным
(Thread-safe) с помощью sync.RWMutex, чтобы с ним можно было работать из горутин.
• Как переписать этот Set на Generics (Go 1.18+), чтобы он работал не только с int,
но и с string или любыми другими типами.

Вот детальный разбор двух продвинутых тем: превращение нашего множества в универсальное (через Generics) и обеспечение его безопасности при работе из параллельных потоков (Thread-safety).

1. Делаем множество универсальным (Generics в Go 1.18+)

Раньше для int и string приходилось писать два разных класса. Теперь с помощью дженериков можно создать один Set, который будет работать с любым типом данных, поддерживающим сравнение (числа, строки, указатели).
Для этого в Go используется встроенное ограничение comparable.
go
package main

import "fmt"

type void struct{}

// Set использует дженерик [T comparable], что позволяет передавать любой сравниваемый тип
type Set[T comparable] struct {
	elements map[T]void
}

// NewSet инициализирует множество для конкретного типа T
func NewSet[T comparable](size int) *Set[T] {
	return &Set[T]{
		elements: make(map[T]void, size),
	}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = void{}
}

func (s *Set[T]) Has(value T) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func main() {
	// Создаем множество строк (string)
	stringSet := NewSet[string](3)
	stringSet.Add("golang")
	stringSet.Add("docker")
	fmt.Println("Есть golang?:", stringSet.Has("golang")) // true

	// Создаем множество целых чисел (int)
	intSet := NewSet[int](3)
	intSet.Add(42)
	fmt.Println("Есть 42?:", intSet.Has(42)) // true
}
Use code with caution.

2. Делаем множество потокобезопасным (Thread-safe)

В Go встроенные карты (map) категорически запрещено одновременно читать и писать из разных горутин. Если две горутины попытаются одновременно записать данные в одну карту, программа завершится с критической ошибкой (panic: concurrent map writes).
Чтобы защитить данные, мы используем sync.RWMutex (Read-Write Mutex). Он работает эффективнее обычного sync.Mutex, так как:
• Позволяет множеству горутин одновременно читать данные (метод RLock), не блокируя друг друга.
• Разрешает только одной горутине писать или удалять данные (метод Lock), полностью блокируя всех остальных на это время.
go
package main

import (
	"fmt"
	"sync"
)

type void struct{}

// ConcurrentSet объединяет наше множество с мьютексом для защиты данных
type ConcurrentSet[T comparable] struct {
	sync.RWMutex // Встраиваем мьютекс прямо в структуру
	elements     map[T]void
}

func NewConcurrentSet[T comparable]() *ConcurrentSet[T] {
	return &ConcurrentSet[T]{
		elements: make(map[T]void),
	}
}

// Add требует эксклюзивной блокировки на запись
func (s *ConcurrentSet[T]) Add(value T) {
	s.Lock()         // Блокируем карту для всех остальных
	defer s.Unlock() // Гарантированно разблокируем в конце функции
	s.elements[value] = void{}
}

// Remove также требует блокировки на запись
func (s *ConcurrentSet[T]) Remove(value T) {
	s.Lock()
	defer s.Unlock()
	delete(s.elements, value)
}

// Has использует RLock, позволяя другим горутинам читать параллельно
func (s *ConcurrentSet[T]) Has(value T) bool {
	s.RLock()         // Блокируем только для записи, читать могут все
	defer s.RUnlock() // Освобождаем "читательскую" блокировку
	_, exists := s.elements[value]
	return exists
}

func main() {
	set := NewConcurrentSet[int]()
	var wg sync.WaitGroup

	// Запускаем 1000 горутин, которые одновременно пишут в наше множество
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			set.Add(val)
		}(i)
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("Успешно добавлено элементов:", len(set.elements))
	fmt.Println("Есть ли элемент 500?:", set.Has(500))
}









*/

type void struct{}
type Set struct {
	data map[int]void
}

func NewSet(length int) *Set {
	//newSet := &Set{data: make(map[int]void, length)}

	newSet := &Set{
		data: make(map[int]void, length),
	}
	return newSet
}

func (s *Set) Add(value ...int) {
	for index, item := range value {
		fmt.Println("--", index, item)
		s.data[item] = void{}
	}
}

func (s *Set) Has(value int) bool {
	_, ok := s.data[value]
	return ok
}

func (s *Set) Remove(value int) {
	delete(s.data, value)
}

func (s *Set) Contains(value int) bool {
	_, ok := s.data[value]
	return ok
}
func (s *Set) Values() []int {
	values := make([]int, 0, len(s.data))
	for k, v := range s.data {
		fmt.Println(k, v)
		values = append(values, k)
	}
	fmt.Println(values)
	return values
}

//set := NewSet(4)
//set.Add(1, 2, 3)
//fmt.Println(set.Values())
//fmt.Println(set.Has(3))
//fmt.Println(set.Has(13))

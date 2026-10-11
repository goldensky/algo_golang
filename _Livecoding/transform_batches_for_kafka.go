package main

import (
	"fmt"
	"math/rand"
)

/*

Нужно передать данные из некоторого источника некоторому потребителю.
При этом источник отдает данные небольшими пачками (~ десятки записей),
а потребитель оптимальнее работает с крупными батчами (~тысячи записей).

Реальный пример -  поставка данных из очередей типа Kafka  в базу Clickhouse.

Источник:
 - условно бесконечный
 - Источник никогда не возвращает более MaxItems записей за один вызов Next.
 - в рамках одной сессии  (одного вызова функции Pipe) источник каждый раз возвращает
 новые данные на каждый вызов Next.
 - Однако, после перезапуска источник начнет с прошлой "подтвержденной" позиции, задаваемой cookie.

Поэтому каждое значение cookie, которое вернул вызов Next, после сохранения данных в приемнике,
должно быть фиксировано вызовом Commit, причем строго в той же последовательности, в которой их вернул Next.


Приемник:
 - не может обработать более MaxItems за один раз

 Базовый уровень:
Требуется реализовать функцию func Pipe(p Producer, c Concumer) error, которая
читает данные из источника,
группирует их в буфер размером не более MaxItems
сохраняет в приемник,
после чего фиксирует прогресс в источнике

==================================

pending_cookies - список cookie, которые относятся к элементам, полностью лежащим в текущем big_batch.
Читаем пачку из Producer.
Если она целиком помещается в big_batch - добавляем, а её cookie кладем в pending_cookies.
Если она не помещается:
Дополняем big_batch до максимума.
Отправляем big_batch в Consumer.
Если успешно - коммитим все pending_cookies (они все гарантированно сохранены).
Очищаем big_batch и pending_cookies.
Оставшийся «хвост» кладем в новый big_batch.

Важно: cookie от этой пачки мы пока не коммитим, но так как хвост уже в новом big_batch,
мы сразу добавляем этот cookie в pending_cookies для нового батча.
Когда источник иссякнет, если остался непустой big_batch - отправляем его и коммитим оставшиеся pending_cookies.


type Item int

type Producer interface {
}

type Consumer interface {
	consume() error
}

func (c Consumer) consume(bigBatch []Item) {
	fmt.Println("consume big batch === ", c, bigBatch)
}
*/

type Item int

const MaxBatchSize = 100

// cookie - position marker
type Producer interface {
	Next() (items []Item, cookie int, err error)
	Commit(cookie []int) error
}

type Consumer interface {
	Process(items []Item) error
}

//=======================================================

type ProducerSource struct {
	producedCount    int // cookie
	limit            int
	committedCookies []int // что было закоммичено
}

func (p *ProducerSource) Next() ([]Item, int, error) {
	if p.producedCount >= p.limit {
		return nil, 0, nil
	}

	batchSize := rand.Intn(16) + 10

	if p.producedCount+batchSize > p.limit {
		batchSize = p.limit - p.producedCount
	}

	items := make([]Item, batchSize)
	for i := 0; i < batchSize; i++ {
		items[i] = Item(rand.Intn(1000))
		p.producedCount++
	}

	return items, p.producedCount, nil

}

func (p *ProducerSource) Commit(cookie []int) error {
	if len(cookie) == 0 {
		return nil
	}
	p.committedCookies = append(p.committedCookies, cookie...)
	fmt.Println("committed cookies:", p.committedCookies)
	return nil
}

// =========================================================
type ConsumerStorage struct {
	storedItems []Item
}

func (c *ConsumerStorage) Process(items []Item) error {
	if len(items) == 0 {
		return nil
	}
	c.storedItems = append(c.storedItems, items...)
	return nil
}

func Pipe(p Producer, c Consumer) error {
	// Инициализируем срезы с емкостью (capacity), но длиной (length) 0
	bigBatch := make([]Item, 0, MaxBatchSize)
	var pendingCookies []int

	for {
		items, cookie, err := p.Next()
		if err != nil {
			return fmt.Errorf("ошибка чтения из источника: %w", err)
		}

		// Пустой срез без ошибки - это стандартный сигнал EOF (End Of File/Stream) в Go
		if len(items) == 0 {
			break
		}

		spaceLeft := MaxBatchSize - len(bigBatch)

		if len(items) <= spaceLeft {
			// Пачка целиком помещается в текущий bigBatch
			bigBatch = append(bigBatch, items...)
			pendingCookies = append(pendingCookies, cookie)
		} else {
			// Пачка НЕ помещается целиком. Разрезаем её.
			bigBatch = append(bigBatch, items[:spaceLeft]...)
			remaining := items[spaceLeft:]

			// 1. Отправляем заполненный bigBatch в Consumer
			if err := c.Process(bigBatch); err != nil {
				return fmt.Errorf("ошибка сохранения в приемник: %w", err)
			}

			// 2. Коммитим ВСЕ cookie, которые полностью ушли в этот обработанный bigBatch
			if err := p.Commit(pendingCookies); err != nil {
				return fmt.Errorf("ошибка коммита в источнике: %w", err)
			}

			// 3. Начинаем формировать новый bigBatch с "хвоста"
			bigBatch = remaining

			// 4. Cookie от этой "разрезанной" пачки теперь относится к элементам
			// в новом bigBatch. Мы не коммитим его сейчас, но запоминаем.
			pendingCookies = []int{cookie}
		}
	}

	// === ФИНАЛЬНЫЙ СБРОС (Final Flush) ===
	// Источник иссяк, но в bigBatch могли остаться необработанные данные
	if len(bigBatch) > 0 {
		if err := c.Process(bigBatch); err != nil {
			return fmt.Errorf("ошибка финального сохранения: %w", err)
		}
		if err := p.Commit(pendingCookies); err != nil {
			return fmt.Errorf("ошибка финального коммита: %w", err)
		}
	}

	return nil
}
func main() {
	c := &ConsumerStorage{}
	p := &ProducerSource{limit: 230}

	err := Pipe(p, c)
	fmt.Println(err)

	fmt.Println(p.producedCount)
	fmt.Println(len(c.storedItems))
	fmt.Println(p.committedCookies)
}

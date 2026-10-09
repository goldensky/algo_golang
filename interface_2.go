package main

import (
	"errors"
	"fmt"
)

type employee struct {
	name   string
	age    string
	id     int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
	update(e employee, salary int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (m *memoryStorage) insert(e employee) error {
	m.data[e.id] = e

	return nil
}

func (m *memoryStorage) get(id int) (employee, error) {
	e, exists := m.data[id]
	if !exists {
		return employee{}, errors.New("employee not found")
	}

	return e, nil
}

func (m *memoryStorage) delete(id int) error {
	e, exists := m.data[id]
	if !exists {
		return errors.New("employee not found")
	}
	fmt.Println("deleting employee", e.name)
	delete(m.data, id)
	return nil
}

func (m *memoryStorage) update(e employee, new_salary int) error {
	id, exists := m.data[e.id]

	//e, exists := m.data[id]
	if !exists {
		return errors.New("employee not found")
	}
	e.salary = new_salary
	fmt.Println("updating employee", id, e.name, e.salary)
	return nil
}

// -------------------------------------------------
type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Println("inserting employee", e.name)
	return nil
}

func (s *dumbStorage) get(id int) (employee, error) {
	fmt.Println("getting employee", id)
	e := employee{
		id: id,
	}
	return e, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Println("deleting employee", id)
	return nil
}

func (s *dumbStorage) update(e employee, salary int) error {
	fmt.Println("updating employee", e.name, e.salary)
	return nil
}

func PrintType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}

func SpawnEmployee(s storage) {
	for i := 1; i < 18; i++ {
		err := s.insert(employee{id: i})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//
//var s storage
//fmt.Println(s == nil)
//s = newMemoryStorage()
//fmt.Println(s == nil)
//fmt.Printf("%+v\n", s, "%T %v\n", s, s)
//ds := newDumbStorage()
//fmt.Println(ds == nil)
//fmt.Printf("%+v\n", ds, "%T %v\n", ds, ds)
//
//SpawnEmployee(s)
//
//for i := 1; i <= 10; i++ {
//fmt.Println(s.get(i))
//}
//
//PrintType(3)
//PrintType(6.7)
//PrintType("hjkl")

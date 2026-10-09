package main

import "fmt"

type Vehicle interface {
	move(name string) error
}

type Car struct {
	name string
}

type Aircraft struct {
	name string
}

func (c Car) move(name string) error {
	fmt.Println("car move ", name)
	return nil
}

func (a Aircraft) move(name string) error {
	fmt.Println("aircraft move ", name)
	return nil
}

func drive(v Vehicle, name string) error {
	if err := v.move(name); err != nil {
		return err
	}
	return nil
}

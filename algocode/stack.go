package main

type Stack struct {
	data         []int
	max_value    int
	len_of_stack int
}

func StackBuilder() Stack {
	stack := Stack{}
	stack.data = make([]int, 0)
	stack.len_of_stack = 0
	return stack
}

func (s Stack) Push(v int) {
	s.data = append(s.data, v)
	s.len_of_stack++
	if v > s.max_value {
		s.max_value = v
	}
}

func (s Stack) Pop() int {
	if s.len_of_stack == 0 {
		return -1
	}
	s.len_of_stack--
	return s.data[s.len_of_stack-1]
}

func (s Stack) Len() int {
	return s.len_of_stack
}

func (s Stack) Values() []int {
	return s.data
}

func (s Stack) Max() int {
	return s.max_value
}

//stack := Stack{}
//fmt.Println(stack, stack.data)
//stack.Push(5)
//stack.Push(12)
//stack.Push(1)
//stack.Push(3)
//fmt.Println(stack)
//fmt.Println("Len", stack.Len())
//fmt.Println("Max", stack.Max())
//fmt.Println()
//fmt.Println("Pop", stack.Pop())
//fmt.Println("Pop", stack.Pop())
//
//fmt.Println("Len", stack.Len())
//fmt.Println("Max", stack.Max())

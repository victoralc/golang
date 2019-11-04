package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) IsEmpty() bool {
	return len(stack.values) == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	last := stack.values[stack.Size() - 1]
	stack.values = stack.values[:stack.Size() - 1]
	return last, nil
}

func main() {
	stack := Stack{}

	fmt.Println("Stack created with size: ", stack.Size())
	fmt.Println("Is Empty? ", stack.IsEmpty())

	stack.Push("Go Language")
	stack.Push(2019)
	stack.Push(3.14)
	stack.Push("Software Engineer")
	stack.Push("Google")
	stack.Push(true)

	fmt.Println("Stack size after add elements: ", stack.Size())
	fmt.Println("Is Empty? ", stack.IsEmpty())

	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		fmt.Println("Pop element: ", v)
		fmt.Println("Stack size: ", stack.Size())
		fmt.Println("Is Empty? ", stack.IsEmpty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

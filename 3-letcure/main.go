package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	val  int
	Next *Node
}

func main() {
	s := New()
	s.Push(2)
	s.Push(3)
	s.PrintReverse()
	s.Print()
	fmt.Println(s.Pop())
	s.Print()
	fmt.Println(s.Peek())
	s.Print()
}

type Stack struct {
	length int
	arr    []int
}

func New() Stack {
	return Stack{length: 0}
}

func (s *Stack) Pop() int {
	if s.length == 0 {
		panic("Stack is empty")
	}
	s.length--
	return s.arr[s.length]
}

func (s *Stack) Peek() int {
	if s.length == 0 {
		panic("Stack is empty")
	}
	return s.arr[s.length-1]
}

func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
	s.length++
}

func (s *Stack) Clear() {
	s.arr = []int{}
	s.length = 0
}

func (s *Stack) Contains(x int) bool {
	for i := 0; i < s.length; i++ {
		if s.arr[i] == x {
			return true
		}
	}
	return false
}

func (s *Stack) Increment() {
	for i := 0; i < s.length; i++ {
		s.arr[i]++
	}
}
func (s *Stack) Print() {
	fmt.Print("[ ")
	for i := 0; i < s.length; i++ {
		fmt.Print(strconv.Itoa(s.arr[i]) + " ")
	}
	fmt.Print("]")
	fmt.Println()
}

func (s *Stack) PrintReverse() {
	fmt.Print("[ ")
	for i := s.length - 1; i >= 0; i-- {
		fmt.Print(strconv.Itoa(s.arr[i]) + " ")
	}
	fmt.Print("]")
	fmt.Println()
}

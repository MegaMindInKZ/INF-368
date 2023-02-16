package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	val  int
	Next *Node
}

func NewNode(val int) Node {
	return Node{val: val}
}
func AddVal(n *Node, val int) *Node {
	if n == nil {
		t := NewNode(val)
		n = &t
		return n
	}
	n.Next = AddVal(n.Next, val)
	return n
}

func main() {
	n := NewNode(2)
	n = *AddVal(&n, 3)
	n = *AddVal(&n, 4)
	fmt.Println(n.Next.Next.val)
	s := NewStack()
	s.Push(2)
	s.Push(3)
	s.PrintReverse()
	s.Print()
	fmt.Println(s.Pop())
	s.Print()
	fmt.Println(s.Peek())
	s.Print()
	s.Push(3)
	s.Increment()
	s.Print()
}

type Stack struct {
	length int
	arr    []int
}

func NewStack() Stack {
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

package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	l.InsertAfter(6, l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	l.MoveToFront(l.Back())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	l2 := list.New()
	l2.PushBackList(l)

	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	l.Init()
	fmt.Println(l.Len())

	l3 := list.New()
	for i := 0; i < 5; i++ {
		l3.PushBack(i)
	}

	for e := l3.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	for e := l3.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == 2 {
			l3.Remove(e)
		}
	}
	fmt.Println()

	for e := l3.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
}

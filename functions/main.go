package main

import "fmt"

// this shit is cool, very "dynamic-statically" typed languaged?
type Fn[T any] func(arg T) T

func (r Fn[T]) Function(x T) {
	fmt.Println(x)
}

func main() {
	var X Fn[any] = func(arg any) any {
		fmt.Println(arg)
		return arg
	}

	X("Hello")
	X.Function("World")
}

// wtf even is this idk
// type Instance[T any] interface{ Subset(arg T) }
// type Structure[T any] struct{}
//
// func (x Structure[T]) Subset(arg T) {
// 	fmt.Println(arg)
// }
//
// func main() {
// 	var X Instance[string] = Structure[string]{}
// 	X.Subset("hello")
//
// }

/// pointers r coo'l
// type Box struct {
// 	Value any
// }

// func Add(x any, v int) {
// 	var a *Box
// 	if i, o := x.(*Box); o {
// 		a = *&i
// 	} else {
// 		return
// 	}
//
// 	var b int
// 	if i, o := a.Value.(int); o {
// 		b = *&i
// 	} else {
// 		return
// 	}
//
// 	b += v
// 	a.Value = b
// }
//
// func main() {
// 	x := Box{Value: 42}
// 	Add(&x, 1)
// 	fmt.Println(x.Value) // 43
// }

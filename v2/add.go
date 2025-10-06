// Package learninggoch10exercise provides an addition utility,
// made as an exercise on versioning in Go
//
// The book for which this exercise was done:
// [Learning Go by Jon Bodner]: https://www.oreilly.com/library/view/learning-go-2nd/9781098139285
package learninggoch10exercise

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two [Number]s and returns the sum
//
// More on addition: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T { return a + b }

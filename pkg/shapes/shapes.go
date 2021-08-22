// Package shapes provides general 2D geometrical
// forms and functions to calculate with them.
package shapes

// Shape specifies a two dimensional geometrical object
// and methods to calculate their carateristics.
type Shape interface {
	// Area returns the amount of covered two dimensional
	// space by the shape.
	Area() float64
	// Perimeter returns the sum of the path which the
	// given shape surrounds.
	Perimeter() float64
}

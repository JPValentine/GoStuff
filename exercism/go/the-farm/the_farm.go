package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
var ErrNegative = errors.New("negative fodder")
var ErrDivZero = errors.New("division by zero")

func SillyNephewError(cows int) error {
	x := errors.New(fmt.Sprintf("silly nephew, there cannot be %v cows", cows))
	return x
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	floatAmount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && floatAmount >= 0.0 {
		out := (floatAmount * 2) / float64(cows)
		return out, nil
	} else if (err == ErrScaleMalfunction || err == nil) && floatAmount < 0.0 {
		return 0.0, ErrNegative
	} else if cows == 0 {
		return 0.0, ErrDivZero
	} else if cows < 0 {
		return 0.0, SillyNephewError(cows)
	} else if err != nil {
		return 0.0, err
	} 
	return floatAmount / float64(cows), err
}

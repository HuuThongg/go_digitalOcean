package main

import (
	"fmt"
	"errors"
)

func runValidation(number int) error {
	err := validateValue(number)
	if err != nil {
		return fmt.Errorf("run error: %w", err)
	}
	return nil
}


func  validateValue(number int) error {
	if number == 1 {
		return fmt.Errorf("That's odd")
	} else if number == 2 {
		return errUhOh1
	}
	return nil
}
	// sentinel error.
	var errUhOh = fmt.Errorf("uh h0")
	func giveMeError() error {
		return errUhOh
	}

	var( errUhOh1 = errors.New("Uh oh"))

	// Custom wrapped Errors

	type ValueError struct {
		Value int
		Err   error
	}

	func newValueError(value int, err error) *ValueError {
		return &ValueError{
			Value: value,
			Err:   err,
		}
	}
	func (ve* ValueError) Error() string {
		return fmt.Sprintf("value error: %s", ve.Err)
	}
	func validateValue1 (number int) error {
		if number == 1 {
			return newValueError(number,fmt.Errorf("That's odd"))
		} else if number == 2 {
			return newValueError(number, errUhOh)
		}
		return nil
	}
	func runValidation1(number int) error {
		err := validateValue1(number)
		if err != nil {
			return fmt.Errorf("run error: %w", err)
		}
		return nil
	}
func main() {
	for num :=1; num <= 3; num++ {
		fmt.Printf("Validating %d...", num)
		err := runValidation(num)
		if err == errUhOh1 || errors.Unwrap(err) == errUhOh1 {
			fmt.Println("oh no!")
		} else if err != nil {
			fmt.Println("There was an error", err)
		} else {
			fmt.Println("valid")
		}
	}

	err := giveMeError()
	if err == errUhOh {
		// "uh oh" error code
	}

	// Custom wrapped Errors
	for num :=1; num <= 3; num++ {
		fmt.Printf("Validating %d...", num)
		err := runValidation1(num)
		if err == errUhOh1 || errors.Unwrap(err) == errUhOh ||
		errors.Unwrap(errors.Unwrap(err)) == errUhOh{
			fmt.Println("oh no!")
		} else if err != nil {
			fmt.Println("There was an error", err)
		} else {
			fmt.Println("valid")
		}
	}
}
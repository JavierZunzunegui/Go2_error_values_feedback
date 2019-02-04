// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import "github.com/JavierZunzunegui/Go2_error_values_feedback/xerrors"

// compile time check the provided type implements error
var _ error = MyError{}

func isMyError(err error) bool {
	_, ok := err.(MyError)
	return ok
}

func LastMyError(err error) (MyError, bool) {
	lastErr := xerrors.Last(err, isMyError)
	if lastErr == nil {
		return MyError{}, false
	}
	return lastErr.(MyError), true
}
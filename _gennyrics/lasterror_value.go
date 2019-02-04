package _gennyrics

import (
	"github.com/cheekybits/genny/generic"
	"github.com/JavierZunzunegui/Go2_error_values_feedback/xerrors"
)

type MyErrorType generic.Type

// compile time check the provided type implements error
var _ error = MyErrorType{}

func isMyErrorType(err error) bool {
	_, ok := err.(MyErrorType)
	return ok
}

func LastMyErrorType(err error) (MyErrorType, bool) {
	lastErr := xerrors.Last(err, isMyErrorType)
	if lastErr == nil {
		return MyErrorType{}, false
	}
	return lastErr.(MyErrorType), true
}
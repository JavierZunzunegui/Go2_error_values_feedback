package _gennyrics

import (
	"github.com/cheekybits/genny/generic"
	"github.com/JavierZunzunegui/Go2_error_values_feedback/xerrors"
)

type MyErrorName generic.Type
type MyErrorType generic.Type

// compile time check the provided type implements error
var _ error = MyErrorType(nil)

func isMyErrorName(err error) bool {
	_, ok := err.(MyErrorType)
	return ok
}

func LastMyErrorName(err error) MyErrorType {
	lastErr := xerrors.Last(err, isMyErrorName)
	if lastErr == nil {
		return nil
	}
	return lastErr.(MyErrorType)
}
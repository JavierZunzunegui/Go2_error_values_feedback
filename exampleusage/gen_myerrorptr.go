// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import "github.com/JavierZunzunegui/Go2_error_values_feedback/xerrors"

// compile time check the provided type implements error
var _ error = (*MyErrorPtr)(nil)

func isMyErrorPtr(err error) bool {
	_, ok := err.(*MyErrorPtr)
	return ok
}

func LastMyErrorPtr(err error) *MyErrorPtr {
	lastErr := xerrors.Last(err, isMyErrorPtr)
	if lastErr == nil {
		return nil
	}
	return lastErr.(*MyErrorPtr)
}

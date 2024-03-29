// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import "github.com/JavierZunzunegui/Go2_error_values_feedback/xerrors"

// compile time check the provided type implements error
var _ error = (Barer)(nil)

func isBarer(err error) bool {
	_, ok := err.(Barer)
	return ok
}

func LastBarer(err error) Barer {
	lastErr := xerrors.Last(err, isBarer)
	if lastErr == nil {
		return nil
	}
	return lastErr.(Barer)
}

# This repo is for use in the discussions at https://github.com/golang/go/issues/29934 only.

It shows hot to use [go generate](https://blog.golang.org/generate) to make typed functions for the proposed `Last(error, func(error) bool) error`.
It uses [genny](https://github.com/cheekybits/genny).

```go
type MyError struct {/* */}
func (*MyError) Error() string {/* */}

//go:generate genny ...
// and this produces (in a different file but same package):
func LastMyError(err error) *MyError {/* */}
// note the return type is not error by *MyError
```

It has 3 packages:
- xerrros: represents the proposed changed to the standard library
- _gennyrics: the files I use alongside genny to generate the types code
- exampleusage: a main, where you can validate everything works as expected

The output of running the main is:
```
*MyErrorPtr to MyErrorPtr: pointer implements error - I support foo!
*MyErrorPtr to Foorer: pointer implements error - I support foo!
MyError to MyError: value implements error - I support bar!
MyError to Barer: value implements error - I support bar!
MyError{*MyErrorPtr} to MyErrorPtr: pointer implements error - I support foo!
MyError{*MyErrorPtr} to MyError: value implements error - I support bar!
MyError{*MyErrorPtr} to Foorer: pointer implements error - I support foo!
MyError{*MyErrorPtr} to Barer: value implements error - I support bar!
*MyErrorPtr{MyError} to MyErrorPtr: pointer implements error - I support foo!
*MyErrorPtr{MyError} to MyError: value implements error - I support bar!
*MyErrorPtr{MyError} to Foorer: pointer implements error - I support foo!
*MyErrorPtr{MyError} to Barer: value implements error - I support bar!
```
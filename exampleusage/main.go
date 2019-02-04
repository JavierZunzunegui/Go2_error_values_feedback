package main

import "fmt"

//go:generate genny -in=../_gennyrics/lasterror.go -out=gen_myerrorptr.go -pkg=main gen "MyErrorName=MyErrorPtr MyErrorType=*MyErrorPtr"
type MyErrorPtr struct {
	wrapped error
}

func (err *MyErrorPtr) Error() string { return "pointer implements error" }

func (err *MyErrorPtr) Wrapped() error { return err.wrapped }

func (err *MyErrorPtr) Foo() string { return "I support foo!" }

//go:generate genny -in=../_gennyrics/lasterror_value.go -out=gen_myerror.go -pkg=main gen "MyErrorType=MyError"
type MyError struct {
	wrapped error
}

func (err MyError) Error() string { return "value implements error" }

func (err MyError) Wrapped() error { return err.wrapped }

func (err MyError) Bar() string { return "I support bar!" }

//go:generate genny -in=../_gennyrics/lasterror.go -out=gen_foorer.go -pkg=main gen "MyErrorName=Foorer MyErrorType=Foorer"
type Foorer interface {
	error
	Foo() string
}

//go:generate genny -in=../_gennyrics/lasterror.go -out=gen_barer.go -pkg=main gen "MyErrorName=Barer MyErrorType=Barer"
type Barer interface {
	error
	Bar() string
}

func check(label string, err error) {
	if e := LastMyErrorPtr(err); e != nil {
		fmt.Println(label, "to MyErrorPtr:", e.Error(), "-", e.Foo())
	}
	if e, ok := LastMyError(err); ok {
		fmt.Println(label, "to MyError:", e.Error(), "-", e.Bar())
	}
	if e := LastFoorer(err); e != nil {
		fmt.Println(label, "to Foorer:", e.Error(), "-", e.Foo())
	}
	if e := LastBarer(err); e != nil {
		fmt.Println(label, "to Barer:", e.Error(), "-", e.Bar())
	}
}

func main() {
	/*
		Output:
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
	*/
	check("nil", nil)
	check("*MyErrorPtr", &MyErrorPtr{})
	check("MyError", MyError{})
	check("MyError{*MyErrorPtr}", MyError{&MyErrorPtr{}})
	check("*MyErrorPtr{MyError}", &MyErrorPtr{MyError{}})
}

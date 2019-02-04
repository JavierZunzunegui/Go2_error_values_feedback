package main

//go:generate genny -in=../_gennyrics/lasterror.go -out=gen_myerrorptr.go -pkg=main gen "MyErrorName=MyErrorPtr MyErrorType=*MyErrorPtr"
type MyErrorPtr struct {
	wrapped error
}

func (err *MyErrorPtr) Error() string {return "pointer implements error"}

func (err *MyErrorPtr) Wrapped() error {return err.wrapped}

func (err *MyErrorPtr) Foo() {}

//ignore go:generate genny -in=../_gennyrics/lasterror.go -out=gen_myerror.go -pkg=main gen "MyErrorName=MyError MyErrorType=*MyError"
type MyError struct {
	wrapped error
}

func (err MyError) Error() string {return "value implements error"}

func (err MyError) Wrapped() error {return err.wrapped}

func (err MyError) Bar() {}

type Foorer interface {
	Foo()
}

type Barer interface {
	Bar()
}

func main() {
	//go:generate genny -in=../_gennyrics/map.go -out=gen_map_string_string.go -pkg=collections gen "Key=string Value=string"

}

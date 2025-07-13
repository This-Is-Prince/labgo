package errorunderstanding

import (
	"errors"
	"fmt"
	"io"
)

type First struct {
	MsgFirst string
	Err1     error
}

func (f *First) Error() string {
	return fmt.Sprintf("First error: %s, Error 1: %v", f.MsgFirst, f.Err1)
}

func (f *First) Unwrap() error {
	return f.Err1
}

type Second struct {
	MsgSecond string
	Err2      error
}

func (s *Second) Error() string {
	return fmt.Sprintf("Second error: %s, Error 2: %v", s.MsgSecond, s.Err2)
}

func (s *Second) Unwrap() error {
	return s.Err2
}

type Third struct {
	MsgThird string
	Err3     error
}

func (t *Third) Error() string {
	return fmt.Sprintf("Third error: %s, Error 3: %v", t.MsgThird, t.Err3)
}

func (t *Third) Unwrap() error {
	return t.Err3
}

var (
	Error3 = &Third{
		MsgThird: "This is the third error",
		Err3:     io.EOF,
	}
	Error2 = &Second{
		MsgSecond: "This is the second error",
		Err2:      Error3,
	}
	Error1 = &First{
		MsgFirst: "This is the first error",
		Err1:     Error2,
	}
	Error4 = &First{
		MsgFirst: "This is the first error with a nil second error",
		Err1:     nil,
	}
	Error5 = errors.New("This is a simple error")
)

func check() error {
	return Error1
}

type ErrorStr string

func (e *ErrorStr) Error() string {
	return string(*e)
}

func Learn(run bool) {
	if !run {
		return
	}

	var err error
	fmt.Println(err)
	fmt.Printf("%T\n\n", err)
	err = errors.New("this is an error")
	fmt.Println(err)
	fmt.Printf("%T\n\n", err)

	errStr := ErrorStr("This is a custom error")
	err = &errStr
	fmt.Println(err)
	fmt.Printf("%T\n\n", err)

	switch errStr := err.(type) {
	case *ErrorStr:
		fmt.Println(*errStr)
		fmt.Printf("%T\n\n", errStr)
		fmt.Println("Error type:", errStr)
		fmt.Println("This is a custom error type")
	case error:
		fmt.Println("This is a standard error type")
	}

	errStruct := check()
	fmt.Println(errStruct)
	fmt.Printf("%T\n\n", errStruct)

	isEofError := errors.Is(errStruct, Error5)
	fmt.Println("Is EOF error:", isEofError)

	var errCustom error

	isFound := errors.As(errStruct, &errCustom)
	fmt.Println("Is First error found:", isFound)
	fmt.Println(errCustom)
}

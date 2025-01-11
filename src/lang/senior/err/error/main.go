package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// 不需要格式化字符串 尽量用errors.New()
	e1 := errors.New("OK")
	fmt.Println(e1)

	e2 := fmt.Errorf("封装 %T", e1)
	fmt.Println(e2)

	var e3 = errors.New("permission denied")

	if errors.Is(e3, os.ErrPermission) {
		fmt.Println(e3)
	}

	exampleAssertError()

	exampleCreateBasicError()

	exampleCreateWrapError()

	exampleUnwrap()

	exampleAssertChanWithAs()
}

func exampleCreateBasicError() {
	err := errors.New("this is demo error")
	basicErr := fmt.Errorf("some context: %v", err)
	if _, ok := basicErr.(interface{ Unwrap() error }); !ok {
		fmt.Println("basicErr is a errorString")
	}
}

func exampleCreateWrapError() {
	err := errors.New("this is demo error")
	wrapErr := fmt.Errorf("some context: %w", err)
	if _, ok := wrapErr.(interface{ Unwrap() error }); ok {
		fmt.Println("wrapError is a wrapError")
	}
}

func exampleAssertError() {
	e1 := &os.PathError{
		Op:   "write",
		Path: "/root/demo.text",
		Err:  os.ErrPermission,
	}

	assertError(e1)

	e2 := fmt.Errorf("not an os.PathError")

	assertError(e2)
}

func assertError(err error) {
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("it's an os.PathError, operation: %s, path : %s, msg: %v \n", e.Op, e.Path, e.Err)
	}
}

func exampleUnwrap() {
	err := fmt.Errorf("write file error: %w", os.ErrPermission)
	if errors.Unwrap(err) == os.ErrPermission {
		fmt.Println("permission denied")
	}
}

func exampleAssertChanWithAs() {
	e1 := &os.PathError{
		Op:   "write",
		Path: "/root/demo.text",
		Err:  os.ErrPermission,
	}

	e2 := fmt.Errorf("some context: %w", e1)

	var target *os.PathError
	if errors.As(e2, &target) {
		fmt.Printf("it's an os.PathError, operation: %s, path : %s, msg: %v \n", target.Op, target.Path, target.Err)
	}

}

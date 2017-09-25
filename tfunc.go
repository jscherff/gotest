// See article https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
// Code from https://github.com/benbjohnson/testing

package gotest

import (
	"fmt"
	"path/filepath"
	"runtime"
	"reflect"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
	}
}

// NotEquals fails the test if exp is not equal to act.
func NotEquals(tb testing.TB, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d:\n\n\texp: NOT %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
	}
}


// HereInfo prints useful information about an execution point
func HereInfo() {

	// Need space only for PC of caller.
	pc := make([]uintptr, 1)

	// Skip PC of Callers and decorator function.
	n := runtime.Callers(2, pc)

	// Return if no PCs available.
	if n == 0 { return }

	// Obtain the caller frame.
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	// Print file and line number, then return.
	fmt.Printf("%s:%d\n", frame.File, frame.Line)
}

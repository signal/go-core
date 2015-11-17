// assert provides helpful assertion methods for any kind of test case.
package assert

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func errorf(t *testing.T, skip int, msgf string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		file, line = "???", 0
	} else if wd, err := os.Getwd(); err == nil {
		file = strings.SplitN(file, wd, 2)[1][1:]
	}
	errArgs := []interface{}{file, line}
	t.Errorf("[%s:%d]\n\r"+msgf+"\n\n", append(errArgs, args...)...)
}

// True expects actual to be true, else call t.Error
func True(t *testing.T, what string, actual bool) {
	isTrue(t, what, actual)
}

// False expects actual to be false, else call t.Error
func False(t *testing.T, what string, actual bool) {
	if actual {
		errorf(t, 1, "For %s expected %v to be equal %v", what, actual, false)
	}
}

// Equal is basic equality checking. If the two provided types are not equal, call t.Error.
func Equal(t *testing.T, what string, expected interface{}, actual interface{}) {
	isEqual(t, what, expected, actual)
}

// DeepEqual does equality checking using reflect.DeepEqual. If the two provided types are not equal, call t.Error.
func DeepEqual(t *testing.T, what string, expected, actual interface{}) {
	isDeepEqual(t, what, expected, actual)
}

// ZeroValue asserts that actual is the zero-value for its type.
func ZeroValue(t *testing.T, what string, actual interface{}) {
	zeroValue := reflect.Zero(reflect.ValueOf(actual).Type()).Interface()
	isDeepEqual(t, "zero value of "+what, zeroValue, actual)
}

// NotEqual performs inequality checking. If the two provided types are equal, call t.Error.
func NotEqual(t *testing.T, what string, expected interface{}, actual interface{}) {
	if expected == actual {
		errorf(t, 1, "For %s expected %v NOT to equal %v", what, actual, expected)
	}
}

// NotDeepEqual performs inequality checking using reflect.DeepEqual. If the two provided types are not equal,
// call t.Error.
func NotDeepEqual(t *testing.T, what string, expected interface{}, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		errorf(t, 1, "For %s expected %v NOT to equal %v", what, actual, expected)
	}
}

// Nil asserts that the provided value is nil, else call t.Error.
func Nil(t *testing.T, what string, actual interface{}) {
	if actual != nil {
		errorf(t, 1, "For %s expected %v to be %v", what, actual, nil)
	}
}

// NotNil asserts that the provided value is not nil, else call t.Error.
func NotNil(t *testing.T, what string, actual interface{}) {
	if actual == nil {
		errorf(t, 1, "For %s expected %v NOT to be %v", what, actual, nil)
	}
}

// NoError ensures no value for error was provided, else call t.Error.
func NoError(t *testing.T, err error) {
	if err != nil {
		errorf(t, 1, "Expected no error, but received %q", err)
	}
}

// Error asserts that an error was provided and it's message.
func Error(t *testing.T, expectedWhat string, err error) {
	if err == nil {
		errorf(t, 1, "Expected error with message %q but no error received",
			expectedWhat)
	} else {
		isEqual(t, "error", expectedWhat, err.Error())
	}
}

// Len asserts that a slice/array/string has an expected length.
func Len(t *testing.T, what string, expectedLen int, actual interface{}) {
	rv := reflect.ValueOf(actual)
	isEqual(t, "length of "+what, expectedLen, reflectLen(t, rv))
}

func reflectLen(t *testing.T, rv reflect.Value) int {
	kind := rv.Kind()
	if kind == reflect.Array || kind == reflect.Slice || kind == reflect.String {
		return rv.Len()
	} else {
		errorf(t, 2, "len is not supported: %v", rv)
		t.FailNow()
	}
	return 0
}

// SameElements asserts that two unordered slices/arrays/strings contain the same elements.
func SameElements(t *testing.T, what string, expected interface{}, actual interface{}) {
	rvExpected := reflect.ValueOf(expected)
	rvActual := reflect.ValueOf(actual)
	lenExpected := reflectLen(t, rvExpected)
	isEqual(t, "element lengths", lenExpected, reflectLen(t, rvActual))
	for i := 0; i < lenExpected; i++ {
		elem := rvExpected.Index(i).Interface()
		isTrue(t, fmt.Sprintf("contains %v", elem), contains(elem, rvActual))
	}
}

func contains(expected interface{}, actualSlice reflect.Value) bool {
	lenActual := actualSlice.Len()
	for i := 0; i < lenActual; i++ {
		actual := actualSlice.Index(i).Interface()
		if expected == actual {
			return true
		}
	}
	return false
}

func isTrue(t *testing.T, what string, actual bool) {
	if !actual {
		errorf(t, 2, "For %s expected %v to be equal %v", what, actual, true)
	}
}

func isEqual(t *testing.T, what string, expected interface{}, actual interface{}) {
	if expected != actual {
		errorf(t, 2, "For %s expected %v to equal %v", what, actual, expected)
	}
}

func isDeepEqual(t *testing.T, what string, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		errorf(t, 2, "For %s expected %v to equal %v", what, actual, expected)
	}
}

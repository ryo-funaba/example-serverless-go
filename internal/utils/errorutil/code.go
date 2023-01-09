package errorutil

import (
	"fmt"

	"github.com/pkg/errors"
)

type customError struct {
	code Code
	err  error
}

func ErrorCode(err error) int {
	if err == nil {
		return OK.Number()
	}

	var e customError

	ok := errors.As(err, &e)
	if ok {
		return e.code.Number()
	}

	return Unknown.Number()
}

func Errorf(code Code, format string, args ...interface{}) error {
	if code == OK {
		return nil
	}

	err := customError{
		code: code,
		err:  errors.Errorf(format, args...),
	}

	err.PrintCodeAndMsg()

	return err
}

func (e customError) Error() string {
	return fmt.Sprintf("%s", e.err)
}

func (e customError) PrintCodeAndMsg() {
	fmt.Printf("Code: %d, Msg: %s\n", e.code.Number(), e.Error())
}

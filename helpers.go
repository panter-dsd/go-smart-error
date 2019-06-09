package smarterror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var projectRootPath string

// InitProjectRootPath ...
func InitProjectRootPath(pathToSkip string) {
	_, p, _, _ := runtime.Caller(1)
	index := strings.Index(p, pathToSkip)
	if index >= 0 {
		projectRootPath = p[:index]
	}
}

// FullError ...
func FullError(err error) string {
	if e, ok := err.(IError); ok {
		return e.FullError()
	}
	return err.Error()
}

// Wrap ...
func Wrap(err error) IError {
	return Wrapf(err, "")
}

// Wrapf ...
func Wrapf(err error, message string, a ...interface{}) IError {
	if err == nil {
		return nil
	}
	var (
		pc       uintptr
		fileName string
		line     int
	)
	for i := 1; ; i++ {
		var ok bool
		pc, fileName, line, ok = runtime.Caller(i)
		if !ok || !strings.Contains(fileName, "go-smart-error") {
			break
		}
	}
	return &Error{
		msg: func() string {
			if len(message) == 0 {
				return ""
			}
			return fmt.Sprintf(message, a...)
		}(),
		err:        err,
		shortError: false,
		file:       fileName,
		line:       line,
		funcName:   runtime.FuncForPC(pc).Name(),
	}
}

// New ...
func New(message string, a ...interface{}) IError {
	return Wrapf(errors.New(""), message, a...)
}

package smarterror

import (
	"fmt"
	"strings"
)

// Error ...
type Error struct {
	msg        string
	err        error
	shortError bool

	file     string
	line     int
	funcName string
}

func (b *Error) Error() string {
	if len(b.msg) == 0 {
		return b.err.Error()
	}
	if b.shortError {
		return b.msg
	}
	return fmt.Sprintf("%s: %v", b.msg, b.err)
}

// FullError ...
func (b *Error) FullError() string {
	if len(b.msg) == 0 {
		return fmt.Sprintf(`%v (%s:%d)`, b.err, b.fileName(), b.line)
	}
	if b.shortError {
		return fmt.Sprintf(`%s (%s:%d)`, b.msg, b.fileName(), b.line)
	}
	return fmt.Sprintf(`%s: %v (%s:%d)`, b.msg, b.err, b.fileName(), b.line)
}

// Cause ...
func (b *Error) Cause() error {
	return b.err
}

// Message ...
func (b *Error) Message() string {
	return b.msg
}

func (b *Error) fileName() string {
	return strings.Replace(b.file, projectRootPath, "", 1)
}

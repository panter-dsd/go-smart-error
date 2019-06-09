package smarterror

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmartErrorTestNew(t *testing.T) {
	err := New("some")
	errStr := FullError(err)
	assert.True(t, strings.Contains(errStr, "some:  (/usr/lib64/go/src/testing/testing.go"), errStr)
}

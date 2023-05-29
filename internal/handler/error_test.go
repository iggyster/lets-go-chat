package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors_AddError(t *testing.T) {
	errs := &Errors{}

	errs.AddError("field", "message", "detail")

	assert.Equal(t, 1, errs.Count())
}
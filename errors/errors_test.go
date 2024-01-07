package errors

import (
	stderrors "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtend(t *testing.T) {
	base := stderrors.New("base")
	extended := Extend(base, "error")
	assert.Equal(t, "base: error", extended.Error())
}

func TestExtendWithSubs(t *testing.T) {
	base := stderrors.New("base")
	sub1 := stderrors.New("sub1")
	sub2 := stderrors.New("sub2")
	extended := Extend(base, "error", sub1, sub2)
	assert.Equal(t, "base: error: sub1: sub2", extended.Error())
}

func TestExtendf(t *testing.T) {
	base := stderrors.New("base")
	extended := Extendf(base, "error with value %d", 42)
	assert.Equal(t, "base: error with value 42", extended.Error())
}

func TestExtendfWithSubs(t *testing.T) {
	base := stderrors.New("base error")
	sub1 := stderrors.New("sub1")
	sub2 := stderrors.New("sub2")
	extended := Extendf(base, "error with value %d", sub1, 42, sub2)
	assert.Equal(t, "base error: error with value 42: sub1: sub2", extended.Error())
}

func TestExtends(t *testing.T) {
	base := stderrors.New("base error")
	sub := stderrors.New("sub error")
	extended := Extend(base, "error", sub)
	assert.True(t, Extends(extended, base))
	assert.False(t, Extends(extended, sub))

	extended2 := Extend(extended, "error")
	assert.True(t, Extends(extended2, extended))
}

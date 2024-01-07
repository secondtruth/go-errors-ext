package errors

import (
	"fmt"
)

// Extend returns a new error that extends the base error and optionally wraps given sub errors.
func Extend(base error, msg string, subs ...error) error {
	if base == nil {
		panic("base error must not be nil")
	}
	nsubs := 0
	for _, err := range subs {
		if err != nil {
			nsubs++
		}
	}
	e := &extendedError{
		base: base,
		msg:  msg,
	}
	if nsubs > 0 {
		e.subs = make([]error, 0, nsubs)
		for _, err := range subs {
			if err != nil {
				e.subs = append(e.subs, err)
			}
		}
	}
	return e
}

// Extendf returns a new formatted error that extends the base error and optionally wraps given sub errors.
// args can be both arbitrary values or errors to be wrapped.
func Extendf(base error, format string, args ...interface{}) error {
	var subs []error
	var fmtArgs []any
	for _, arg := range args {
		if err, ok := arg.(error); ok {
			subs = append(subs, err)
		} else {
			fmtArgs = append(fmtArgs, arg)
		}
	}
	msg := fmt.Sprintf(format, fmtArgs...)
	return Extend(base, msg, subs...)
}

// Extends reports whether any error in err's tree extends target.
func Extends(err, target error) bool {
	if e, ok := err.(interface{ Base() error }); ok {
		if e.Base() == target {
			return true
		}
		return Extends(e.Base(), target)
	}
	return false
}

// ExtendedError is an error that extends another error. It implements the Base() method.
type ExtendedError interface {
	Base() error
}

// ChainedErrors is an error that wraps other errors. It implements the Subs() method.
type ChainedErrors interface {
	Subs() []error
}

type extendedError struct {
	base error
	msg  string
	subs []error
}

func (e *extendedError) Error() string {
	var b []byte
	b = append(b, e.base.Error()+": "+e.msg...)
	for _, err := range e.subs {
		b = append(b, ": "+err.Error()...)
	}
	return string(b)
}

// Base returns the base error.
func (e *extendedError) Base() error {
	return e.base
}

// Subs returns the sub errors.
func (e *extendedError) Subs() []error {
	return e.subs
}

func (e *extendedError) Unwrap() []error {
	return append([]error{e.base}, e.subs...)
}

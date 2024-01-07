# Go Errors Extensions

This library provides a set of extensions for error creation and handling in Go.

## Installation

To install the go-errors-ext library, use the following command:

	go get -u github.com/secondtruth/go-errors-ext

## Quick Start

After installation, you can import the library into your Go project:

```go
import errorsx "github.com/secondtruth/go-errors-ext/errors"
```

> [!NOTE]
> Renaming the package to `errorsx` is not always necessary, but is recommended to do it always
> to make a clear distinction between the standard `errors` package and this one.

## Usage

### Extending errors

The `Extend` and `Extendf` functions can be used to extend an existing error:

```go
base := errors.New("base error")

extendedPlain := errorsx.Extend(base, "error")
extendedFormatted := errorsx.Extendf(base, "error with value %d", 42)
```

Sub-errors can also be provided:

```go
base := errors.New("base error")
sub1 := errors.New("suberror 1")
sub2 := errors.New("suberror 2")

extendedPlain := errorsx.Extend(base, "error", sub1, sub2)
extendedFormatted := errorsx.Extendf(base, "error with value %d", sub1, 42, sub2)
```

In `Extendf`, sub-errors can be placed anywhere in the format args.

### Testing errors

The `Extends` function can be used to test if an error extends another error:

```go
base := errors.New("base error")
sub := errors.New("suberror")
extended1 := errorsx.Extend(base, "error", sub)
extended2 := errorsx.Extend(extended1, "error")

extends1 := errorsx.Extends(extended1, base) // true
extends2 := errorsx.Extends(extended2, extended1) // true
extendsSub := errorsx.Extends(extended1, sub) // false
```

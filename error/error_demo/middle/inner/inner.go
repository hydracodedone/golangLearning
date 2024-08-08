package inner

import "github.com/pkg/errors"

var InnerError error = errors.New("error in inner")

func Inner() error {
	return InnerError
}

func InnerWithStack() error {
	return errors.WithStack(InnerError)
}

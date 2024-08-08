package middle

import (
	"error_demo/middle/inner"
	"fmt"

	"github.com/pkg/errors"
)

func Middle() error {
	err := inner.Inner()
	if err != nil {
		return fmt.Errorf("middle recv:%v", err)
	}
	return err
}
func MiddleWithWrap() error {
	err := inner.Inner()
	if err != nil {
		return errors.Wrap(err, "in middle")
	}
	return err
}

func MiddleWithStack() error {
	err := inner.Inner()
	return errors.WithStack(err)
}

package example

import (
	"errors"
	"fmt"
)

func Test() error {
	fmt.Print("test")

	return errors.New("haha")
}

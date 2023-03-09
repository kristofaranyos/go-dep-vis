package example

import (
	ees "errors"
	"fmt"
)

func Test() error {
	fmt.Print("test")

	return ees.New("haha")
}

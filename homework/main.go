package homework

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := myFunc()
	if err != nil {
		fmt.Sprintf("%T\n%v", errors.Cause(err), errors.Cause(err))
	}
}

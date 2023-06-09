package handler

import (
	"fmt"
)

func ExampleErrors_AddError() {
	errs := &Errors{}

	errs.AddError("field", "message", "detail")

	fmt.Println(errs.Count())
	// Output: 1
}

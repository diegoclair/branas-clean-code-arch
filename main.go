package main

import (
	"fmt"

	"github.com/diegoclair/branas-clean-code-arch/validator"
)

func main() {
	fmt.Println(validator.IsValidDocumentNumber("111.111.111-11"))
	fmt.Println(validator.IsValidDocumentNumber("123.456.789-99"))
	fmt.Println(validator.IsValidDocumentNumber("935.411.347-80"))
}

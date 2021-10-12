package main

import (
	"fmt"

	"github.com/diegoclair/branas-clean-code-arch/utils"
)

func main() {
	fmt.Println(utils.IsValidDocumentNumber("111.111.111-11"))
	fmt.Println(utils.IsValidDocumentNumber("123.456.789-99"))
	fmt.Println(utils.IsValidDocumentNumber("935.411.347-80"))
}

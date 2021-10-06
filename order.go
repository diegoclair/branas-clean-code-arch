package main

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/validator"
)

func MakeOrder(document string) error {

	if !validator.IsValidDocumentNumber(document) {
		return errors.New("invalid document")
	}

	return nil
}

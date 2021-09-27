package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(IsValidDocumentNumber("111.111.111-11"))
	fmt.Println(IsValidDocumentNumber("123.456.789-99"))
	fmt.Println(IsValidDocumentNumber("935.411.347-80"))
}

var notNumberRE = regexp.MustCompile(`\D`)

func IsValidDocumentNumber(document string) bool {

	document = cleanNumber(document)
	isCPF := len(document) == 11
	isCNPJ := len(document) == 14
	if !isCPF && !isCNPJ {
		return false
	}
	if isCPF {
		return validateCPF(document)
	}
	return validateCNPJ(document)
}

func validateCPF(cpf string) bool {

	lenToFirstDigit := len(cpf) - 2
	lenToSecondDigit := len(cpf) - 1

	if invalidEqualNumbers(cpf) {
		return false
	}

	firstDigit, _ := strconv.Atoi(cpf[lenToFirstDigit:lenToSecondDigit])
	calculatedFirstDigit := calculateCPFDigit(cpf[:lenToFirstDigit])
	isCorrect := firstDigit == calculatedFirstDigit
	if !isCorrect {
		return false
	}

	secondDigit, _ := strconv.Atoi(cpf[lenToSecondDigit:])
	calculatedSecondDigit := calculateCPFDigit(cpf[:lenToSecondDigit])
	isCorrect = secondDigit == calculatedSecondDigit
	return isCorrect
}

func calculateCPFDigit(document string) int {
	multiplier := len(document) + 1
	sum := 0
	for i := 0; i < len(document); i++ {
		pos, _ := strconv.Atoi(string(document[i]))
		sum += pos * multiplier
		multiplier--
	}
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func invalidEqualNumbers(document string) bool {
	var allNumbersAreEqual bool
	var digit string
	for _, val := range document {
		if len(digit) == 0 {
			digit = string(val)
		}
		if string(val) == digit {
			allNumbersAreEqual = true
			continue
		}
		allNumbersAreEqual = false
		break
	}
	return allNumbersAreEqual
}

func validateCNPJ(cnpj string) bool {

	//need to be implemented
	return true
}

// cleanNumber remove all not number characters
func cleanNumber(value string) string {
	return notNumberRE.ReplaceAllString(value, "")
}

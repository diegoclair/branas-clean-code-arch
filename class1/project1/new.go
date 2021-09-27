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

	cleanedDocument := cleanNumber(document)
	isCPF := len(cleanedDocument) == 11
	isCNPJ := len(cleanedDocument) == 14
	if !isCPF && !isCNPJ {
		return false
	}
	if isCPF {
		return validateCPF(cleanedDocument)
	}
	return validateCNPJ(cleanedDocument)
}

func validateCPF(cpf string) bool {

	lenToFirstDigit := len(cpf) - 2
	lenToSecondDigit := len(cpf) - 1

	if invalidEqualNumbers(cpf) {
		return false
	}

	firstDigit, _ := strconv.Atoi(cpf[lenToFirstDigit:lenToSecondDigit])
	secondDigit, _ := strconv.Atoi(cpf[lenToSecondDigit:])
	calculatedFirstDigit, calculatedSecondDigit := calculateCPFDigits(cpf)

	return firstDigit == calculatedFirstDigit && secondDigit == calculatedSecondDigit

}

func calculateCPFDigits(document string) (firstDigit, secondDigit int) {
	lenTocalculateFirstDigit, lenTocalculateSecondDigit := 8, 9
	factorForFirstDigit, factorForSecondDigit := 10, 11
	sumFirstDigit, sumSecondDigit := 0, 0
	for i := 0; i < len(document); i++ {
		pos, _ := strconv.Atoi(string(document[i]))
		if i <= lenTocalculateFirstDigit {
			sumFirstDigit += pos * factorForFirstDigit
			factorForFirstDigit--
		}
		if i <= lenTocalculateSecondDigit {
			sumSecondDigit += pos * factorForSecondDigit
			factorForSecondDigit--
		}
	}
	//234 287
	fmt.Println(sumFirstDigit, sumSecondDigit)
	rest := sumFirstDigit % 11
	if rest >= 2 {
		firstDigit = 11 - rest
	}

	rest = sumSecondDigit % 11
	if rest >= 2 {
		secondDigit = 11 - rest
	}
	return
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

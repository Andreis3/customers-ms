package valueobject

import (
	"slices"
	"strings"
	"unicode"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

const (
	CNPJLength            = 14
	CNPJFirstDigitIdx     = 12
	CNPJSecondDigitIdx    = 13
	CNPJModuleDivisor     = 11
	CNPJBlacklistLength   = 10
	CNPJASCIIZero         = '0'
	CNPJFirstDigitWeight  = 5
	CNPJSecondDigitWeight = 6
)

var blackListCNPJ = []string{
	"00000000000000", "11111111111111", "22222222222222", "33333333333333",
	"44444444444444", "55555555555555", "66666666666666", "77777777777777",
	"88888888888888", "99999999999999",
}

type CNPJ struct {
	CNPJ      string
	Validator validator.Validator
}

func NewCNPJ(cnpj string, validator validator.Validator) *CNPJ {
	return &CNPJ{CNPJ: cnpj, Validator: validator}
}

func (c *CNPJ) Validate() {
	cleanedCNPJ := cleanCNPJ(c.CNPJ)
	c.Validator.Assert(validator.NotBlank(cleanedCNPJ), "cnpj", validator.NotBlankField)
	c.Validator.Assert(validateCNPJ(cleanedCNPJ), "cnpj", "cnpj: is invalid")
}

func cleanCNPJ(cnpj string) string {
	var sb strings.Builder
	for _, r := range cnpj {
		if unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func validateCNPJ(cnpj string) bool {
	if len(cnpj) != CNPJLength || slices.Contains(blackListCNPJ, cnpj) {
		return false
	}
	return validateDigitCNPJ(cnpj, CNPJFirstDigitIdx, CNPJFirstDigitWeight) &&
		validateDigitCNPJ(cnpj, CNPJSecondDigitIdx, CNPJSecondDigitWeight)
}

func validateDigitCNPJ(cnpj string, position, startWeight int) bool {
	sum := 0
	for i, char := range cnpj[:position] {
		sum += int(char-CNPJASCIIZero) * (startWeight - i)
	}

	rest := sum % CNPJModuleDivisor
	if rest < 2 {
		rest = 0
	} else {
		rest = CNPJModuleDivisor - rest
	}

	return rest == int(cnpj[position]-CNPJASCIIZero)
}

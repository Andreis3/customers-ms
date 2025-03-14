package valueobject

import (
	"regexp"
	"slices"
	"strconv"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

const (
	CNPJLength   = 14
	ModuleBase   = 11
	MinWeight    = 2
	MaxWeight    = 9
	BlackListMod = 2
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

	// Validações
	c.Validator.CheckField(validator.NotBlank(cleanedCNPJ), "cnpj", validator.NotBlankField)
	c.Validator.CheckField(len(cleanedCNPJ) == CNPJLength, "cnpj", "cnpj: must have 14 characters")
	c.Validator.CheckField(!slices.Contains(blackListCNPJ, cleanedCNPJ), "cnpj", "cnpj: must be a valid CNPJ number")
	c.Validator.CheckField(validateCNPJ(cleanedCNPJ), "cnpj", "cnpj: is invalid, must be a valid CNPJ number calculated with the module 11 algorithm")
}

func cleanCNPJ(cnpj string) string {
	regex := regexp.MustCompile("[^0-9]")
	return regex.ReplaceAllString(cnpj, "")
}

func validateCNPJ(cnpj string) bool {
	if len(cnpj) != CNPJLength {
		return false
	}

	return validateDigitCNPJ(cnpj, CNPJLength-2) &&
		validateDigitCNPJ(cnpj, CNPJLength-1)
}

func validateDigitCNPJ(cnpj string, position int) bool {
	sum := 0
	weight := MinWeight

	for i := position - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weight
		weight++
		if weight > MaxWeight {
			weight = MinWeight
		}
	}

	rest := sum % ModuleBase
	expectedDigit := 0
	if rest >= BlackListMod {
		expectedDigit = ModuleBase - rest
	}

	return strconv.Itoa(expectedDigit) == string(cnpj[position])
}

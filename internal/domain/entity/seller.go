package entity

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

type Seller struct {
	User
	CompanyName   string
	CNPJ          string
	BusinessName  string
	FundationDate time.Time
	Addresses     []Address
	Validator     validator.Validator
}

func SellerBuilder() *Seller {
	return &Seller{}
}

func (s *Seller) SetID(id string) *Seller {
	s.ID = id
	return s
}

func (s *Seller) SetEmail(email string) *Seller {
	s.Email = email
	return s
}

func (s *Seller) SetPassword(password string) *Seller {
	s.Password = password
	return s
}

func (s *Seller) SetCompanyName(CompanyName string) *Seller {
	s.CompanyName = CompanyName
	return s
}

func (s *Seller) SetCNPJ(cnpj string) *Seller {
	s.CNPJ = cnpj
	return s
}

func (s *Seller) SetBusinessName(businessName string) *Seller {
	s.BusinessName = businessName
	return s
}

func (s *Seller) SetFundationDate(fundationDate time.Time) *Seller {
	s.FundationDate = fundationDate
	return s
}

func (s *Seller) SetAddresses(addresses []Address) *Seller {
	s.Addresses = addresses
	return s
}

func (s *Seller) Build() *Seller {
	return s
}

func (s *Seller) Validate() *validator.Validator {
	user := s.User.Validate()

	for key, value := range user {
		s.Validator.AddFieldError(key, value)
	}

	s.Validator.CheckField(validator.NotBlank(s.CompanyName), "company_name", validator.NotBlankField)
	s.Validator.CheckField(validator.NotBlank(s.CNPJ), "cnpj", validator.NotBlankField)
	s.Validator.CheckField(validator.NotBlank(s.BusinessName), "business_name", validator.NotBlankField)

	return &s.Validator
}

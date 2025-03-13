//go:build unit
// +build unit

package entity_test

import (
	"fmt"
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/validator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: ADDRES", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when build new address", func() {
				entity := entity.SellerBuilder().
					SetID("any_id").
					SetEmail("any_email").
					SetPassword("any_password").
					SetCompanyName("any_company_name").
					SetBusinessName("any_business_name").
					SetCNPJ("11122233344").
					SetEmail("any_email").
					SetFundationDate(time.Now()).
					SetCreatedAT(time.Now()).
					SetUpdatedAT(time.Now()).
					Build()

				validate := entity.Validate()

				Expect(validate.Errors()).To(BeNil())
			})
		})

		Context("error cases", func() {
			It("should return an error when address is empty", func() {
				entity := entity.SellerBuilder().Build()

				validate := entity.Validate()

				Expect(validate.Errors()).NotTo(BeNil())
				Expect(validate.Errors()).To(HaveLen(10))
				Expect(validate.Errors()).To(ContainElement(fmt.Errorf("cnpj: %s", validator.NotBlankField)))
				Expect(validate.Errors()).To(ContainElement(fmt.Errorf("company_name: %s", validator.NotBlankField)))
			})
		})
	})
})

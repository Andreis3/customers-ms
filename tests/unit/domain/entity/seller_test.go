//go:build unit
// +build unit

package entity_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/validator"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: SELLER", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when build new seller", func() {
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

				Expect(validate.Errors()).To(BeEmpty())
			})
		})

		Context("error cases", func() {
			It("should return an error when seller is empty", func() {
				entity := entity.SellerBuilder().Build()

				validate := entity.Validate()

				Expect(validate.Errors()).NotTo(BeNil())
				Expect(validate.Errors()).To(HaveLen(5))
				Expect(validate.Errors()).To(ContainElement(fmt.Sprintf("cnpj: %s", validator.NotBlankField)))
				Expect(validate.Errors()).To(ContainElement(fmt.Sprintf("company_name: %s", validator.NotBlankField)))
			})
		})
	})
})

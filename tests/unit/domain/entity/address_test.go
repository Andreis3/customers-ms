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

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: ADDRESS", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when build new address", func() {
				entity := entity.AddressBuilder().
					SetID("any_id").
					SetCity("any_city").
					SetComplement("any_password").
					SetStreet("any_street").
					SetCountry("any_country").
					SetNumber("any_number").
					SetState("any_state").
					SetPostalCode("any_postal_code").
					SetCreatedAT(time.Now()).
					SetUpdatedAT(time.Now()).
					Build()

				validate := entity.Validate()

				Expect(validate.Errors()).To(BeNil())
			})
		})

		Context("error cases", func() {
			It("should return an error when address is empty", func() {
				entity := entity.AddressBuilder().Build()

				validate := entity.Validate()

				Expect(validate.Errors()).NotTo(BeNil())
				Expect(validate.Errors()).To(HaveLen(12))
				Expect(validate.Errors()).To(ContainElement(fmt.Errorf("country: %s", validator.NotBlankField)))
				Expect(validate.Errors()).To(ContainElement(fmt.Errorf("state: %s", validator.NotBlankField)))
			})
		})
	})
})

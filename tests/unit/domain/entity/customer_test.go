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

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: CUSTOMER", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when build new customer", func() {
				entity := entity.CustomerBuilder().
					SetID("any_id").
					SetPassword("any_password").
					SetFirstName("any_first_name").
					SetLasName("any_last_name").
					SetCPF("11122233344").
					SetEmail("any_email").
					SetDateOfBirth(time.Now()).
					SetCreatedAT(time.Now()).
					SetUpdatedAT(time.Now()).
					Build()

				validate := entity.Validate()

				Expect(validate.Errors()).To(BeEmpty())
			})
		})

		Context("error cases", func() {
			It("should return an error when customer is empty", func() {
				entity := entity.CustomerBuilder().Build()

				validate := entity.Validate()

				Expect(validate.Errors()).NotTo(BeNil())
				Expect(validate.Errors()).To(HaveLen(4))
				Expect(validate.Errors()).To(ContainElement(fmt.Sprintf("email: %s", validator.NotBlankField)))
				Expect(validate.Errors()).To(ContainElement(fmt.Sprintf("first_name: %s", validator.NotBlankField)))
			})
		})
	})
})

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

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: CUSTOMER", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when build new customer", func() {
				entity := entity.CustomerBuilder().
					SetID("any_id").
					SetEmail("any_email").
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

				Expect(validate.Errors()).To(BeNil())
			})
		})
		Context("error cases", func() {
			It("should return an error when apikey is empty", func() {
				entity := entity.CustomerBuilder().Build()

				validate := entity.Validate()

				Expect(validate.Errors()).NotTo(BeNil())
				Expect(validate.Errors()).To(HaveLen(8))
				Expect(validate.Errors()).To(ContainElement(fmt.Errorf("email: %s", validator.NotBlankField)))
				Expect(validate.Errors()).To(ContainElement(fmt.Errorf("first_name: %s", validator.NotBlankField)))
			})
		})
	})
})

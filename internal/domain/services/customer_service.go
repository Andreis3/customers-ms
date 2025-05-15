package services

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/interfaces"
)

type CustomerService struct {
	customerRepository interfaces.CustomerRepository
}

func NewCustomerService(customerRepository interfaces.CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (s *CustomerService) ExistCustomerByEmail(ctx context.Context, email string) bool {
	customer, err := s.customerRepository.FindCustomerByEmail(ctx, email)
	return err == nil && customer != nil
}

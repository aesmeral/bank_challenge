package services

import (
	"errors"

	"github.com/aesmeral/bank_challenge/models"
)

type TransactionData struct {
	Value       int64
	Type        string
	Description string
}

func (t TransactionData) ValidateAndPerform(customer models.Customer) (map[string]interface{}, error) {
	validatorMessage, err := t.validator()
	if err != nil {
		return map[string]interface{}{
			"message": validatorMessage,
		}, err
	}

	transactionMessage, err := t.perform(customer)
	if err != nil {
		return map[string]interface{}{
			"message": "Transaction failed",
		}, err
	}

	return transactionMessage, nil
}

func (t TransactionData) validator() (string, error) {
	if t.Value <= 0 {
		return "Value must be greater than 0", errors.New("value must be greater than 0")
	}

	if t.Type != "c" && t.Type != "d" {
		return "Type must be c or d", errors.New("type must be c or d")
	}

	if len(t.Description) < 1 || len(t.Description) > 10 {
		return "Description must be between 1 and 10 characters", errors.New("description must be between 1 and 10 characters")
	}

	return "Transaction created", nil
}

func (t TransactionData) perform(customer models.Customer) (map[string]interface{}, error) {
	// code to perform transaction
	return map[string]interface{}{
		"message": "Transaction performed",
	}, nil
}

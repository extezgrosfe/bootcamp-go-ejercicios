package internal

import (
	"errors"
	"fmt"
)

var (
	ErrTransactionNotFound  = errors.New("la transaccion no existe")
	ErrTransactionsNotFound = errors.New("no se encontraron transacciones")
	ErrCreateTransaction    = errors.New("error al crear la transacción")
)

type Service interface {
	GetAll() (Transactions, error)
	Get(id int) (*Transaction, error)
	Create(transaction *Transaction) (*Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// GetAll returns all transactions
func (s *service) GetAll() (Transactions, error) {
	transactions := s.repository.GetAll()
	if len(transactions) == 0 {
		return nil, ErrTransactionsNotFound
	}
	return transactions, nil
}

// Get returns a transaction by id
func (s *service) Get(id int) (*Transaction, error) {
	transaction := s.repository.Get(id)
	if transaction == nil {
		return nil, ErrTransactionNotFound
	}

	return transaction, nil
}

// Create creates a new transaction
func (s *service) Create(transaction *Transaction) (*Transaction, error) {
	transaction, err := s.repository.Store(transaction)
	if err != nil {
		fmt.Println(err)
		return nil, ErrCreateTransaction
	}

	return transaction, nil
}

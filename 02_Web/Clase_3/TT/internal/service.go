package internal

import (
	"errors"
	"fmt"
)

var (
	ErrTransactionNotFound  = errors.New("la transaccion no existe")
	ErrTransactionsNotFound = errors.New("no se encontraron transacciones")
	ErrCreateTransaction    = errors.New("error al crear la transacción")
	ErrUpdateTransaction    = errors.New("error al actualizar la transacción")
	ErrNilTransaction       = errors.New("la transaccion que se intenta guardar es nula")
)

type Service interface {
	GetAll() (Transactions, error)
	Get(id int) (*Transaction, error)
	Create(transaction *Transaction) (*Transaction, error)
	Update(transaction *Transaction) (*Transaction, error)
	UpdateCodigoAndMonto(id int, codigo string, monto float64) (*Transaction, error)
	Delete(id int) error
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
	transactions, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, ErrTransactionsNotFound
	}
	return transactions, nil
}

// Get returns a transaction by id
func (s *service) Get(id int) (*Transaction, error) {
	transaction, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}

	return transaction, nil
}

// Update updates a transaction
func (s *service) Update(transaction *Transaction) (*Transaction, error) {
	transaction, err := s.repository.Update(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// UpdateCodigoAndMonto updates a transaction monto and codigo
func (s *service) UpdateCodigoAndMonto(id int, codigo string, monto float64) (*Transaction, error) {
	transaction, err := s.repository.Get(id)
	if transaction == nil {
		return nil, ErrTransactionNotFound
	}
	if err != nil {
		return nil, err
	}

	if codigo != "" {
		transaction.Codigo = codigo
	}
	if monto != 0 {
		transaction.Monto = monto
	}

	transaction, err = s.repository.Update(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Delete deletes a transaction
func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

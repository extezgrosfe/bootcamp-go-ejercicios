package internal

import (
	"encoding/json"
	"os"
)

type Transaction struct {
	Id       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
}

type Transactions []*Transaction

type Repository interface {
	GetAll() (Transactions, error)
	Get(id int) (*Transaction, error)
	Store(transaction *Transaction) (*Transaction, error)
	Update(transaction *Transaction) (*Transaction, error)
	Delete(id int) error
}

type repository struct {
	fileName string
}

func NewRepository(fileName string) Repository {
	return &repository{fileName}
}

// GetAll returns all transactions
func (r *repository) GetAll() (Transactions, error) {
	return r.readFile()
}

// Get returns a transaction by id
func (r *repository) Get(id int) (*Transaction, error) {
	transactions, err := r.readFile()
	if err != nil {
		return nil, err
	}

	for _, t := range transactions {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, ErrTransactionNotFound
}

// getLastId returns the last id
func (r *repository) getLastId() int {
	transactions, err := r.readFile()
	if err != nil {
		return -1
	}

	if len(transactions) == 0 {
		return 0
	}
	return transactions[len(transactions)-1].Id
}

// Store stores a transaction
func (r *repository) Store(transaction *Transaction) (*Transaction, error) {
	transactions, err := r.readFile()
	if err != nil {
		return nil, err
	}

	if transaction == nil {
		return nil, ErrCreateTransaction
	}

	transaction.Id = r.getLastId() + 1
	if transaction.Id == 0 {
		return nil, ErrCreateTransaction
	}

	transactions = append(transactions, transaction)
	return transaction, r.writeFile(transactions)
}

// Update updates a transaction
func (r *repository) Update(transaction *Transaction) (*Transaction, error) {
	if transaction == nil {
		return nil, ErrNilTransaction
	}

	transactions, err := r.readFile()
	if err != nil {
		return nil, err
	}

	for i, t := range transactions {
		if t.Id == transaction.Id {
			transactions[i] = transaction
			return transaction, r.writeFile(transactions)
		}
	}
	return nil, ErrTransactionNotFound
}

// Delete deletes a transaction
func (r *repository) Delete(id int) error {
	transactions, err := r.readFile()
	if err != nil {
		return err
	}

	for i, t := range transactions {
		if t.Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return r.writeFile(transactions)
		}
	}
	return ErrTransactionNotFound
}

// writeFile writes transactions to a file
func (r *repository) writeFile(transactions Transactions) error {
	fileData, err := json.Marshal(transactions)
	if err != nil {
		return err
	}

	return os.WriteFile(r.fileName, fileData, 0644)
}

// readFile reads data from a file
func (r *repository) readFile() (Transactions, error) {
	var transactions Transactions
	fileData, err := os.ReadFile(r.fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileData, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

package internal

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

var transactions Transactions

type Repository interface {
	GetAll() Transactions
	Get(id int) *Transaction
	Store(transaction *Transaction) (*Transaction, error)
	Update(transaction *Transaction) (*Transaction, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// GetAll returns all transactions
func (r *repository) GetAll() Transactions {
	return transactions
}

// Get returns a transaction by id
func (r *repository) Get(id int) *Transaction {
	for _, t := range transactions {
		if t.Id == id {
			return t
		}
	}
	return nil
}

// getLastId returns the last id
func getLastId() int {
	if len(transactions) == 0 {
		return 0
	}
	return transactions[len(transactions)-1].Id
}

// Store stores a transaction
func (r *repository) Store(transaction *Transaction) (*Transaction, error) {
	if transaction == nil {
		return nil, ErrCreateTransaction
	}

	transaction.Id = getLastId() + 1

	transactions = append(transactions, transaction)
	return transaction, nil
}

// Update updates a transaction
func (r *repository) Update(transaction *Transaction) (*Transaction, error) {
	if transaction == nil {
		return nil, ErrNilTransaction
	}

	for i, t := range transactions {
		if t.Id == transaction.Id {
			transactions[i] = transaction
			return transaction, nil
		}
	}
	return nil, ErrTransactionNotFound
}

// Delete deletes a transaction
func (r *repository) Delete(id int) error {
	for i, t := range transactions {
		if t.Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}
	return ErrTransactionNotFound
}

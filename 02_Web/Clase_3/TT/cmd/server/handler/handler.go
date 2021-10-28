package handler

import (
	"errors"
	"goweb_clase3_tt/internal"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrTransactionAmountRequired   = errors.New("el campo monto es requerido")
	ErrTransactionCodeRequired     = errors.New("el campo codigo es requerido")
	ErrTransactionDateRequired     = errors.New("el campo fecha es requerido")
	ErrTransactionSenderRequired   = errors.New("el campo emisor es requerido")
	ErrTransactionReceiverRequired = errors.New("el campo receptor es requerido")
	ErrTransactionCurrencyRequired = errors.New("el campo moneda es requerido")
)

type Error struct {
	Message string `json:"message"`
}

type Handler interface {
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	UpdateCodigoAndMonto() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type handler struct {
	service internal.Service
}

func NewHandler(service internal.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !verifyToken(c.GetHeader("token")) {
			c.String(http.StatusUnauthorized, "acceso denegado")
			return
		}
		transactions, err := h.service.GetAll()
		if err != nil {
			if errors.Is(err, internal.ErrTransactionsNotFound) {
				c.JSON(http.StatusNotFound, Error{Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, transactions)
	}
}

func (h *handler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !verifyToken(c.GetHeader("token")) {
			c.String(http.StatusUnauthorized, "acceso denegado")
			return
		}
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		transaction, err := h.service.Get(id)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, Error{Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, transaction)
	}
}

func (h *handler) Create() gin.HandlerFunc {
	type request = internal.Transaction
	return func(c *gin.Context) {
		if !verifyToken(c.GetHeader("token")) {
			c.String(http.StatusUnauthorized, "acceso denegado")
			return
		}
		var transaction request
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		err := validateTransaction(&transaction)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, Error{Message: err.Error()})
			return
		}

		newTransaction, err := h.service.Create(&transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newTransaction)
	}
}

func (h *handler) Update() gin.HandlerFunc {
	type request = internal.Transaction
	return func(c *gin.Context) {
		if !verifyToken(c.GetHeader("token")) {
			c.String(http.StatusUnauthorized, "acceso denegado")
			return
		}
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		var transaction request
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		err = validateTransaction(&transaction)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, Error{Message: err.Error()})
			return
		}

		transaction.Id = id

		updatedTransaction, err := h.service.Update(&transaction)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, Error{Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedTransaction)
	}
}

func (h *handler) UpdateCodigoAndMonto() gin.HandlerFunc {
	type request struct {
		Codigo string  `json:"codigo"`
		Monto  float64 `json:"monto"`
	}
	return func(c *gin.Context) {
		if !verifyToken(c.GetHeader("token")) {
			c.String(http.StatusUnauthorized, "acceso denegado")
			return
		}
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		updatedTransaction, err := h.service.UpdateCodigoAndMonto(id, req.Codigo, req.Monto)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, Error{Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedTransaction)
	}
}

func (h *handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !verifyToken(c.GetHeader("token")) {
			c.String(http.StatusUnauthorized, "acceso denegado")
			return
		}
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		}

		err = h.service.Delete(id)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, Error{Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// validateTransaction asserts that the transaction is valid and every field is not empty
func validateTransaction(transaction *internal.Transaction) error {
	if transaction.Monto == 0 {
		return ErrTransactionAmountRequired
	}
	if transaction.Codigo == "" {
		return ErrTransactionCodeRequired
	}
	if transaction.Fecha == "" {
		return ErrTransactionDateRequired
	}
	if transaction.Emisor == "" {
		return ErrTransactionSenderRequired
	}
	if transaction.Receptor == "" {
		return ErrTransactionReceiverRequired
	}
	if transaction.Moneda == "" {
		return ErrTransactionCurrencyRequired
	}

	return nil
}

// verifyToken verifies that the token is valid
func verifyToken(token string) bool {
	return token == os.Getenv("API_TOKEN")
}

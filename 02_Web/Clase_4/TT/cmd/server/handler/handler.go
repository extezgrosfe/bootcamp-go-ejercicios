package handler

import (
	"errors"
	"goweb_clase4_tt/internal"
	"goweb_clase4_tt/pkg/web"
	"net/http"
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

// GetAll godoc
// @Summary Get all transactions
// @Description Get all transactions
// @Tags transactions
// @Accept  json
// @Produce  json
// @Success 200 {array} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /transactions [get]
func (h *handler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactions, err := h.service.GetAll()
		if err != nil {
			if errors.Is(err, internal.ErrTransactionsNotFound) {
				c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err))
				return
			}
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transactions, nil))
	}
}

// GetById godoc
// @Summary Get transaction by id
// @Description Get transaction by id
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /transactions/{id} [get]
func (h *handler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		transaction, err := h.service.Get(id)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err))
				return
			}
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transaction, nil))
	}
}

// Create godoc
// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body internal.Transaction true "Transaction"
// @Success 201 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 422 {object} web.Response
// @Router /transactions [post]
func (h *handler) Create() gin.HandlerFunc {
	type request = internal.Transaction
	return func(c *gin.Context) {
		var transaction request
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		err := validateTransaction(&transaction)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err))
			return
		}

		newTransaction, err := h.service.Create(&transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err))
			return
		}
		c.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, newTransaction, nil))
	}
}

// Update godoc
// @Summary Update a transaction
// @Description Update a transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Param transaction body internal.Transaction true "Transaction"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 422 {object} web.Response
// @Router /transactions/{id} [put]
func (h *handler) Update() gin.HandlerFunc {
	type request = internal.Transaction
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		var transaction request
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		err = validateTransaction(&transaction)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err))
			return
		}

		transaction.Id = id

		updatedTransaction, err := h.service.Update(&transaction)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err))
				return
			}
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedTransaction, nil))
	}
}

// UpdateCodigoAndMonto godoc
// @Summary Update codigo and monto of a transaction
// @Description Update codigo and monto of a transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Param transaction body internal.Transaction true "Transaction"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /transactions/{id} [patch]
func (h *handler) UpdateCodigoAndMonto() gin.HandlerFunc {
	type request struct {
		Codigo string  `json:"codigo"`
		Monto  float64 `json:"monto"`
	}
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		updatedTransaction, err := h.service.UpdateCodigoAndMonto(id, req.Codigo, req.Monto)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err))
				return
			}
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedTransaction, nil))
	}
}

// Delete godoc
// @Summary Delete a transaction
// @Description Delete a transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} web.Response
// @Failure 500 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /transactions/{id} [delete]
func (h *handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
			return
		}

		err = h.service.Delete(id)
		if err != nil {
			if errors.Is(err, internal.ErrTransactionNotFound) {
				c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err))
				return
			}
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, nil, nil))
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

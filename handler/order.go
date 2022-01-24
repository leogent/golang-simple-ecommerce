package handler

import (
	"net/http"
	"simple-ecommerce/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	OrderProduct(*gin.Context)
}

type orderHandler struct {
	repo repository.OrderRepository
}

func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo: repository.NewOrderRepository(),
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context) {
	prodIDstr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodIDstr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quantityIDstr := ctx.Param("quantity")
	quantityID, err := strconv.Atoi(quantityIDstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetFloat64("userID")
	if err := h.repo.OrderProduct(int(userID), prodID, quantityID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, "Product Successfully Ordered")
}

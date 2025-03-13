package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createRentalRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

func (h *Handler) createRental(ctx *gin.Context) {
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	instrumentId, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid instrument id")
		return
	}

	var req createRentalRequest

	if err := ctx.BindJSON(&req); err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Rent.CreateRental(callerId, instrumentId, req.StartDate, req.EndDate)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getRental(ctx *gin.Context) {
	rentalId, err := strconv.Atoi(ctx.Param("rental_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid rental id")
		return
	}

	rental, err := h.service.Rent.GetRental(rentalId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rental)
}

func (h *Handler) getUserRentals(ctx *gin.Context) {
	userId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	rentals, err := h.service.Rent.GetUserRentals(userId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rentals)
}

func (h *Handler) getInstrumentRentals(ctx *gin.Context) {
	instrumentId, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid instrument id")
		return
	}

	rentals, err := h.service.Rent.GetInstrumentRentals(instrumentId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rentals)
}

func (h *Handler) returnInstrument(ctx *gin.Context) {
	rentalId, err := strconv.Atoi(ctx.Param("rental_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid rental id")
		return
	}

	if err := h.service.Rent.ReturnInstrument(rentalId); err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

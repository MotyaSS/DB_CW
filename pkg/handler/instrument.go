package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func setupInstFilter(ctx *gin.Context) (entity.InstFilter, error) {
	var f entity.InstFilter

	// Handle multiple categories
	categories := ctx.QueryArray("category")
	for _, category := range categories {
		f.AddCategory(category)
	}

	// Handle multiple manufacturers
	manufacturers := ctx.QueryArray("manufacturer")
	for _, manufacturer := range manufacturers {
		f.AddManufacturer(manufacturer)
	}

	// Handle price floor
	if priceFloor := ctx.Query("price_floor"); priceFloor != "" {
		price, err := decimal.NewFromString(priceFloor)
		if err != nil {
			return f, fmt.Errorf("incorrect price_floor format")
		}
		f.AddPriceFloor(price)
	}

	// Handle price ceiling
	if priceCeil := ctx.Query("price_ceil"); priceCeil != "" {
		price, err := decimal.NewFromString(priceCeil)
		if err != nil {
			return f, fmt.Errorf("incorrect price_ceil format")
		}
		f.AddPriceCeil(price)
	}

	// Handle pagination
	param := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(param)
	if err != nil {
		return f, fmt.Errorf("incorrect page format")
	}
	f.AddPage(page)

	return f, nil
}

func (h *Handler) getAllInstruments(ctx *gin.Context) {
	// TODO: should return instruments with discount if exists
	filter, err := setupInstFilter(ctx)
	if err != nil {
		httpErr := &httpError.ErrorWithStatusCode{}
		ok := errors.As(err, &httpErr)
		if ok {
			abortWithStatusCode(ctx, httpErr.HTTPStatus, httpErr.Msg)
		} else {
			abortWithStatusCode(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	res, err := h.service.Instrument.GetAllInstruments(filter)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"items": res},
	)
}

func (h *Handler) getInstrument(ctx *gin.Context) {
	// TODO: should return instrument with discount if exists
	id, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithError(ctx, err)
	}
	inst, err := h.service.Instrument.GetInstrument(id)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		inst,
	)
}

func (h *Handler) createInstrument(ctx *gin.Context) {
	var inst entity.Instrument
	if err := ctx.BindJSON(&inst); err != nil {
		abortWithError(ctx, err)
		return
	}
	id, err := h.service.Instrument.CreateInstrument(inst)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"id": id},
	)
}

func (h *Handler) deleteInstrument(ctx *gin.Context) {
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	instId, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	err = h.service.Instrument.DeleteInstrument(callerId, instId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusNoContent, nil,
	)
}

func (h *Handler) rentInstrument(ctx *gin.Context) {
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	instId, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	id, err := h.service.Rent.CreateRental(callerId, instId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"id": id},
	)
}

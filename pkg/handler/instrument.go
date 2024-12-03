package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func setupInstFilter(ctx *gin.Context) (entity.InstFilter, error) {
	var f entity.InstFilter
	param, exists := ctx.GetQuery("category")
	if exists {
		f.AddCategory(param)
	}
	param, exists = ctx.GetQuery("manufacturer")
	if exists {
		f.AddManufacturer(param)
	}

	param = ctx.DefaultQuery("page", "0")
	page, err := strconv.Atoi(param)
	if err != nil {
		return f, fmt.Errorf("incorrect page format")
	}
	f.AddPage(page)

	param, exists = ctx.GetQuery("floor")
	if exists {
		d, err := decimal.NewFromString(param)
		if err != nil {
			return f, fmt.Errorf("incorrect price format")
		}
		f.AddPriceFloor(&d)
	}

	param, exists = ctx.GetQuery("ceil")
	if exists {
		d, err := decimal.NewFromString(param)
		if err != nil {
			return f, fmt.Errorf("incorrect price format")
		}
		f.AddPriceCeil(&d)
	}

	return f, nil
}

func (h *Handler) getAllInstruments(ctx *gin.Context) {
	a, err := setupInstFilter(ctx)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, err.Error())
		return
	}

	slog.Info("InstFilter initialized correctly", "InstFilter", a)
	res, err := h.service.Instrument.GetAllInstruments(a)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (h *Handler) getInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"instrument number ",
			ctx.Param("inst_id"),
		},
	)
}

func (h *Handler) addInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		"instrument added",
	)
}

func (h *Handler) deleteInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"instrument deleted",
			ctx.Param("inst_id"),
		},
	)
}

func (h *Handler) rentInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		"instrument rented",
	)
}

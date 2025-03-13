package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// Структура для приема данных с фронтенда
type repairRequest struct {
	RepairStartDate string          `json:"repair_start_date" binding:"required"`
	RepairEndDate   string          `json:"repair_end_date" binding:"required"`
	RepairCost      decimal.Decimal `json:"repair_cost" binding:"required"`
	Description     string          `json:"description" binding:"required"`
}

func (h *Handler) createRepair(ctx *gin.Context) {
	// Получаем ID пользователя из контекста
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	// Получаем ID инструмента из URL
	instrumentId, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid instrument id")
		return
	}

	var req repairRequest
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	// Парсим даты из строк в time.Time
	startDate, err := time.Parse("2006-01-02", req.RepairStartDate)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid start date format")
		return
	}

	endDate, err := time.Parse("2006-01-02", req.RepairEndDate)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid end date format")
		return
	}

	repair := entity.Repair{
		InstrumentId:    instrumentId,
		RepairStartDate: startDate,
		RepairEndDate:   endDate,
		RepairCost:      req.RepairCost,
		Description:     req.Description,
	}

	id, err := h.service.Repair.CreateRepair(callerId, repair)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllRepairs(ctx *gin.Context) {
	// Получаем ID пользователя из контекста
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	// Получаем ID инструмента из URL
	instrumentId, err := strconv.Atoi(ctx.Param("inst_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid instrument id")
		return
	}
	repairs, err := h.service.Repair.GetInstrumentRepairs(callerId, instrumentId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, repairs)
}

func (h *Handler) getRepair(ctx *gin.Context) {
	// Получаем ID пользователя из контекста
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	// Получаем ID ремонта из URL
	repairId, err := strconv.Atoi(ctx.Param("repair_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid repair id")
		return
	}

	repair, err := h.service.Repair.GetRepair(callerId, repairId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, repair)
}

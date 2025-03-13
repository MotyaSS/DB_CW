package handler

import (
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBackup(c *gin.Context) {
	// Получаем ID пользователя из контекста для проверки прав
	callerId, err := h.getCallerId(c)
	if err != nil {
		abortWithError(c, err)
		return
	}

	// Проверяем права доступа (только администратор может управлять бэкапами)
	if err := h.service.Authorisation.CheckPermission(callerId, entity.RoleAdmin); err != nil {
		abortWithError(c, err)
		return
	}

	backupPath, err := h.service.Backup.CreateBackup()
	if err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"backup_path": backupPath,
	})
}

func (h *Handler) listBackups(c *gin.Context) {
	// Получаем ID пользователя из контекста для проверки прав
	callerId, err := h.getCallerId(c)
	if err != nil {
		abortWithError(c, err)
		return
	}

	// Проверяем права доступа (только администратор может управлять бэкапами)
	if err := h.service.Authorisation.CheckPermission(callerId, entity.RoleAdmin); err != nil {
		abortWithError(c, err)
		return
	}

	backups, err := h.service.Backup.ListBackups()
	if err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"backups": backups,
	})
}

func (h *Handler) restoreBackup(c *gin.Context) {
	// Получаем ID пользователя из контекста для проверки прав
	callerId, err := h.getCallerId(c)
	if err != nil {
		abortWithError(c, err)
		return
	}

	// Проверяем права доступа (только администратор может управлять бэкапами)
	if err := h.service.Authorisation.CheckPermission(callerId, entity.RoleAdmin); err != nil {
		abortWithError(c, err)
		return
	}

	var input struct {
		BackupPath string `json:"backup_path" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		abortWithStatusCode(c, http.StatusBadRequest, "invalid input")
		return
	}

	if err := h.service.Backup.RestoreFromBackup(input.BackupPath); err != nil {
		abortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "backup restored successfully",
	})
}

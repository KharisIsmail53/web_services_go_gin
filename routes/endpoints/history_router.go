package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HistroyRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/history", func(c *gin.Context) {
		var history []models.History
		db.Order("created_at DESC").Find(&history)
		c.JSON(http.StatusOK, schema.HistorySchema{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    history,
		})
	})

	router.GET("/history/:id", func(c *gin.Context) {
		var history models.HistoryByID
		if err := db.Preload("TransaksiByID").
					 Preload("TransaksiByID.Akad").
					 Preload("TransaksiByID.Beras").
					 Preload("TransaksiByID.Zakat").
					 Where("id = ?", c.Param("id")).First(&history).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error":"History not found"})
			return
		}
		c.JSON(http.StatusOK, schema.HistoryByID{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    history,
		})
	})

	
}
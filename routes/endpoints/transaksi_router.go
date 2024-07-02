package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"gin-framework-services/services"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransaksiRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/transaksi", func(c *gin.Context) {
		var transaksi []models.Transaksi
		db.Order("created_at DESC").Find(&transaksi)
		c.JSON(http.StatusOK, schema.TransaksiSchema{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    transaksi,
		})
	})

	router.GET("/transaksi/:id", func(c *gin.Context) {
		var transaksi models.TransaksiByID
		if db.Preload("Akad").Preload("Zakat").Preload("Beras").Where("id = ?", c.Param("id")).First(&transaksi).Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi not found"})
			return
		}
		c.JSON(http.StatusOK, schema.TransaksiByID{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    transaksi,
		})
	})

	router.POST("/transaksi", func(c *gin.Context) {
		var transaksi models.Transaksi
		if err := c.ShouldBindJSON(&transaksi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(transaksi)
		services.CreateTransaksi(c, db, &transaksi)
	})
}
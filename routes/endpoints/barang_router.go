package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BarangRouter(router *gin.Engine, db *gorm.DB) {
	// Endpoint GET untuk mengambil semua user
	router.GET("/barang", func(c *gin.Context) {
		var barang []models.Barang
		db.Order("created_at DESC").Find(&barang)
		response := schema.BarangSchema{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    barang,
		}
		c.JSON(http.StatusOK, response)
	})
	
	router.GET("/barang/:id", func(c *gin.Context) {
		var barang models.BarangByID
		if err := db.Preload("Kategori").Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
			return
		}
		responds := schema.BarangByID{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    barang,
		}
		c.JSON(http.StatusOK, responds)
	})

	// Endpoint POST untuk menambahkan user baru
	router.POST("/barang", func(c *gin.Context) {
		var barang models.Barang
		if err := c.ShouldBindJSON(&barang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&barang).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := schema.BarangCRUD{
			Message: "Data Has Been Created",
			Status:  http.StatusCreated,
			Data:    barang,
		}
		c.JSON(http.StatusCreated, response)
	})

	router.PUT("/barang/:id", func(c *gin.Context) {
		var barang models.Barang
		if err := db.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
			return
		}
		if err := c.ShouldBindJSON(&barang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&barang).Updates(&barang).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responds := schema.BarangCRUD{
			Message: "Data Has Been Updated",
			Status:  http.StatusOK,
			Data:    barang,
		}
		c.JSON(http.StatusOK, responds)
	})

	router.DELETE("/barang/:id", func(c *gin.Context) {
		var barang models.Barang
		if err := db.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
			return
		}
		if err := db.Delete(&barang).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responds := schema.BarangCRUD{
			Message: "Data Has Been Deleted",
			Status:  http.StatusOK,
			Data:    barang,
		}
		c.JSON(http.StatusOK, responds)
	})


}

package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KategoriRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/kategori", func(c *gin.Context) {
		var kategori []models.Kategori
		db.Order("created_at DESC").Find(&kategori)
		responds := schema.KategoriSchema{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    kategori,
		}
		c.JSON(http.StatusOK, responds)
	})

	router.GET("/kategori/:id", func(c *gin.Context) {
		var kategori models.Kategori
		if err := db.Where("id = ?", c.Param("id")).First(&kategori).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
			return
		}
		responds := schema.KategoriByID{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    kategori,
		}
		c.JSON(http.StatusOK, responds)
	})

	router.POST("/kategori", func(c *gin.Context) {
		var kategori models.Kategori
		if err := c.ShouldBindJSON(&kategori); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&kategori).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.KategoriCRUD{
			Message: "Data Has Been Created",
			Status:  http.StatusOK,
			Data:    kategori,
		})
	})

	router.PUT("/kategori/:id", func(c *gin.Context) {
		var kategori models.Kategori
		if err := db.Where("id = ?", c.Param("id")).First(&kategori).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
			return
		}
		if err := c.ShouldBindJSON(&kategori); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&kategori).Updates(&kategori).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responds := schema.KategoriCRUD{
			Message: "Data Has Been Updated",
			Status:  http.StatusOK,
			Data:    kategori,
		}
		c.JSON(http.StatusOK, responds)
	})

	router.DELETE("/kategori/:id", func(c *gin.Context) {
		var kategori models.Kategori
		if err := db.Where("id = ?", c.Param("id")).First(&kategori).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
			return
		}
		if err := db.Delete(&kategori).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responds := schema.KategoriCRUD{
			Message: "Data Has Been Deleted",
			Status:  http.StatusOK,
			Data:    kategori,
		}
		c.JSON(http.StatusOK, responds)
	})
}

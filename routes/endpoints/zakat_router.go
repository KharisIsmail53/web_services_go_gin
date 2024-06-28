package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ZakatRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/zakat", func(c *gin.Context) {
		var zakat []models.Zakat
		db.Order("jenis_zakat ASC").Find(&zakat)
		c.JSON(http.StatusOK, schema.ZakatSchema{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    zakat,
		})
	})

	router.GET("/zakat/:id", func(c *gin.Context) {
		var zakat models.Zakat
		if err := db.Where("id = ?", c.Param("id")).First(&zakat).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.ZakatByID{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    zakat,
		})
	})

	router.POST("/zakat", func(c *gin.Context) {
		var zakat models.Zakat
		if err := c.ShouldBindJSON(&zakat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&zakat).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, schema.ZakatCRUD{
			Message: "Data Has Been Created",
			Status:  http.StatusCreated,
			Data:    zakat,
		})
	})

	router.PUT("/zakat/:id", func(c *gin.Context) {
		var zakat models.Zakat
		if err := db.Where("id = ?", c.Param("id")).First(&zakat).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindJSON(&zakat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&zakat).Updates(&zakat).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.ZakatCRUD{
			Message: "Data Has Been Updated",
			Status:  http.StatusOK,
			Data:    zakat,
		})
	})

	router.DELETE("/zakat/:id", func(c *gin.Context) {
		var zakat models.Zakat
		if err := db.Where("id = ?", c.Param("id")).First(&zakat).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := db.Delete(&zakat).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.ZakatCRUD{
			Message: "Data Has Been Deleted",
			Status:  http.StatusOK,
			Data:    zakat,
		})
	})
}
package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BerasRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/beras", func(c *gin.Context) {
		var beras []models.Beras
		db.Order("harga_beras DESC").Find(&beras)
		c.JSON(http.StatusOK, schema.BerasSchema{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    beras,
		})
	})

	router.GET("/beras/:id", func(c *gin.Context) {
		var beras models.Beras
		if err := db.Where("id = ?", c.Param("id")).First(&beras).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.BerasByID{
			Message: "Data Received",
			Status:  http.StatusOK,
			Data:    beras,
		})
	})

	router.POST("/beras", func(c *gin.Context) {
		var beras models.Beras
		if err := c.ShouldBindJSON(&beras); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&beras).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.BerasCRUD{
			Message: "Data Has Been Created",
			Status:  http.StatusOK,
			Data:    beras,
		})
	})

	router.PUT("/beras/:id", func(c *gin.Context) {
		var beras models.Beras
		if err := db.Where("id = ?", c.Param("id")).First(&beras).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindJSON(&beras); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&beras).Updates(&beras).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.BerasCRUD{
			Message: "Data Has Been Updated",
			Status:  http.StatusOK,
			Data:    beras,
		})
	})

	router.DELETE("/beras/:id", func(c *gin.Context) {
		var beras models.Beras
		if err := db.Where("id = ?", c.Param("id")).First(&beras).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := db.Delete(&beras).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.BerasCRUD{
			Message: "Data Has Been Deleted",
			Status:  http.StatusOK,
			Data:    beras,
		})
	})
}
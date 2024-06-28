package endpoints

import (
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AkadRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/akad", func(c *gin.Context) {
		var akad []models.Akad
		db.Order("jenis_akad ASC").Find(&akad)
		c.JSON(http.StatusOK, schema.AkadSchema{
			Status:  http.StatusOK,
			Message: "Data Received",
			Data:    akad,
		})
	})

	router.GET("/akad/:id", func(c *gin.Context) {
		var akad models.Akad
		if err := db.Where("id = ?", c.Param("id")).First(&akad).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.AkadByID{
			AkadSchema: schema.AkadSchema{
				Status:  http.StatusOK,
				Message: "Data Received",
			},
			Data: akad,
		})
	})

	router.POST("/akad", func(c *gin.Context) {
		var akad models.Akad
		if err := c.ShouldBindJSON(&akad); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&akad).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, schema.AkadCRUD{
			AkadSchema: schema.AkadSchema{
				Status:  http.StatusCreated,
				Message: "Data Created",
			},
			Data: akad,
		})
	})

	router.PUT("/akad/:id", func(c *gin.Context) {
		var akad models.Akad
		if err := db.Where("id = ?", c.Param("id")).First(&akad).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindJSON(&akad); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&akad).Updates(&akad).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.AkadCRUD{
			AkadSchema: schema.AkadSchema{
				Status:  http.StatusOK,
				Message: "Data Updated",
			},
			Data: akad,
		})
	})

	router.DELETE("/akad/:id", func(c *gin.Context) {
		var akad models.Akad
		if err := db.Where("id = ?", c.Param("id")).First(&akad).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := db.Delete(&akad).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schema.AkadCRUD{
			AkadSchema: schema.AkadSchema{
				Status:  http.StatusOK,
				Message: "Data Deleted",
			},
			Data: akad,
		})
	})
}
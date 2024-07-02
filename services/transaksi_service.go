package services

import (
	"errors"
	"gin-framework-services/models"
	"gin-framework-services/schema"
	"gin-framework-services/crud"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransaksi(c *gin.Context, db *gorm.DB, transaksi *models.Transaksi) (schema.TransaksiCRUD, error) {
	// for fitrah beras
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return schema.TransaksiCRUD{}, tx.Error
	}

	zakat, err := crud.GetZakat(db, *transaksi.IDZakat)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return schema.TransaksiCRUD{}, err
	}
	if zakat.JenisZakat == "Fitrah" {
		akad, err := crud.GetAkad(tx, *transaksi.IDAkad)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return schema.TransaksiCRUD{}, err
		}
		if akad.JenisAkad == "Beras" {
			beras, err := crud.GetBeras(tx, *transaksi.IDBeras)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return schema.TransaksiCRUD{}, err
			}
			total_beras := 3.5 * float32(*transaksi.JumlahKeluarga)
			transaksi.JumlahLiteran = &total_beras
			update_stock_beras := beras.Stock - total_beras
			if update_stock_beras < 0 {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
				return schema.TransaksiCRUD{}, errors.New("Not enough stock")
			}
			beras.Stock = update_stock_beras
			if err := tx.Model(&beras).Updates(&beras).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return schema.TransaksiCRUD{}, err
			}
			var history_transaksi models.History
			history_transaksi.IDTransaksi = uuid.UUID(transaksi.ID)
			history_transaksi.Tahun = time.Now().Year()
			if err := tx.Create(&history_transaksi).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return schema.TransaksiCRUD{}, err
			}
			if err := tx.Create(&transaksi).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			if err := tx.Commit().Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return schema.TransaksiCRUD{}, err
			}	
			c.JSON(http.StatusCreated, schema.TransaksiCRUD{
				Message: "Transaksi berhasil ditambahkan",
				Status:  http.StatusCreated,
				Data:    *transaksi,
			})
		}
	} else if zakat.JenisZakat == "Fidyah" {
		return schema.TransaksiCRUD{}, errors.New("Tipe zakat tidak valid")
	} else {
		return schema.TransaksiCRUD{}, errors.New("Tipe zakat tidak valid")
	}
	return schema.TransaksiCRUD{}, nil
}

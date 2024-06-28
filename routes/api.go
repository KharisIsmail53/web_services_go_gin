package routes

import (
	"gin-framework-services/routes/endpoints"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func APIRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	endpoints.BarangRouter(router, db)
	endpoints.KategoriRouter(router, db)
	endpoints.BerasRouter(router, db)
	endpoints.ZakatRouter(router, db)
	endpoints.AkadRouter(router, db)
	endpoints.TransaksiRouter(router, db)

	return router
}

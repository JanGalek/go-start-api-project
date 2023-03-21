package endpoints

import (
	"github.com/gin-gonic/gin"
	"jan-galek/api/src/libs/database"
	"net/http"
)

func Homepage(c *gin.Context) {
	err, _ := database.GetConnection()

	if err != nil {
		panic("DB ERROR")
	}

	c.JSON(http.StatusOK, "ok")
}

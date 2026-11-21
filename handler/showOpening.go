package handler

import (
	"net/http"

	"github.com/galvao0/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHadler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParametter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	SendSuccess(c, "show-opening", opening)
}

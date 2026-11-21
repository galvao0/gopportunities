package handler

import (
	"net/http"

	"github.com/galvao0/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	SendSuccess(c, "list-openings", openings)
}

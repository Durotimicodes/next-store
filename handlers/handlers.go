package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcomepage(c *gin.Context) {
	//call the HTML context to render template
	c.JSON(http.StatusOK, "welcome-page.html", gin.H{"title": "Home page"})
}


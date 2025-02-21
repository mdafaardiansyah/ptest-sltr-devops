package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func WelcomeHandler(c *gin.Context) {
	name := c.Param("nama")
	if name == "" {
		c.String(http.StatusOK, "Anonymous")
		return
	}

	c.String(http.StatusOK, "Selamat datang "+name)
}

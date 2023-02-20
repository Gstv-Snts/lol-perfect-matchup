package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "eriirj"})
}

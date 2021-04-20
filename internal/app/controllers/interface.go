package controllers

import (
	"github.com/gin-gonic/gin"
	"ops-go/internal/app/models"
)

type CommonInterfaces interface {
	Get(c *gin.Context)
	Query(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Option() models.Option
}
package controllers

import (
	"net/http"
	"ameicosmeticos/app/contracts"
	"github.com/gin-gonic/gin"
)

type CanceledUsersController struct {
	contracts.ICanceledUsersService
}

func (this *CanceledUsersController) FindById(c *gin.Context) {
	response, _ := this.GetUserById(1)
	c.JSON(http.StatusOK, gin.H{"data": response})
}
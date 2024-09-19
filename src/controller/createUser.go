package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josiasbarreto-dev/meu-primeiro-crud-go/src/configuration/internal_errors"
	"github.com/josiasbarreto-dev/meu-primeiro-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := internal_errors.NewBadRequestError(fmt.Sprintf("There are some incorrect fields= %s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

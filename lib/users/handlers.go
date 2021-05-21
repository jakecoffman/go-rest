package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/rest/lib/models"
	"strconv"
)

func List(c *gin.Context) {
	results, err := models.FindUsers(c)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, results)
	}
}

func Add(c *gin.Context) {
	var instance models.User
	if err := c.BindJSON(&instance); err != nil {
		return
	}

	if err := models.InsertUser(c, &instance); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, instance)
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.GetUserById(c, id)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(200, user)
	}
}

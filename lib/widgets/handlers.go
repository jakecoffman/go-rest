package widgets

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/rest/lib/db"
	"github.com/jakecoffman/rest/lib/models"
	"log"
)

func List(c *gin.Context) {
	results := []models.Widget{}
	err := db.DB.SelectContext(c, &results, `select * from widget`)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, results)
}

func Add(c *gin.Context) {
	var instance models.Widget
	if err := c.BindJSON(&instance); err != nil {
		return
	}

	stmt, err := db.DB.PrepareNamedContext(c, `insert into widget (type, owner_id) values(:type, :owner_id) returning *`)
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}

	err = stmt.Get(&instance, instance)
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, instance)
}

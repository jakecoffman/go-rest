package authors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/rest/lib/db"
	"github.com/jakecoffman/rest/lib/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"strconv"
)

type listQuery struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
}

func list(c *gin.Context) {
	var query listQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, err.Error())
		return
	}

	order := fmt.Sprintf("%s %s", query.Sort, query.Order)
	authors, err := models.Authors(qm.Limit(query.Limit), qm.Offset(query.Offset), qm.OrderBy(order)).All(c, db.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	if authors == nil {
		authors = []*models.Author{}
	}
	c.JSON(200, authors)
}

func add(c *gin.Context) {
	var instance models.Author
	if err := c.BindJSON(&instance); err != nil {
		return
	}

	if err := instance.Insert(c, db.DB, boil.Infer()); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, instance)
}

func get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	author, err := models.FindAuthor(c, db.DB, id)
	if err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, author)
	}
}

func patch(c *gin.Context) {
	// using a map here to get patching columns
	var update map[string]interface{}
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(400, err.Error())
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))

	num, err := models.Authors(models.AuthorWhere.ID.EQ(id)).UpdateAll(c, db.DB, update)
	if num < 1 {
		c.JSON(404, "Not found")
		return
	}
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	// unfortunate sqlboiler doesn't support postgres returning clause
	get(c)
}

func remove(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	instance := models.Author{ID: id}
	i, _ := instance.Delete(c, db.DB)
	if i == 1 {
		c.JSON(200, "deleted")
	} else {
		c.JSON(404, "not found")
	}
}

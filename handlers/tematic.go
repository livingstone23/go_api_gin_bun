package handlers

import (
	"context"
	"go_gin_bun/db"
	"go_gin_bun/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

//Fuction to get all the tematics
func Tematic_get(c *gin.Context) {

	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	var tematics []model.TematicModel

	err := db.NewSelect().Model(&tematics).OrderExpr("id desc").Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tematics": tematics,
	})
}

//Function to get a tematic by id
func Tematic_get_by_id(c *gin.Context) {

	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	var tematic model.TematicModel

	//Confirm if the register exist
	exist, err_exist := db.NewSelect().Model(&tematic).Where("id = ?", c.Param("id")).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(500, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"error",
			"message": "The tematic does not exist",
		})
		return
	}



	err := db.NewSelect().Model(&tematic).Where("id = ?", c.Param("id")).Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"status":"error",
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":"success",
		"tematic": tematic,
	})
}


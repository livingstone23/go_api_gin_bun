package handlers

import (
	"context"
	"go_gin_bun/db"
	"go_gin_bun/dto"
	"go_gin_bun/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

// Fuction to get all the tematics
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


// Function to get a tematic by id
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
			"status":  "error",
			"message": "The tematic does not exist",
		})
		return
	}

	err := db.NewSelect().Model(&tematic).Where("id = ?", c.Param("id")).Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"tematic": tematic,
	})
}


// Function to create a tematic
func Tematic_post(c *gin.Context) {
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var tematic dto.TematicDto

	if err := c.ShouldBindJSON(&tematic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Validate is the name is empty
	if tematic.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Insertamos el nuevo registro  con BUN
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	model := &model.TematicModel{Name: tematic.Name, Slug: slug.Make(tematic.Name)}

	_, err := db.NewInsert().Model(model).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"tematic": tematic,
	})
}


// Function to update a tematic
func Tematic_put(c *gin.Context) {
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var tematic dto.TematicDto

	if err := c.ShouldBindJSON(&tematic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Validate is the name is empty
	if tematic.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Prepare the conectio to the database
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//ask if the register exist
	var tematic_exist model.TematicModel
	exist, err_exist := db.NewSelect().Model(&tematic_exist).Where("id = ?", c.Param("id")).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The tematic does not exist",
		})
		return
	}

	//Update the register
	model := &model.TematicModel{Name: tematic.Name, Slug: slug.Make(tematic.Name)}

	_, err := db.NewUpdate().Model(model).Where("id = ?", c.Param("id")).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"tematic": tematic,
	})
}


// Function to delete a tematic
func Tematic_delete(c *gin.Context) {
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//ask if the register exist
	var tematic_exist model.TematicModel
	exist, err_exist := db.NewSelect().Model(&tematic_exist).Where("id = ?", c.Param("id")).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The tematic does not exist",
		})
		return
	}

	//Delete the register
	_, err := db.NewDelete().Model(&tematic_exist).Where("id = ?", c.Param("id")).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "The tematic was deleted !!!",
	})
}


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

// Fuction to get all the movies
func Movie_get(c *gin.Context) {

	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	var movies []model.MovieModel

	err := db.NewSelect().Model(&movies).Relation("Tematic").OrderExpr("id desc").Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tematics": movies,
	})
}

// Function to get a movie by id
func Movie_get_by_id(c *gin.Context) {

	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//var movie model.MovieModel

	//Confirm if the register exist
	id := c.Param("id")
	exist, err_exist := db.NewSelect().Model((*model.MovieModel)(nil)).Where("id = ?", id).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(500, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The movie does not exist",
		})
		return
	}

	//Get the movie
	var movieFound model.MovieModel
	err := db.NewSelect().Model(&movieFound).Relation("Tematic").Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"movie":  movieFound,
	})
}

// Function to create a new movie
func Movie_post(c *gin.Context) {

	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var movie dto.MovieDto

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data",
		})
		return
	}

	//validate if the name is empty
	if movie.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	db := bun.NewDB(db.Connection(), mysqldialect.New())

	model := &model.MovieModel{Name: movie.Name, Slug: slug.Make(movie.Name), Description: string(movie.Description), Year: movie.Year, TematicID: movie.TematicID}

	_, err := db.NewInsert().Model(model).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error creating the movie",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"movie":  movie,
	})
}


// Function to update a movie
func Movie_put(c *gin.Context) {
	
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var movie dto.MovieDto

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data",
		})
		return
	}

	//validate if the name is empty
	if movie.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Prepare the connection to the database
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//ask if the register exist
	var movie_exist model.MovieModel
	exist, err_exist := db.NewSelect().Model(&movie_exist).Where("id = ?", c.Param("id")).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The movie does not exist",
		})
		return
	}

	//Update the register
	model := &model.MovieModel{Name: movie.Name, Slug: slug.Make(movie.Name), Description: string(movie.Description), Year: movie.Year, TematicID: movie.TematicID}

	_, err := db.NewUpdate().Model(model).Where("id = ?", c.Param("id")).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error updating the movie",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"movie":  movie,
	})
}

// Function to delete a movie
func Movie_delete(c *gin.Context) {
	
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//ask if the register exist
	var movie_exist model.MovieModel
	exist, err_exist := db.NewSelect().Model(&movie_exist).Where("id = ?", c.Param("id")).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The movie does not exist",
		})
		return
	}

	//Delete the register
	_, err := db.NewDelete().Model((*model.MovieModel)(nil)).Where("id = ?", c.Param("id")).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error deleting the movie",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
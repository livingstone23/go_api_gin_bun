package handlers

import (
	"context"
	"go_gin_bun/db"
	"go_gin_bun/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func Movie_picture_upload(c *gin.Context) {
	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"state":   "error",
			"message": "Error in the file",
		})
		return
	}

	c.SaveUploadedFile(file, "public/uploads/movies/"+file.Filename)

	//insert
	id := c.Param("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	//Prepare the connection
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	model := model.MoviePictureModel{Name: file.Filename, MovieID: idInt64}

	_, err = db.NewInsert().Model(&model).Exec(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"state":   "error",
			"message": "Error in the file",
		})
		return
	}

	//return the response
	c.JSON(http.StatusCreated, gin.H{
		"state":   "success",
		"message": "File uploaded successfully!!!",
		"file":    file.Filename,
	})
}

// List all pictures of a movie
func Movie_picture_get(c *gin.Context) {
	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	//Prepare the connection
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	var pictures []model.MoviePictureModel

	err := db.NewSelect().Model(&pictures).Where("movie_id = ?", c.Param("id")).OrderExpr("id desc").Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pictures": pictures,
	})
}

// Delete the register of a picture and the file
func Movie_picture_delete(c *gin.Context) {
	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	//Prepare the connection
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//Confirm if the register exist
	id := c.Param("id")
	exist, err_exist := db.NewSelect().Model((*model.MoviePictureModel)(nil)).Where("id = ?", id).Exists(context.TODO())
	if err_exist != nil {
		c.JSON(500, gin.H{
			"message": "Error conecting to the database",
		})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The picture does not exist",
		})
		return
	}

	//Get the picture
	var picture model.MoviePictureModel
	err := db.NewSelect().Model(&picture).Where("id = ?", id).Scan(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Error conecting to the database",
		})
		return
	}

	//Delete the file
	err = os.Remove("public/uploads/movies/" + picture.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Error deleting the file",
		})
		return
	}

	//Delete the register
	_, err = db.NewDelete().Model(&picture).Where("id = ?", id).Exec(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Error deleting the register",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "The picture has been deleted",
	})
}

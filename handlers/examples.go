package handlers

import (
	"net/http"
	"go_gin_bun/dto"
	"github.com/gin-gonic/gin"
)

//Function get is a example of how to create a handler in gin
func Example_get(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from FUNC example get",
	})
}

//function post is a example of how to create a handler in gin
func Example_post(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello from FUNC example post",
	})
}

//function post and getting model
func Example_post_with_model(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var body dto.TematicDto

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in the body",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello from FUNC example post"+ "with model: " + body.Name,
		"Authorization":c.Request.Header.Get("Authorization"),
	})
}



//function post with argument a example of how to create a handler in gin
func Example_post_argument(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	id := c.Param("id")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello from FUNC example post"+ "with argument: " + id,
	})
}

//Funcion post with parameter query string a example of how to create a handler in gin
func Example_post_query_string(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")
	
	id := c.Query("id")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello from FUNC example post"+ "with argument: " + id,
	})
}


//function put is a example of how to create a handler in gin
func Example_put(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello from FUNC example put",
	})
}

//function delete is a example of how to create a handler in gin
func Example_delete(c *gin.Context) {

	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from FUNC example delete",
	})
}

//Function to upload a file
func Example_upload_file(c *gin.Context) {
	
	//Custom header response
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"state": "error",
			"message": "Error in the file",
		})
		return
	}

	c.SaveUploadedFile(file, "public/uploads/pictures/"+file.Filename)

	c.JSON(http.StatusCreated, gin.H{
		"state": "success",
		"message": "File uploaded successfully!!!",
		"file": file.Filename,
	})
}
package main

import (
	"fmt"
	"go_gin_bun/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Startint project GO_API_GIN_BUN!!!!")

	prefix := "/api/v1/"

	router := gin.Default()
	router.GET(prefix+"ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Enable release mode, IMPORTANT ONLY FOR THE TEST!!!
	gin.SetMode(gin.ReleaseMode)

	//Example of how to create a handler in gin
	router.GET(prefix+"example", handlers.Example_get)
	router.POST(prefix+"example", handlers.Example_post)
	router.POST(prefix+"example/:id", handlers.Example_post_argument)
	router.POST(prefix+"example_model", handlers.Example_post_with_model)
	router.POST(prefix+"example_query_string", handlers.Example_post_query_string)
	router.PUT(prefix+"example", handlers.Example_put)
	router.DELETE(prefix+"example", handlers.Example_delete)
	router.POST(prefix+"upload", handlers.Example_upload_file)

	//Temmatics routes
	router.GET(prefix+"tematics", handlers.Tematic_get)
	router.GET(prefix+"tematics/:id", handlers.Tematic_get_by_id)


	errorVars := godotenv.Load()
	if errorVars != nil {
		panic("Error loading .env file")
	}

	router.Run(":" + os.Getenv("API_PORT"))

}

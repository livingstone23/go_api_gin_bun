package main

import (
	"fmt"
	"go_gin_bun/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go_gin_bun/db"
)

func main() {
	fmt.Println("Startint project GO_API_GIN_BUN!!!!")

	// Cargar variables de entorno
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    // Conectar a la base de datos
    sqldb := db.Connection()
    defer sqldb.Close() // Asegurarse de cerrar la conexión cuando main termine


	prefix := "/api/v1/"

	//Enable release mode, IMPORTANT ONLY FOR THE TEST!!!
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET(prefix+"ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Configurar proxies de confianza
    router.SetTrustedProxies([]string{"127.0.0.1"})

	

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
	router.POST(prefix+"tematics", handlers.Tematic_post)
	router.PUT(prefix+"tematics/:id", handlers.Tematic_put)
	router.DELETE(prefix+"tematics/:id", handlers.Tematic_delete)

	//Movies routes
	router.GET(prefix+"movies", handlers.Movie_get)
	router.GET(prefix+"movies/:id", handlers.Movie_get_by_id)
	router.POST(prefix+"movies", handlers.Movie_post)
	router.PUT(prefix+"movies/:id", handlers.Movie_put)
	router.DELETE(prefix+"movies/:id", handlers.Movie_delete)

	//Movies picture routes
	router.POST(prefix+"movies_picture/:id", handlers.Movie_picture_upload)
	router.GET(prefix+"movies_picture/:id", handlers.Movie_picture_get)
	router.DELETE(prefix+"movies_picture/:id", handlers.Movie_picture_delete)


	//Security routes
	router.POST(prefix+"user", handlers.Security_register)
	router.POST(prefix+"login", handlers.Security_login)
	router.POST(prefix+"secure", handlers.Security_protect)



	router.Run(":" + os.Getenv("API_PORT"))

}

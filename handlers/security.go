package handlers

import (
	"context"
	"go_gin_bun/db"
	"go_gin_bun/dto"
	"go_gin_bun/jwt"
	"go_gin_bun/middleware_custom"
	"go_gin_bun/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"golang.org/x/crypto/bcrypt"
)

// Func Security_register save a user in the database
func Security_register(c *gin.Context) {
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var body dto.UserDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data",
		})
		return
	}

	//Bind the body of the request to the struct
	//validate if the name is empty
	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Validate if the email is empty
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Validate if the PASSWORD is empty
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Prepare the connection
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//Validate is the email exist in the database
	var UserExist model.UserModel
	err_exist := db.NewSelect().Model(&UserExist).Where("email = ?", body.Email).Scan(context.TODO())
	if err_exist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The email already exist",
		})
		return
	}

	//generate the hash of the password
	cost := 8
	bytes, _ := bcrypt.GenerateFromPassword([]byte(body.Password), cost)

	model := model.UserModel{Name: body.Name, Email: body.Email, Telephone: body.Telephone, Password: string(bytes)}

	_, err := db.NewInsert().Model(&model).Exec(context.TODO())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Error creating the user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "The user has been created",
	})

}

// Func Security_login validate the user and return a token

func Security_login(c *gin.Context) {
	c.Writer.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	var body dto.LoginDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data",
		})
		return
	}

	//Validate if the email is empty
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Validate if the PASSWORD is empty
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The name is required",
		})
		return
	}

	//Prepare the connection
	db := bun.NewDB(db.Connection(), mysqldialect.New())

	//Validate if the user exist
	var user model.UserModel
	err := db.NewSelect().Model(&user).Where("email = ?", body.Email).Scan(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The user does not exist",
		})
		return
	}

	//Validate the password
	passwordbytes := []byte(body.Password)
	passwordBd := []byte(user.Password)

	//Validate the password
	errPassword := bcrypt.CompareHashAndPassword(passwordBd, passwordbytes)

	if errPassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The password is incorrect",
		})
		return
	} else {

		//Generate the token
		jwtKey, errJWT := jwt.GenerateJWT(user.Email, user.Name, user.ID)

		if errJWT != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Error generating the token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  jwtKey,
		})
		return
	}
}

// Func to confirm the middleware that confirms the token

func Security_protect(c *gin.Context) {
	if middleware_custom.ValidateJWT(c.GetHeader("Authorization")) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Authorized",
	})
}

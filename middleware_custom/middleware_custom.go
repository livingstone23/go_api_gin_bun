package middleware_custom

import (
	"context"
	"fmt"
	"go_gin_bun/db"
	"go_gin_bun/model"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func ValidateJWT(header string) int {
	errorVA := godotenv.Load()
	if errorVA != nil {
		return 0
	}

	myKey := []byte(os.Getenv("SECRET_JWT"))

	if len(header) == 0 {
		return 0
	}

	splitBearer := strings.Split(header, " ")
	if len(splitBearer) != 2 {
		return 0
	}

	splitToken := strings.Split(splitBearer[1], ".")
	if len(splitToken) != 3 {
		return 0
	}

	tk := strings.TrimSpace(splitBearer[1])
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:")
		}
		return myKey, nil
	})

	if err != nil {
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var user model.UserModel
		db := bun.NewDB(db.Connection(), mysqldialect.New())
		err := db.NewSelect().Model(&user).Where("email = ?", claims["email"]).Scan(context.TODO())

		if err != nil {
			return 0
		}
		return 1

	} else {
		return 0
	}
}

package routes

import (
	"invoice/routehandler"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	_ "invoice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title  Documenting API
// @version 1

// @localhost:8181

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(authMiddleware())
	{
		router.GET("/invoicedetails", routehandler.GetInvoiceDetails)
		router.POST("/createinvoice", routehandler.CreateInvoice)
	}

	return router
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "header missing"})
			ctx.Abort()
			return
		}
		jwtKey := os.Getenv("SECRET_KEY")

		if !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			ctx.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}

			return []byte(jwtKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token "})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

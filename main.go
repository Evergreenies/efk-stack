package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Made HealthCheck call")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	}
}

func GetDiv() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		idParsed, _ := strconv.Atoi(id)
		div := 10 / idParsed // intentiaonlly we are rising error just
		// in order to test proper loggin traceback in elasticsearch

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("division is %d", div),
		})
	}
}

func main() {
	router := gin.Default()

	log.Println("Registered a group at root")
	service := router.Group("/")
	service.GET("/health-check", HealthCheck())
	service.GET("/divide/:id", GetDiv())
	service.GET("/", HealthCheck())

	log.Println("🌠 Server starting at port :8080")
	router.Run(":8080")
}

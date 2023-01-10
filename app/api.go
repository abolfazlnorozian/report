package app

import (
	"fmt"
	"net/http"
	"os"
	"report/client"
	"report/helper/env"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

func RunServer() {
	mongoclient := client.MongoDB()
	redisclient := client.Redis()
	server := client.SERVER()
	ctx := client.Ctx()
	env.Load()
	port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "9000"
	// }

	defer mongoclient.Disconnect(ctx)
	value, err := redisclient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	fmt.Println(server.Run(":" + port))

}

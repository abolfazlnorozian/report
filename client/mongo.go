package client

import (
	"context"
	"fmt"
	"os"
	"report/helper/env"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client
)

func init() {
	env.Load()
	ctx = context.TODO()
	MGN := os.Getenv("MONGODB_URI")

	mongoconn := options.Client().ApplyURI(MGN)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	server = gin.Default()
}
func MongoDB() *mongo.Client {
	return mongoclient

}
func SERVER() *gin.Engine {
	return server
}
func Ctx() context.Context {
	return ctx
}

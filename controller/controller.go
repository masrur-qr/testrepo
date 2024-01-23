package controller

import (
	"context"
	"docker/app/structs"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DataList(c *gin.Context) {
	var SliceToSend []structs.User

	MongoConnection()
	client := MongoConnection()
	Connection := client.Database("DockerSession").Collection("Users")

	Result, err := Connection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for Result.Next(context.TODO()) {
			var UserTemp structs.User
			Result.Decode(&UserTemp)

			SliceToSend = append(SliceToSend, UserTemp)
		}
		Result.Close(context.TODO())

		c.JSON(200,SliceToSend)
	}

}
func DataPost(c *gin.Context) {
	var UserTemp structs.User
	c.ShouldBindJSON(&UserTemp)

	if UserTemp.Login == "" || UserTemp.Name == "" || UserTemp.Password == "" {
		c.JSON(404, "Emty field")
	} else {
		MongoConnection()
		client := MongoConnection()
		Connection := client.Database("DockerSession").Collection("Users")

		result, err := Connection.InsertOne(context.TODO(), UserTemp)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("result: %v\n", result)
			c.JSON(200, "Success")
		}
	}

}

func MongoConnection() *mongo.Client {
	var optionsUrl = options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	var client, err = mongo.Connect(context.Background(), optionsUrl)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return client

}

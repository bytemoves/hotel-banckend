package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/bytemoves/hotel-backend/api"
	"github.com/bytemoves/hotel-backend/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"



func main () {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "Goku",
		LastName: "Saitama",
	}


	res ,err := coll.InsertOne(ctx,user)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(res)

	listenAddr := flag.String("listenAddr",":5000","The listen addr of the api server")
	flag.Parse()


	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	
	apiv1.Get("/user",api.HandleGetUsers)
	apiv1.Get("/user:id",api.HandleGetUser)

	app.Listen(*listenAddr)
}





package main

import (
	"context"
	"flag"
	
	"log"

	"github.com/bytemoves/hotel-backend/api"
	"github.com/bytemoves/hotel-backend/db"
	// "github.com/bytemoves/hotel-backend/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)



//const userColl = "users"



// Create a new fiber instance with custom config
var config = fiber.Config{
    
    ErrorHandler: func(c *fiber.Ctx, err error) error {
      
        return c.JSON(map[string]string{"error": err.Error()})
    },
}
func main () {

	

	listenAddr := flag.String("listenAddr",":5000","The listen addr of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	
	//handler initialization

	var (

		userHandler = api.NewUserHandler(db.NewMongoUserStore(client,db.DBNAME))
		hotelStore = db.NewMongoHotelStore(client)
		roomStore = db.NewMongoRoomStore(client,hotelStore)
		hotelHandler  = api.NewHotelHandler(hotelStore,roomStore)
		app = fiber.New(config)
		apiv1 = app.Group("/api/v1")


	)


	//userhandler
	apiv1.Delete("/user/:id",userHandler.HandlePutUser)
	apiv1.Delete("/user/:id",userHandler.HandleDeleteUser)
	apiv1.Post("/user",userHandler.HandlePostUser)
	apiv1.Get("/user",userHandler.HandleGetUser)
	apiv1.Get("/user/:id",userHandler.HandleGetUser)


	///hotel Handlers

	apiv1.Get("/hotel",hotelHandler.HandleGetHotels)

	app.Listen(*listenAddr)
}





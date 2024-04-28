package main

import (
	"context"
	"flag"

	"log"

	"github.com/bytemoves/hotel-backend/api"
	"github.com/bytemoves/hotel-backend/api/middleware"
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

		//userHandler = api.NewUserHandler(db.NewMongoUserStore(client))
		hotelStore = db.NewMongoHotelStore(client)
		roomStore = db.NewMongoRoomStore(client,hotelStore)
		userStore = db.NewMongoUserStore (client)
		store = &db.Store{
			Hotel: hotelStore,
			Room: roomStore,
			User: userStore,

		}
		userHandler = api.NewUserHandler(userStore)
		hotelHandler  = api.NewHotelHandler(store)

		authHandler = api.NewAuthHandler(userStore)
		app = fiber.New(config)
		auth = app.Group("/api")
		apiv1 = app.Group("/api/v1",middleware.JWTAuthentication)


	)
	///auth

	auth.Post("/auth/",authHandler.HandleAuthenticate)

	//Versioned API routes

	//userhandler
	apiv1.Delete("/user/:id",userHandler.HandlePutUser)
	apiv1.Delete("/user/:id",userHandler.HandleDeleteUser)
	apiv1.Post("/user",userHandler.HandlePostUser)
	apiv1.Get("/user",userHandler.HandleGetUser)
	apiv1.Get("/user/:id",userHandler.HandleGetUser)


	///hotel Handlers

	apiv1.Get("/hotel",hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id",hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms",hotelHandler.HotelGetRooms)

	app.Listen(*listenAddr)
}





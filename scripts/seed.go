package main

//insert hotel to db go to store and make a hotel and room store
import (
	"context"
	
	"log"

	"github.com/bytemoves/hotel-backend/db"
	"github.com/bytemoves/hotel-backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	roomStore db.RoomStore
	hotelStore db.HotelStore

	ctx = context.Background()

)

func seedHotel(name string , location string , rating int){


	hotel := types.Hotel{
		Name: name,
		Location: location,
		Rooms: []primitive.ObjectID{},
		Rating: rating,
	}


	rooms := []types.Room{
		{
			Type: types.SeasideRoomType,
			BasePrice: 99.9,
		},

		{
			Type: types.DeluxeRoomType,
			BasePrice: 199.9,

		},

		{
			Type: types.SinglePersonRomType,
			BasePrice: 19.9,
		},
	}


	insertedHotel , err := hotelStore.InsertHotel(ctx,&hotel)
	if err != nil{
		log.Fatal(err)
	}

	for _,room  := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil{
		log.Fatal(err)
	}

	}
	

}


func main () {

	seedHotel("Bellucia","France" , 3)
	seedHotel("The cozy hotel","Netheralnds",4)
	seedHotel("halal","London",1)
	



	

}

func init() {
	var err error

	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}


	if err := client.Database(db.DBNAME).Drop(ctx); err!= nil{
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client,hotelStore)

}
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
	userStore db.UserStore

	ctx = context.Background()

)


func seedUser(fname,lname,email string) {
	user , err := types.NewUserFromParams(types.CreateUserParams{
		Email: email,
		FirstName: fname,
		LastName: lname,
		Password: "securepassfornow",
	})

	if err != nil{
		log.Fatal(err)
	}

	_,err = userStore.InsertUser(context.TODO(),user)
	if err != nil{
		log.Fatal(err)
	}


}


func seedHotel(name string , location string , rating int){


	hotel := types.Hotel{
		Name: name,
		Location: location,
		Rooms: []primitive.ObjectID{},
		Rating: rating,
	}


	rooms := []types.Room{
		{
			Size: "small",
			Price: 9.9,
		},

		{
			Size: "normal",
			Price: 19.9,

		},

		{
			Size: "kingsize",
			Price: 199.9,
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
	seedUser("goo","bar","foobar@gmail.com")
	



	

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
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)

	

	

}
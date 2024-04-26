package db

import (
	"context"

	"github.com/bytemoves/hotel-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type RoomStore interface{
	InsertRoom (context.Context, *types.Room) (*types.Room,error)
}

type MongoRoomStore struct{
	client *mongo.Client
	coll *mongo.Collection

	HotelStore
}

func NewMongoRoomStore (client *mongo.Client,hotelStore HotelStore) *MongoRoomStore{
	return &MongoRoomStore{
		client: client,
		coll: client.Database(DBNAME).Collection("Rooms"),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom (ctx context.Context, room *types.Room) (*types.Room, error) {
	resp , err := s.coll.InsertOne(ctx,room)

	if err != nil {
		return nil , err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)
	//update the hotel	
	filter := bson.M{"_id":room.HotelID}

	update := bson.M{"$psuh": bson.M{"rooms":room.ID}}

	if err := s.HotelStore.Update(ctx , filter , update); err != nil {
		return nil , err
	}





	return room , nil

}
package db

import (
	"context"
	// "go/build/constraint"

	"github.com/bytemoves/hotel-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/params"
	// "google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1/schema/predict/params"
	// "google.golang.org/genproto/googleapis/cloud/retail/v2"
)


const userColl = "users"





type UserStore interface{
	GetUserByID (context.Context,string) (*types.User,error)
	GetUsers(context.Context) ([]*types.User,error)
	InsertUser(context.Context,*types.User) (*types.User,error)
	DeleteUser(context.Context,string) error
	UpdateUser(ctx context.Context,  filter bson.M, params types.UpdateUserParams) error
	
}

type MongoUserStore struct{
	client * mongo.Client
	
	coll *mongo.Collection
}

func NewMongoUserStore( client *mongo.Client) *MongoUserStore{
	
	return &MongoUserStore{
		client: client,
		coll : client.Database(DBNAME).Collection(userColl),
	}
}


func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, param types.UpdateUserParams) error{
	update := bson.D{
		{
		"$set",params.ToBSON(),
		},
	}

	_ , err := s.coll.UpdateOne(ctx,filter,update)
	if err != nil{
		return err

	}

	return nil

}

func (s *MongoUserStore) DeleteUser(ctx context.Context , id string) error {
	oid ,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}

	res,err := s.coll.DeleteOne(ctx,bson.M{"_id":oid})

	if err != nil {
		return err
	}
// TODO later
	if res.DeletedCount == 0 {
		return nil

	}
	return nil

}

func (s *MongoUserStore) InsertUser(ctx context.Context , user *types.User) (*types.User,error){
	res , err := s.coll.InsertOne(ctx,user)
	if err != nil{
		return nil,err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return user,nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User,error){
	 
	 cur , err := s.coll.Find(ctx,bson.M{})

	 if err != nil{
		return nil,err
	 }

	 var users []*types.User
	 if err   := cur.All(ctx,&users); err != nil{
		return nil,err
	 } 

	 return users, nil



}

func (s *MongoUserStore) GetUserByID(ctx context.Context,id string) (*types.User,error){
	
	oid ,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil,err
	}


	var user types.User
	if err := s.coll.FindOne(ctx,bson.M{"_id":oid}).Decode(&user); err!=nil{
		return nil,err
	}

	return &user,nil

}
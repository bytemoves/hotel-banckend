package api

import (
	"bytes"
	"context"
	"encoding/json"
	
	"log"
	"net/http/httptest"
	"testing"

	"github.com/bytemoves/hotel-backend/db"
	"github.com/bytemoves/hotel-backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

const (
	testdburi = "mongodb://localhost:27017"
	dbname    = "hotel-resrvation-test"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}

}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)

	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "foo@bar.com",
		FirstName: "foobar",
		LastName:  "pop",
		Password:  "lldkdjdjeoodf",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("context-type", "application/json")

	resp, err := app.Test(req)
	if err != nil{
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	if len (user.ID) == 0 {
		t.Errorf("expecting a user id to be set ")
	}
	if len (user.EnctyptedPassword) > 0 {
		t.Errorf("expecting the expected password not to be included in json response ")
	}
	if user.FirstName != params.FirstName{
		t.Errorf("Expected First name %s but got %s",params.FirstName,user.FirstName)
	}
	if user.LastName != params.LastName{
		t.Errorf("Expected Last name %s but got %s",params.LastName,user.LastName)
	}
	if user.Email != params.Email{
		t.Errorf("Expected Email %s but got %s",params.Email,user.Email)
	}
 
}

package api

import (
	"errors"
	// "path/filepath"

	"github.com/bytemoves/hotel-backend/db"
	"github.com/bytemoves/hotel-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "github.com/bytemoves/hotel-backend/types"
	"github.com/gofiber/fiber/v2"
)

// retrieve
type UserHandler struct {
	UserStore db.UserStore
}

//constructor function

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}




func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error{
	var (
		// values bson.E
		params types.UpdateUserParams
		userID = c.Params("id")
	)
	oid ,err := primitive.ObjectIDFromHex(userID)
	if err != nil{
		return err
	}
	if err := c.BodyParser(&params); err != nil{
		return err
	}
	filter := bson.M{"_id":oid}
	if err := h.UserStore.UpdateUser(c.Context(),filter,params); err != nil{
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}



func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error{
	userID := c.Params("id")
	if err := h.UserStore.DeleteUser(c.Context(),userID);err != nil{
		return err
	}
	return c.JSON(map[string]string{"deleted":userID})

}



func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil{
		return nil
	}
	if errors := params.Validate(); len(errors) > 0{
		return c.JSON(errors)
		
	}

	user ,err := types.NewUserFromParams(params)
	if err != nil{
		return nil
	}

	InsertedUser ,err := h.UserStore.InsertUser(c.Context(),user)
	if err != nil{
		return err
	}
	return c.JSON(InsertedUser)
}


func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		
	)

	user, err := h.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		
		if errors.Is(err,mongo.ErrNoDocuments){
			return c.JSON(map[string]string{"error": "not found"})
		}

		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users ,err := h.UserStore.GetUsers(c.Context())
	if err != nil{
		return err
	}
	return c.JSON(users)
	
}

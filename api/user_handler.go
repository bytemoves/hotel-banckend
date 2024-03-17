package api

import (
	"github.com/bytemoves/hotel-backend/db"
	"github.com/bytemoves/hotel-backend/types"
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

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil{
		return nil
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

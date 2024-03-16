package api

import (
	"github.com/bytemoves/hotel-backend/types"
	"github.com/gofiber/fiber/v2"
)


func HandleGetUsers(c *fiber.Ctx) error{
	u := types.User{
		FirstName: "Goku",
		LastName: "Saitama",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error{
	return c.JSON("Goku")
}
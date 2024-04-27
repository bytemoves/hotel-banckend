package api

import (
	"github.com/bytemoves/hotel-backend/db"
	"github.com/gofiber/fiber/v2"
)


type HotelHandler struct{
	hotelStore db.HotelStore
	roomStore db.RoomStore

}


func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hs,
		roomStore: rs,
	}


}
type HotelQUeryParams struct {
	Rooms bool
}
func (h *HotelHandler) HandleGetHotels ( c *fiber.Ctx) error {
	var qparams HotelQUeryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}

	
	hotels , err := h.hotelStore.GetHotels(c.Context(),nil)
	if err != nil {
		return err
	}

	return c.JSON((hotels))


}


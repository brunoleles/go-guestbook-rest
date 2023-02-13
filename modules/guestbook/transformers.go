package guestbook

import (
	"main/models"
	"main/support"
)

func TransformGuestbook(data models.GuestbookModel) support.TransformedMap {
	return support.TransformedMap{
		"id":      data.ID,
		"name":    data.Name,
		"message": data.Message,
	}
}

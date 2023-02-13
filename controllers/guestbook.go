package controllers

import (
	. "main/database"
	"main/models"
	"main/support"
	"net/http"

	"github.com/gin-gonic/gin"
)

type guestbookPostRequest struct {
	//models.GuestbookModel
	Name    string `json:"name" form:"name"`
	Message string `json:"message" form:"message"`
}

type guestbookDeleteRequest struct {
	ID string `json:"id" form:"id"`
}

func GuestbookGet(c *gin.Context) {
	entries := []models.GuestbookModel{}
	DB().Find(&entries)

	c.JSON(http.StatusOK, entries)
}

func GuestbookPost(c *gin.Context) {
	var request guestbookPostRequest
	c.Bind(&request)

	entry := models.GuestbookModel{}
	support.CopyInto(request, &entry)

	if r := DB().Create(&entry); r.Error != nil {
		support.GromErrorResponse(c, r.Error)
		return
	}

	c.JSON(http.StatusOK, support.ResponseOk{
		Data: support.TransformGuestbook(entry),
	})
}

func GuestbookDelete(c *gin.Context) {
	var request guestbookDeleteRequest
	c.Bind(&request)

	entry := models.GuestbookModel{}
	if r := DB().First(&entry, request.ID); r.Error != nil {
		support.GromErrorResponse(c, r.Error)
		return
	}

	if r := DB().Delete(&entry); r.Error != nil {
		support.GromErrorResponse(c, r.Error)
		return
	}

	c.JSON(http.StatusOK, support.ResponseOk{
		Data: true,
	})
}

package controllers

import (
	"main/models"
	"main/support"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GustbookPostRequest struct {
	//models.GuestbookModel
	Name    string `json:"name" form:"name"`
	Message string `json:"message" form:"message"`
}

type GustbookDeleteRequest struct {
	ID string `json:"id" form:"id"`
}

func GuestbookGet(c *gin.Context) {
	entries := []models.GuestbookModel{}
	support.GetDB().Find(&entries)

	c.JSON(http.StatusOK, entries)
}

func GuestbookPost(c *gin.Context) {
	var request GustbookPostRequest
	c.Bind(&request)

	entry := models.GuestbookModel{}
	support.CopyInto(request, &entry)

	if r := support.GetDB().Create(&entry); r.Error != nil {
		support.GromErrorResponse(c, r.Error)
		return
	}

	c.JSON(http.StatusOK, support.ResponseOk{
		Data: support.TransformGuestbook(entry),
	})
}

func GuestbookDelete(c *gin.Context) {
	var request GustbookDeleteRequest
	c.Bind(&request)

	entry := models.GuestbookModel{}
	if r := support.GetDB().First(&entry, request.ID); r.Error != nil {
		support.GromErrorResponse(c, r.Error)
		return
	}

	if r := support.GetDB().Delete(&entry); r.Error != nil {
		support.GromErrorResponse(c, r.Error)
		return
	}

	c.JSON(http.StatusOK, support.ResponseOk{
		Data: true,
	})
}

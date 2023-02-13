package support

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseOk struct {
	Data any `json:"data"`
}

type ResponseNok struct {
	Error   bool   `json:"error"`
	Message string `json:"message",omitempty`
}

func GromErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, ResponseNok{
		Error:   true,
		Message: err.Error(),
	})
}

type TransformedMap map[string]interface{}

func TransformGuestbook(data models.GuestbookModel) TransformedMap {
	return TransformedMap{
		"id":      data.ID,
		"name":    data.Name,
		"message": data.Message,
	}
}

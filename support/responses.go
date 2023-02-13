package support

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseOk struct {
	Data any `json:"data"`
}

type ResponseNok struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
}

func GromErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, ResponseNok{
		Error:   true,
		Message: err.Error(),
	})
}

type TransformedMap map[string]interface{}

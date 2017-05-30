package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	//"strings"
)

//This function gets error key and returns appropriate message regarding to this key
func HandleHttpError(c *gin.Context, err error) {

	// all errors are internal unless equal specified errors or have Error structure and NotValid/BadRequest Key
	statusCode := http.StatusInternalServerError
	if err == ErrNotFound {
		statusCode = http.StatusNotFound
	} else if err == ErrForbidden || err == ErrTokenInvalid {
		statusCode = http.StatusForbidden
	} else if err == ErrObjectInvalid || err == ErrPassRequired || err == ErrPassword ||
		err == ErrPassFormat || err == ErrRateExceed || err == ErrAgentNotActive {
		statusCode = http.StatusBadRequest

	} else if err == ErrDuplicate || strings.Contains(err.Error(), Messages["DuplicateIndex"]) {
		statusCode = http.StatusConflict
	}

	if customError, ok := err.(Error); ok {
		if customError.Key == "NotValid" || customError.Key == "BadRequest" {
			statusCode = http.StatusBadRequest
		}
	}

	c.JSON(statusCode, err.Error())
	c.Abort()
	return

}

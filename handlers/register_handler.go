package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/simpleton/tinker-api/models"
	"github.com/simpleton/tinker-api/db"
	"github.com/twinj/uuid"
***REMOVED***

type RegisterHandler struct {

}

func(this *RegisterHandler***REMOVED*** Register(c echo.Context***REMOVED*** error {
	register := new(models.RegisterInfo***REMOVED***
	response := models.NewJsend(***REMOVED***
	if err := c.Bind(register***REMOVED***; err != nil {
		c.JSON(http.StatusBadRequest, response.StatusCode(http.StatusBadRequest***REMOVED***.Message(err.Error(***REMOVED******REMOVED******REMOVED***
	}
	// check user existed
	if existed, err := db.CheckEmailExisted(register.Email***REMOVED***; err != nil {
		c.JSON(http.StatusBadRequest, response.StatusCode(http.StatusBadRequest***REMOVED***.Message(err.Error(***REMOVED******REMOVED******REMOVED***
	} ***REMOVED*** {
		if existed {
			return c.JSON(http.StatusBadRequest, response.StatusCode(http.StatusBadRequest***REMOVED***.Message(err.Error(***REMOVED******REMOVED******REMOVED***
		}
	}
	salt := uuid.NewV4(***REMOVED***
	db.CreateUser(register.Username, register.Password, register.Email, salt.String(***REMOVED******REMOVED***
	return c.JSON(http.StatusCreated, register***REMOVED***
}

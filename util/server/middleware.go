package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"postinger/util/localization"
	"postinger/util/logwrapper"
	"postinger/util/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// LanHeader - is for headers without auth
type LanHeader struct {
	Language       string `header:"language" binding:"required"`
	Platform       int    `header:"platform" binding:"required"`
	CurrencySymbol string `header:"currencysymbol" binding:"required"`
	CurrencyCode   string `header:"currencycode" binding:"required"`
}

// AuthLanHeader - is for headers with auth
type AuthLanHeader struct {
	Authorization  string `header:"authorization" binding:"required"`
	Language       string `header:"language" binding:"required"`
	Platform       int    `header:"platform" binding:"required"`
	CurrencySymbol string `header:"currencysymbol" binding:"required"`
	CurrencyCode   string `header:"currencycode" binding:"required"`
}

// This middleware ensures that a request will be aborted with an error
// if the user is not logged in
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := AuthLanHeader{}
		lang := c.GetHeader("language")
		errResponse := models.ErrResponse{
			Message: localization.GetMessage(lang, "common.401", nil),
		}
		if err := c.ShouldBindHeader(&h); err != nil {
			logwrapper.Logger.Debugln("error in header to struct : ", err)
			var verr validator.ValidationErrors
			fields := []string{}
			if errors.As(err, &verr) {
				for _, f := range verr {
					fields = append(fields, f.Field())
				}
			}
			errResponse := models.ErrResponse{
				Message: localization.GetMessage(lang, "common.400", map[string]interface{}{
					"Fields": strings.Join(fields, ", "),
				}),
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
			return
		}
		// If there's an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("validUser")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		}
	}
}

// This middleware ensures that a request will be aborted with an error
// if the user is already logged in
func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := LanHeader{}
		lang := c.GetHeader("language")
		errResponse := models.ErrResponse{
			Message: localization.GetMessage(lang, "common.401", nil),
		}
		if err := c.ShouldBindHeader(&h); err != nil {
			logwrapper.Logger.Debugln("error in header to struct : ", err)
			var verr validator.ValidationErrors
			fields := []string{}
			if errors.As(err, &verr) {
				for _, f := range verr {
					fields = append(fields, f.Field())
				}
			}
			errResponse := models.ErrResponse{
				Message: localization.GetMessage(lang, "common.400", map[string]interface{}{
					"Fields": strings.Join(fields, ", "),
				}),
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
			return
		}
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("validUser")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
		}
	}
}

// This middleware sets whether the user is logged in or not
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Set("validUser", false)
		} else {
			logwrapper.Logger.Debugln("token : ", tokenString)
			var tokenData map[string]interface{}
			if json.Unmarshal([]byte(tokenString), &tokenData) == nil {
				// logwrapper.Logger.Debugln("tokenData : ", tokenData)
				c.Set("userData", tokenData)
				c.Set("validUser", true)
			} else {
				c.Set("validUser", false)
			}
		}
	}
}

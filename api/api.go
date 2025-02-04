package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetApiList
// api list
func GetApiList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"teapot":                               "https://api.hellozwz.com/teapot",
		"token by username and password":       "https://api.hellozwz.com/v1/token/username-pwd",
		"token by phone and password":          "https://api.hellozwz.com/v1/token/phone-pwd",
		"token by phone and verification code": "https://api.hellozwz.com/v1/token/phone-code",
		"users":                                "https://api.hellozwz.com/v1/users",
		"users_authority":                      "https://api.hellozwz.com/v1/users/authority",
		"users_password":                       "https://api.hellozwz.com/v1/users/password",
		"books":                                "https://api.hellozwz.com/v1/books",
		"file":                                 "https://api.hellozwz.com/v1/file",
	})
}

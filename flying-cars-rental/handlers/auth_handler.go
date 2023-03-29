package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)
var userDB *sql.DB 

type SignUpRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_password"`
}

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// SignUp function to handle sign up requests
func SignUp(c *gin.Context) {
	var req SignUpRequest
	var userDB *sql.DB 
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Check if user with this mobile number already exists
	if _, ok := userDB[req.Mobile]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	// Validate password
	if req.Password != req.ConfirmPass {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
		return
	}

	// Create new user
	user := &User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Mobile:    req.Mobile,
		Password:  req.Password,
	}
	userDB[req.Mobile] = user

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

// Login function to handle login requests
func Login(c *gin.Context) {
	var req LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Check if user with this mobile number exists
	user, ok := userDB[req.Mobile]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		return
	}

	// Validate password
	if user.Password != req.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate session token
	token, err := GenerateSessionToken(user.Mobile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Set session cookie
	c.SetCookie("session_token", token, cookieMaxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

// Logout function to handle logout requests
func Logout(c *gin.Context) {
	// Delete session cookie
	c.SetCookie("session_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// SessionCheck function to check if user session is valid
func SessionCheck(c *gin.Context) {
	// Get session cookie
	cookie, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session not found"})
		return
	}

	// Verify session token
	mobile, err := VerifySessionToken(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
		return
	}

	// Get user from database
	user, ok := userDB[mobile]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "session valid", "user": user})
}
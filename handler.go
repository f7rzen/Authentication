package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func LoginUser(c *gin.Context) {
	var loginData User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	email := loginData.Email
	user, err := GetUser(db, email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

func RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body"})
		return
	}

	err := CreateUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name})
}

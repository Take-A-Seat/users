package main

import (
	"github.com/Take-A-Seat/storage/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func handleCreateUser(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	 if err :=addUser(user); err !=nil{
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	 }else{
		 c.JSON(http.StatusCreated, gin.H{"error": "Success create user"})

	 }

}

func handleValidateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := getUserByParam("email", user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusNoContent, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": response.Id.Hex(), "firstName": response.FirstName,
		"lastName": response.LastName})

}

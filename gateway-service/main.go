package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type password struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Strength int    `json:"strength"`
}

// temporary collection that replaces db
var passwords = []password{
	{ID: 1, Password: "intel1", Strength: 0},
	{ID: 2, Password: "elyass15@ajilent-ci", Strength: 2},
	{ID: 3, Password: "hodygid757#$!23w", Strength: 1},
}

func main() {
	router := gin.Default()
	router.GET("/api/v1/passwords", getPasswords)
	router.POST("/api/v1/passwords", postPassword)

	router.Run("localhost:10000")
}

func getPasswords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, passwords)
}

func postPassword(c *gin.Context) {
	var newPassword password
	if err := c.BindJSON(&newPassword); err != nil {
		return
	}
	newPassword.ID = passwords[len(passwords)-1].ID + 1
	passwords = append(passwords, newPassword)
	c.IndentedJSON(http.StatusCreated, newPassword)
}

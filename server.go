package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/chats", getUsers)
	router.GET("/users", getUsers)
	router.GET("/messages", getMessages)
	router.POST("/chats", setChats)
	router.POST("/users", setUsers)
	router.POST("/messages", sendMessages)
	router.Run(":1234")
}

func getUsers(c *gin.Context) {
	c.String(200, "get users!")
}

func getChats(c *gin.Context) {
	c.String(200, "get chats!")
}
func getMessages(c *gin.Context) {
	c.String(200, "get ,messages!")
}
func setUsers(c *gin.Context) {
	c.String(200, "set users!")
}

func setChats(c *gin.Context) {
	c.String(200, "set chats!")
}

func sendMessages(c *gin.Context) {
	id := c.Query("id")
	senderId := c.Query("sender")
	message := c.Query("message")

	fmt.Printf("Message sent with id %s and senderId %s and message $s", id, senderId, message)
}

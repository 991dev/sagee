package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/chats", getUsers)
	router.GET("/users", getUsers)
	router.POST("/chats", setChats)
	router.POST("/users", setUsers)
	router.Run()
}

func getUsers(c *gin.Context) {
	c.String(200, "get users!")
}

func getChats(c *gin.Context) {
	c.String(200, "get chats!")
}

func setUsers(c *gin.Context) {
	c.String(200, "set users!")
}

func setChats(c *gin.Context) {
	c.String(200, "set chats!")
}

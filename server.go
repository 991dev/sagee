package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	const (
		host     = "89.23.118.97"
		port     = 5432
		user     = "creator991"
		password = "3228103228A"
		dbname   = "default_db"
	)

	// Connect to the PostgreSQL database
	connString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("Connection String:", connString)

	var errConnect error
	db, errConnect = sql.Open("postgres", connString)
	if errConnect != nil {
		log.Fatal(errConnect)
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal(errPing)
	}
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", setUsers)

	router.GET("/chats", getChats)
	router.POST("/chats", setChats)

	router.GET("/messages", getMessages)
	router.POST("/messages", sendMessages)

	router.Run()
}

func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error querying users from the database: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			log.Printf("Error scanning users from the database: %v", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
		users = append(users, user)
	}

	log.Printf("Users: %+v", users) // Print users for debugging

	c.JSON(200, users)
}

func setUsers(c *gin.Context) {
	username := c.PostForm("username")

	_, err := db.Exec("INSERT INTO users(username) VALUES($1)", username)
	if err != nil {
		c.String(500, "Error inserting user into the database")
		return
	}

	c.String(200, "User inserted successfully")
}

func getChats(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM chats")
	if err != nil {
		log.Printf("Error querying chats from the database: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var chats []string
	for rows.Next() {
		var chat string
		if err := rows.Scan(&chat); err != nil {
			log.Printf("Error scanning chats from the database: %v", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
		chats = append(chats, chat)
	}

	c.JSON(200, chats)
}

func setChats(c *gin.Context) {
	chatName := c.PostForm("chatName")

	_, err := db.Exec("INSERT INTO chats(chatName) VALUES($1)", chatName)
	if err != nil {
		c.String(500, "Error inserting chat into the database")
		return
	}

	c.String(200, "Chat inserted successfully")
}

func getMessages(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM messages")
	if err != nil {
		log.Printf("Error querying messages from the database: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			log.Printf("Error scanning messages from the database: %v", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
		messages = append(messages, message)
	}

	c.JSON(200, messages)
}

func sendMessages(c *gin.Context) {
	senderID := c.PostForm("senderID")
	messageText := c.PostForm("messageText")
	chatID := c.PostForm("chatID")

	_, err := db.Exec("INSERT INTO messages(sender_id, message_text, chat_id) VALUES($1, $2, $3)", senderID, messageText, chatID)
	if err != nil {
		c.String(500, "Error sending message to the database")
		return
	}

	c.String(200, "Message sent successfully")
}

func cleanup() {
	db.Close()
}

func init() {
	defer cleanup()
}

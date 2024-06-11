package cmd

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateMessage(c *gin.Context, db *sql.DB) {
	//Parse JSON request into Message struct
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Insert message into database
	result, err := db.Exec("INSERT INTO messages (channel_id, user_id, message) VALUES (?, ?, ?)", message.ChannelID, message.UserID, message.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//GET ID of newly inserted message
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return id of newly created message
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func ListMessages(c *gin.Context, db *sql.DB) {
	//parse channel ID from URL
	channelID, err := strconv.Atoi(c.Query("channelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Parse optional limit query parameter  from URL
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		//Set limit to 100 if not provided
		limit = 100
	}

	//Parse last message ID query parameter from URL. This is used to get messages after a certain message
	lastMessageID, err := strconv.Atoi(c.Query("lastMessageID"))
	if err != nil {
		//Set last message ID to 0 if not provided
		lastMessageID = 0
	}
	//Query database for messages
	rows, err := db.Query("SELECT m.id, channel_id, user_id, u.username AS user_name, message FROM messages m LEFT JOIN users u ON u.id = m.user_id WHERE channel_id = ? AND m.id > ? ORDER BY m.id ASC LIMIT ?", channelID, lastMessageID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Create slice of messages
	var messages []Message

	//Iterate over rows
	for rows.Next() {
		//Create new message
		var message Message

		//Scan row into message
		err := rows.Scan(&message.ID, &message.ChannelID, &message.UserID, &message.UserName, &message.Text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//Append message into slice
		messages = append(messages, message)
	}

	//Return slice of messages
	c.JSON(http.StatusOK, messages)
}

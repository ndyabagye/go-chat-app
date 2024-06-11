package cmd

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateChannel endpoint
func CreateChannel(c *gin.Context, db *sql.DB) {
	//Parse JSON request into channel struct
	var channel Channel
	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Insert channel into database
	result, err := db.Exec("INSERT INTO channels (name) VALUES (?)", channel.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Get ID of newly inserted channel
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//return ID of newly created channel
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func ListChannels(c *gin.Context, db *sql.DB) {
	//query database for channels
	rows, err := db.Query("SELECT id, name FROM channels")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Create a slice of channels
	var channels []Channel

	//Iterate over channels
	for rows.Next() {
		//Create new channel
		var channel Channel

		//Scan row into channel
		err := rows.Scan(&channel.ID, &channel.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//append channel into slice
		channels = append(channels, channel)

		//return slice of channels
		c.JSON(http.StatusOK, channels)
	}
}

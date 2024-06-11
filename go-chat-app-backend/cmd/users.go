package cmd

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context, db *sql.DB) {
	//parse JSON request body into User struct
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Query database for user
	row := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", user.Username, user.Password)

	//Get ID of user
	var id int
	err := row.Scan(&id)
	if err != nil {
		//Check if user was not found
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		//return error if other error occured
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	//return ID of user
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func CreateUser(c *gin.Context, db *sql.DB) {
	//parse json request into User struct
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Insert user into database
	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Get ID of newly created user
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//return id of newly created user
	c.JSON(http.StatusOK, gin.H{"id": id})
}

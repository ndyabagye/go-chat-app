package main

import (
	"database/sql"
	"fmt"
	"go-chat-app-backend/cmd"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	//get the working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//Print the working directory
	fmt.Println("Working directory: ", wd)

	//Open the SQLite database file
	db, err := sql.Open("sqlite", wd+"/database.db")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	//Create the GIN router
	r := gin.Default()
	if err != nil {
		log.Fatal(err)
	}

	//Creation endpoints
	r.POST("/users", func(c *gin.Context) { cmd.CreateUser(c, db) })
	r.POST("/channels", func(c *gin.Context) { cmd.CreateChannel(c, db) })
	r.POST("/messages", func(c *gin.Context) { cmd.CreateMessage(c, db) })
	//Listing endpoints
	r.GET("/channels", func(c *gin.Context) { cmd.ListChannels(c, db) })
	r.GET("/messages", func(c *gin.Context) { cmd.ListMessages(c, db) })
	//Login endpoint
	r.POST("/login", func(c *gin.Context) { cmd.Login(c, db) })

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

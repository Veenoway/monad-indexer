package main

import (
	"context"
	"monad-indexer/internal/db"

	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	defer db.Conn.Close(context.Background())

	db.Migrate()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bakend en GO"})
	})

	r.Run(":8080")
}






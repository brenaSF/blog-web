package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	// Define a route for serving CSS file
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	fmt.Println("Servindo na porta 8080...")
	r.Run(":8080") // Use Gin's Run function to start the server
}

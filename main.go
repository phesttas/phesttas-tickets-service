package main

import "fmt"
import "github.com/gin-gonic/gin"

func handleTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"I'm": "here!",
	})
}

func main() {
	fmt.Println("Starts here!")
	r := gin.Default()
	r.GET("/", handleTest)
	r.Run()
}

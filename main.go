package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the Go Application!")

	r := gin.New()
	r.GET("/", GetMessage)
	r.POST("/", PostMessage)
	r.GET("/query", GetFromQuery)
	r.GET("/path/:name/:age", GetFromPath)
	r.Run()
}

// GetMessage Gets a simple message.
func GetMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

// PostMessage Accepts a body of any type and returns it.
func PostMessage(c *gin.Context) {
	body := c.Request.Body

	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

// GetFromQuery Parses content from query strings.
func GetFromQuery(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// GetFromPath Parses content from path parameters
func GetFromPath(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

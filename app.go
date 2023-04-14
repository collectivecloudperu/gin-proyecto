package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/miruta", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Mega Crack de la Programación !",
		})
	})
	r.Run()
}

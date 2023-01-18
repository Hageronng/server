package main

import "github.com/gin-gonic/gin"

type Reg struct {
	Email    string
	Name     string
	Password string
}

type LogIn struct {
	Email    string
	Password string
}

func main() {
	server := gin.Default()

	var r Reg

	server.POST("/registration", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    r.Email,
			"Name":     r.Name,
			"Password": r.Password,
		})
	})

	server.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    r.Email,
			"Password": r.Password,
		})
	})

	server.Run(":0808")
}

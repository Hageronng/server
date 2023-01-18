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
	var l LogIn

	server.POST("/registration", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    r.Email,
			"Name":     r.Name,
			"Password": r.Password,
		})
		r.Email = l.Email
		r.Password = l.Password
	})

	server.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    l.Email,
			"Password": l.Password,
		})

		if l.Email == r.Email && l.Password == r.Password {
			server.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"LogIn": "Ok",
				})
			})
		} else {
			server.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"LogIn": "Fail",
				})
			})
		}
	})

	server.Run(":0808")
}

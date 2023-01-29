package controllers

import (
	"bufio"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Reg struct {
	Email    string
	Password string
}

type LogIn struct {
	Email    string
	Password string
}

func controllers() {
	Server := gin.Default()

	var r Reg
	var l LogIn

	Server.POST("/registration", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    r.Email,
			"Password": r.Password,
		})

		var acc string = r.Email + " - " + r.Password + "\n"

		path, _ := os.Open("server/accounts.dat")
		path.WriteString(acc)
		path.Close()
	})

	Server.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    l.Email,
			"Password": l.Password,
		})

		var acc string = l.Email + " - " + l.Password

		path, _ := os.Open("server/accounts.dat")

		scanner := bufio.NewScanner(path)

		line := 1
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), acc) {
				Server.GET("/", func(ctx *gin.Context) {
					ctx.JSON(200, gin.H{
						"LogIn": "Ok",
					})
				})
			} else {
				Server.GET("/", func(ctx *gin.Context) {
					ctx.JSON(200, gin.H{
						"LogIn": "Fail",
					})
				})
				line++
			}
		}

		path.Close()
	})
}

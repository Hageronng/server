package login

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type LogIn struct {
	Email    string
	Password string
}

func login() {
	Server := gin.Default()

	var e LogIn
	var l LogIn

	Server.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    e.Email,
			"Password": e.Password,
		})

		db, err := sql.Open("sqlite3", "data.db")
		if err != nil {
			panic(err)
		}

		login := db.QueryRow("select * from logins where email = $1", e.Email)
		err = login.Scan(&l.Email.email, &l.Password.password)
		if err != nil {
			Server.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"LogIn": "Not found",
				})
			})
		}

		if VerifyPassword(l.Password.password, e.Password) == nil {
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
		}

		login.Close()
		db.Close()
	})
}

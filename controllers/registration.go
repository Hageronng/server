package registration

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Reg struct {
	Email    string
	Password string
}

func registration() {
	Server := gin.Default()

	var r Reg

	Server.POST("/registration", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Email":    r.Email,
			"Password": r.Password,
		})

		HashPassword(r.Password)

		db, err := sql.Open("sqlite3", "data.db")
		if err != nil {
			panic(err)
		}

		db.Exec("insert into logins (email, password) values ($1, $2)", r.Email, hashedPassword)
		db.Close()
	})
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Server := gin.Default()

	Server.Run(":0808")
}

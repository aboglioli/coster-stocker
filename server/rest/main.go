package rest

import (
	"fmt"

	"github.com/aboglioli/coster-stocker/helpers/config"
	"github.com/gin-gonic/gin"
)

func Start() {
	c := config.Get()

	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/stock/:id", getStock)
	r.POST("/stock", createStock)

	r.Run(fmt.Sprintf(":%d", c.Port))
}

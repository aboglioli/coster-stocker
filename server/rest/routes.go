package rest

import (
	"log"
	"net/http"

	"github.com/aboglioli/coster-stocker/stock"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func createStock(c *gin.Context) {
	body := &stock.CreateRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		log.Println(err)
		return
	}

	stockService, err := stock.NewService()
	if err != nil {
		log.Println(err)
		return
	}

	stock, err := stockService.Create(body)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   stock.ID,
		"name": stock.Name,
	})
}

func getStock(c *gin.Context) {
	id := c.Param("id")

	stockService, err := stock.NewService()
	if err != nil {
		log.Println(err)
		return
	}

	stock, err := stockService.Get(id)

	c.JSON(http.StatusOK, gin.H{
		"id":   stock.ID,
		"name": stock.Name,
	})
}

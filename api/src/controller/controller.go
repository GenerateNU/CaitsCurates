package controller

import (
	"CaitsCurates/backend/src/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Serve() *gin.Engine
}

type PgController struct {
	model.Model
}

func (pg *PgController) Serve() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/gifts/:giftId", func(c *gin.Context) {
		id := c.Param("giftId")
		intId, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, pg.GetExampleGift(int64(intId)))
	})
	r.GET("/gifts", func(c *gin.Context) {
		gifts, err := pg.AllExampleGifts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})

	r.POST("/addGift", func(c *gin.Context) {
		var eg model.ExampleGift
		fmt.Print(c)
		if err := c.BindJSON(&eg); err != nil {
			c.JSON(http.StatusBadRequest,  "Failed to unmarshal gift")

			fmt.Print(err)

			return
		}
		insertedGift, err := pg.AddExampleGift(eg)

		if err != nil {
			c.JSON(http.StatusBadRequest, eg.GiftId)
			panic(err)
		}

		c.JSON(http.StatusOK, insertedGift)
	})

	return r
}

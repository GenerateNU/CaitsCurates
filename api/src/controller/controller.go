package controller

import (
	"CaitsCurates/backend/src/model"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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

	// Get incomplete gift requests
	r.GET("/requests/incomplete", func(c *gin.Context) {
		gifts, err := pg.IncompleteRequests()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})

	// Get complete gift requests
	r.GET("/requests/complete", func(c *gin.Context) {
		gifts, err := pg.CompleteRequests()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})

	r.POST("/addGiftResponse", func(c *gin.Context) {
		var input model.GiftResponse
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal respone")

			fmt.Print(err)

			return
		}
		insertedResponse, err := pg.AddResponse(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, insertedResponse)
	})
	r.POST("/addGiftRequest", func(c *gin.Context) {
		var input model.GiftRequest
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal request")
			fmt.Print(err)
			return
		}
		insertedRequest, err := pg.AddRequest(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, insertedRequest)
	})
	r.POST("/addGiftCollection", func(c *gin.Context) {
		var input model.GiftCollection
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal collection")

			fmt.Print(err)

			return
		}
		insertedCollection, err := pg.AddCollection(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, insertedCollection)
	})

	return r
}

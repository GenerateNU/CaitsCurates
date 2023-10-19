package controller

import (
	"CaitsCurates/backend/src/model"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	r.PUT("/requests", func(c *gin.Context) {
		// Get Body Parameters and put in JSON Object
		var input model.GiftRequest
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal gift")
			fmt.Print(err)
			return
		}

		// Model Call to Update GiftRequest
		updatedGiftRequest, err := pg.UpdateGiftRequest(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, updatedGiftRequest)
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
	r.PUT("/updateGiftCollection", func(c *gin.Context) {
		var input model.GiftCollection
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal collection")

			fmt.Print(err)

			return
		}
		updatedCollection, err := pg.UpdateCollection(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, updatedCollection)
	})
	r.GET("/gifts/:id", func(c *gin.Context) {
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		gift, err := pg.GetGift(int64(intId))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gift)
	})

	r.GET("/gifts", func(c *gin.Context) {
		gifts, err := pg.GetAllGifts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})
	r.GET("/responses", func(c *gin.Context) {
		responses, err := pg.AllGiftResponses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, responses)
	})
	r.GET("/collections", func(c *gin.Context) {
		collections, err := pg.AllCollections()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, collections)
	})
	r.POST("/addGift", func(c *gin.Context) {
		var input model.Gift
		fmt.Print(c)
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal gift")

			fmt.Print(err)

			return
		}
		insertedGift, err := pg.AddGift(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, insertedGift)
	})
	// Update Gift Record Given Gift ID
	r.PUT("/gifts/:id", func(c *gin.Context) {

		// Get Gift ID
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		// Get Body Parameters and put in JSON Object
		var input model.Gift
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal gift")
			fmt.Print(err)
			return
		}

		// Model Call to Update Gift
		updatedGift, err := pg.UpdateGift(int64(intId), input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, updatedGift)
	})

	// Delete Gift Record based on Gift ID
	r.DELETE("/gifts/:id", func(c *gin.Context) {

		// Get Gift ID
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		err = pg.DeleteGift(int64(intId))

		// Delete Gift
		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to delete gift")
			fmt.Print(err)
			return
		}
		c.JSON(http.StatusNoContent, "Deleted Gift")
	})
	r.DELETE("/deleteGiftCollection/:id", func(c *gin.Context) {

		// Get GiftCollection ID
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		err = pg.DeleteGiftCollection(int64(intId))

		// Delete GiftCollection
		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to delete GiftCollection")
			fmt.Print(err)
			return
		}
		c.JSON(http.StatusNoContent, "Deleted GiftCollection")
	})

	// Add Gift to Gift Collection
	r.POST("/addGiftCollection/:id", func(c *gin.Context) {
		var input model.Gift

		// Get Gift Collection Id
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal collection")
			fmt.Print(err)
			return
		}

		giftAddedCollection, err := pg.AddGiftToGiftCollection(input, int64(intId))

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, giftAddedCollection)
	})

	// Delete Gift to Gift Collection
	r.DELETE("/removeGiftFromGiftCollection/:giftID/:giftCollectionID", func(c *gin.Context) {
		var input model.Gift

		// Get Gift Collection Id
		collectionID, err := strconv.Atoi(c.Param("giftCollectionID"))
		if err != nil {
			panic(err)
		}

		giftID, err := strconv.Atoi(c.Param("giftID"))
		if err != nil {
			panic(err)
		}

		giftRemovedCollection, err := pg.DeleteGiftFromGiftCollection(int64(giftID), int64(collectionID))

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, giftRemovedCollection)
	})

	return r
}

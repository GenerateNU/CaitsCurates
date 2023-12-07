package controller

import (
	"CaitsCurates/backend/src/model"
	"fmt"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"

	"log"
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

func (pg *PgController) CreateStripeCheckoutSession(c *gin.Context) {
	print(stripe.Key)
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1OEIkBLQbsCsABA6PNVghEsC"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success"),
		CancelURL:  stripe.String(domain + "/cancel"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": s.URL})
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
	r.POST("/create-checkout-session", pg.CreateStripeCheckoutSession)

	// Get complete gift requests
	r.GET("/requests/complete", func(c *gin.Context) {
		gifts, err := pg.CompleteRequests()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})

	// Get customer gift requests
	r.GET("/requests/:id", func(c *gin.Context) {
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		requests, err := pg.GetCustomerRequests(int64(intId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, requests)
	})

	// Add a Gift Response
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
	// Update the Gift Request
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
	// Create a new Gift Request
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
	// Create a new Gift Collection
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
	// Update the Gift Collection
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
	// Get the Gift given the Gift ID
	r.GET("/gifts/:id", func(c *gin.Context) {
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		gift, err := pg.GetGift(int64(intId))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gift)
	})
	// Get all Gifts
	r.GET("/gifts", func(c *gin.Context) {
		gifts, err := pg.GetAllGifts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})
	// Get all Gift Responses
	r.GET("/responses", func(c *gin.Context) {
		responses, err := pg.AllGiftResponses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, responses)
	})
	// Get all the Gifts in a Gift Collection given filter options
	r.GET("/search/:giftCollectionId", func(c *gin.Context) {
		searchTerm := c.Query("q")
		minPriceStr := c.Query("minPrice")
		maxPriceStr := c.Query("maxPrice")
		occasion := c.Query("occasion")
		demographic := c.Query("demographic")
		category := c.Query("category")

		id := c.Param("giftCollectionId")
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid giftCollectionId")
			return
		}

		minPrice, _ := strconv.Atoi(minPriceStr)
		maxPrice, _ := strconv.Atoi(maxPriceStr)
		gifts, err := pg.SearchGifts(int64(intId), searchTerm, minPrice, maxPrice, occasion, demographic, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, gifts)
	})
	// Get all Gift Collections
	r.GET("/collections", func(c *gin.Context) {
		collections, err := pg.AllCollections()
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, collections)
	})
	// Create an endpoint that takes in a customerID and returns all collections with no customerID or a matching customerID.
	r.GET("/collections/:customerId", func(c *gin.Context) {

		// Get Customer ID
		id := c.Param("customerId")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		collections, err := pg.AllCustomerCollections(int64(intId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Oops")
		}
		c.JSON(http.StatusOK, collections)
	})
	// Create a new Gift
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
	// Delete Gift Collection based on Gift Collection ID
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

	// Add gift to gift collection for given customer id and collection name
	r.POST("/addCustomerGiftCollection/:collectionName/:customerId", func(c *gin.Context) {
		var input model.Gift

		collectionName := c.Param("collectionName")
		customerId := c.Param("customerId")

		intId, err := strconv.Atoi(customerId)
		if err != nil {
			panic(err)
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal gift")
			fmt.Print(err)
			return
		}

		giftAddedCollection, err := pg.AddGiftToCustomerCollection(input, collectionName, int64(intId))

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, giftAddedCollection)
	})
	// Remove a Customer from a Gift Collection given Gift Collection Name and Customer ID
	r.POST("/removeCustomerGiftCollection/:collectionName/:customerId", func(c *gin.Context) {
		var input model.Gift

		collectionName := c.Param("collectionName")
		customerId := c.Param("customerId")

		intId, err := strconv.Atoi(customerId)
		if err != nil {
			panic(err)
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal gift")
			fmt.Print(err)
			return
		}

		giftRemovedCollection, err := pg.DeleteGiftFromCustomerCollection(input, collectionName, int64(intId))

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, giftRemovedCollection)
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

	// Retrieve Giftees based on Giftee ID
	r.GET("/giftee/:id", func(c *gin.Context) {

		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		giftee, err := pg.GetGiftee(int64(intId))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, giftee)
	})

	// Create a new Giftee
	r.POST("/addGiftee", func(c *gin.Context) {
		var input model.Giftee
		fmt.Print(c)
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal giftee")

			fmt.Print(err)

			return
		}
		insertedGiftee, err := pg.AddGiftee(input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, insertedGiftee)
	})

	// Update Giftee Information
	r.PUT("/giftee/:id", func(c *gin.Context) {

		// Get Giftee ID
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		var input model.Giftee
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, "Failed to unmarshal giftee")

			fmt.Print(err)

			return
		}

		updatedGiftee, err := pg.UpdateGiftee(int64(intId), input)

		if err != nil {
			c.JSON(http.StatusBadRequest, input)
			panic(err)
		}

		c.JSON(http.StatusOK, updatedGiftee)
	})

	// Delete Giftee based on Giftee ID
	r.DELETE("/giftee/:id", func(c *gin.Context) {

		// Get Giftee ID
		id := c.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		err = pg.DeleteGiftee(int64(intId))

		// Delete Giftee
		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to delete giftee")
			fmt.Print(err)
			return
		}
		c.JSON(http.StatusNoContent, "Deleted Giftee")
	})
	// Update AvailableRequests based on Customer ID
	r.PUT("customer/:id", func(c *gin.Context) {

		// Get Customer ID
		customerID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		// Get request amount
		updatedRequests := c.Query("requests")
		requests, err := strconv.Atoi(updatedRequests)
		if err != nil {
			panic(err)
		}

		updatedCustomerRequests, err := pg.UpdateCustomerAvailableRequests(int64(customerID), int64(requests))
		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to Update AvailableRequests")
			panic(err)
		}

		c.JSON(http.StatusOK, updatedCustomerRequests)
	})

	return r
}

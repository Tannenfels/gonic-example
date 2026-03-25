package main

import (
	"9263-solution/Models"
	"9263-solution/Repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()
	router.GET("/booklist", getBookings)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	router.POST("/book")

	router.Run("localhost:8080")
}

func getBookings(c *gin.Context) {
	userIdStr, userIdExists := c.GetQuery("user_id")
	placeIdStr, placeIdExists := c.GetQuery("place_id")

	if userIdExists && placeIdExists {
		c.AbortWithStatus(400)
	}

	var bookings []Models.Booking
	var err error

	if userIdExists {
		userId, _ := strconv.ParseUint(userIdStr, 10, 64)
		bookings, err = Repositories.GetBookingsByUserId(userId)
		if err != nil {
			c.AbortWithStatus(500)
		}

	} else if placeIdExists {
		placeId, _ := strconv.ParseUint(placeIdStr, 10, 64)
		bookings, err = Repositories.GetBookingsByPlaceId(placeId)
		if err != nil {
			c.AbortWithStatus(500)
		}
	} else {
		c.AbortWithStatus(400)
	}

	c.IndentedJSON(200, bookings)
}

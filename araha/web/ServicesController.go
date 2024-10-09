package web

import (
	"araha/araha/models"
	"araha/araha/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSubscriptionController() gin.HandlerFunc {
	var createSubscription services.NewSubscriptionServices

	//if http.MethodPost {
	//
	//}

	return func(c *gin.Context) {
		var Subscription models.Subscription

		if err := c.ShouldBindJSON(&Subscription); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error creating subscription": err.Error()})
			return
		}

		result, err := createSubscription.CreateSubscription(Subscription)

		if result == 409 {
			c.JSON(http.StatusConflict, gin.H{"message": "Already on this subscription"})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": "Created Subscription successfully !!!",
				"Object": result})
		}

		if err != nil {
			fmt.Printf("Couldnt create a subscription %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	}

}

func UpdateSubscriptionController() gin.HandlerFunc {
	var updateSubscription services.NewSubscriptionServices
	return func(c *gin.Context) {
		var subscription models.Subscription
		if err := c.ShouldBindJSON(&subscription); err != nil {
			c.JSON(http.StatusNotModified, gin.H{"Error updating subscription": err})
		}
		result, err := updateSubscription.UpdateSubscription(subscription)
		if err != nil {
			c.JSON(http.StatusNotModified, gin.H{"Error": err})
		}
		c.JSON(http.StatusOK, gin.H{"data": "Updated Successfully",
			"Object": result})
	}
}

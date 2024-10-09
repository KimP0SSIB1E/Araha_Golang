package mapper

import (
	"araha/araha/constants"
	"araha/araha/exceptions"
	"araha/araha/models"
	"net/http"
	"time"
)

func FindSubscriptionTypes(subscription models.Subscription) models.Subscription {
	var noSubscription models.Subscription

	if subscription.SubscriptionType == constants.NETFLIX {
		var userBalance models.User
		var subscriptionCost float64

		subscriptionCost = 1000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = int(userBalance.Balance) - balance
		subscription.Date = time.Now().String()
		return subscription
	}

	if subscription.SubscriptionType == constants.JUMIA {
		var userBalance models.User
		var subscriptionCost float64

		subscriptionCost = 2000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = int(userBalance.Balance) - balance
		subscription.Date = time.Now().String()
		return subscription
	}

	if subscription.SubscriptionType == constants.GOTV {
		var userBalance models.User
		var subscriptionCost float64

		subscriptionCost = 4000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = int(userBalance.Balance) - balance
		subscription.Date = time.Now().String()
		return subscription
	}
	return noSubscription
}

func subscriptionNotAvailableErrorMessage() exceptions.MyException {
	err := exceptions.MyException{Data: http.StatusExpectationFailed, Object: "This service is currently not available\nPlease try again later"}
	return err
}

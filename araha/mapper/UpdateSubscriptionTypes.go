package mapper

import "araha/araha/models"

func UpdateSubscription(subscription models.Subscription) int {
	subscriptionFound := FindSubscriptionTypes(subscription)
	updatedSubscription := subscriptionFound.Amount + subscription.Amount
	return updatedSubscription
}

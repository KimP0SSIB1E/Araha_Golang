package models

import (
	"araha/araha/constants"
)

type Subscription struct {
	ID               string                     `json:"id"`
	SubscriptionType constants.SubscriptionType `json:"SubscriptionType"`
	Description      string                     `json:"description"`
	Amount           int
	Date             string
	UserId           int
}

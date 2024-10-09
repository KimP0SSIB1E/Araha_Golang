package services

import (
	"araha/araha/exceptions"
	"araha/araha/mapper"
	"araha/araha/models"
	"araha/araha/repository"
	"log"
	"net/http"
)

type CreateSubscriptionServices interface {
	CreateSubscription(subscription models.Subscription) (int, error)
}

type NewSubscriptionServices struct{}

func (nss *NewSubscriptionServices) CreateSubscription(subscription models.Subscription) (int, error) {
	var invalidDetailsException exceptions.MyException
	var foundSubType models.Subscription

	db, err := repository.SubscriptionRepo()
	db.Find(&foundSubType)

	if foundSubType.SubscriptionType == subscription.SubscriptionType {
		return http.StatusConflict, nil
	}
	//var userBalance models.User
	// userBalance.Balance > 0

	if subscription.SubscriptionType != " " {
		subscription = mapper.FindSubscriptionTypes(subscription)
		db, err := repository.SubscriptionRepo()
		db.Save(subscription)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusCreated, nil
	}

	if err != nil {
		log.Fatalf("An error occurred in creating a new subscription: %v ", err)
	}
	return http.StatusBadRequest, &invalidDetailsException
}

type UpdateSubscriptionServices struct{}

func (nss *NewSubscriptionServices) UpdateSubscription(subscription models.Subscription) (int, error) {
	//var cannotUpdateSubscriptionException exceptions.MyException

	if subscription.Amount != 0 || subscription.SubscriptionType != " " {
		foundSubscription := mapper.FindSubscriptionTypes(subscription)
		newSubscription := foundSubscription.Amount + subscription.Amount
		db, err := repository.SubscriptionRepo()
		db.Save(newSubscription)

		if err != nil {
			log.Fatalf("Couldn't update the user subscription %v", err)
			return http.StatusAccepted, nil
		}
	}
	return http.StatusNotModified, nil

}

func (nss *NewSubscriptionServices) DeleteSubscription(subscription models.Subscription) {
	var UnableToDeleteSubscriptionException exceptions.MyException

	db, err := repository.SubscriptionRepo()

	db.Delete(subscription)

	if err != nil {
		log.Fatalf("Couldn't delete subscription %v", UnableToDeleteSubscriptionException)
	}
}

func (nss *NewSubscriptionServices) GetAllSubscription() interface{} {
	var unableToGetAllValuesException exceptions.MyException
	var allSubscription models.Subscription

	db, err := repository.SubscriptionRepo()
	foundSub := db.Find(&allSubscription)

	if err != nil {
		log.Fatalf("Could'nt retrieve all the values from the database %v", unableToGetAllValuesException)
	}

	return foundSub

}

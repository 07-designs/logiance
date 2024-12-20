// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

import (
	"github.com/Shridhar2104/logilo/graphql/models"
)

type AccountInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Accounts struct {
	Orders []*models.Order `json:"orders"`
}

type Mutation struct {
}

type OrderInput struct {
	AccountID string                `json:"accountId"`
	LineItems []*OrderLineItemInput `json:"lineItems"`
}

type OrderLineItemInput struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type PaginationInput struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

type Query struct {
}

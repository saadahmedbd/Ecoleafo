package guestcarthandler

import guestcartservice "github.com/saadahmedbd/Treestore/Rest/Service/guestCartService"

type GuestcartHandler struct {
	service guestcartservice.Guestcartservice
}

func NewGuestCartHandler(service guestcartservice.Guestcartservice) *GuestcartHandler {
	return &GuestcartHandler{service: service}
}

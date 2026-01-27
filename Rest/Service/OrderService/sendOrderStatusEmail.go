package orderservice

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *orderService) sendOrderStatusEmail(order *models.Order) {
	fmt.Printf("[EMAIL] Attempting to send order status email for order %s to %s\n", order.OrderNumber, order.CustomerEmail)
	
	if order.CustomerEmail == "" {
		fmt.Printf("[EMAIL ERROR] No customer email found for order %s\n", order.OrderNumber)
		return
	}
	
	emailService := util.NewEmailService()
	if err := emailService.SendOrderConfirmedEmail(order.CustomerEmail, order.OrderNumber, order.Status); err != nil {
		fmt.Printf("[EMAIL ERROR] Failed to send order status email: %v\n", err)
	} else {
		fmt.Printf("[EMAIL SUCCESS] Order status email sent successfully to %s\n", order.CustomerEmail)
	}
}

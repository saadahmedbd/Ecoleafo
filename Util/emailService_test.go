package util

import (
	"fmt"
	"os"
	"testing"
)

// TestEmailService_SendOrderPlacedEmail tests the order placed email functionality
// To run this test:
// 1. Set RESEND_API_KEY and RESEND_FROM_EMAIL in your .env file
// 2. Run: go test -v ./Util -run TestEmailService_SendOrderPlacedEmail
func TestEmailService_SendOrderPlacedEmail(t *testing.T) {
	// Skip test if API key is not configured
	if os.Getenv("RESEND_API_KEY") == "" {
		t.Skip("RESEND_API_KEY not configured, skipping test")
	}

	emailService := NewEmailService()
	
	// Test data
	testEmail := "test@example.com" // Replace with your test email
	orderNumber := "ORD-1234567890"
	total := 1500.00

	err := emailService.SendOrderPlacedEmail(testEmail, orderNumber, total)
	
	if err != nil {
		t.Errorf("Failed to send order placed email: %v", err)
	} else {
		fmt.Printf("✓ Order placed email sent successfully to %s\n", testEmail)
	}
}

// TestEmailService_SendOrderConfirmedEmail tests the order confirmed email functionality
func TestEmailService_SendOrderConfirmedEmail(t *testing.T) {
	// Skip test if API key is not configured
	if os.Getenv("RESEND_API_KEY") == "" {
		t.Skip("RESEND_API_KEY not configured, skipping test")
	}

	emailService := NewEmailService()
	
	// Test data
	testEmail := "test@example.com" // Replace with your test email
	orderNumber := "ORD-1234567890"
	status := "processing"

	err := emailService.SendOrderConfirmedEmail(testEmail, orderNumber, status)
	
	if err != nil {
		t.Errorf("Failed to send order confirmed email: %v", err)
	} else {
		fmt.Printf("✓ Order confirmed email sent successfully to %s\n", testEmail)
	}
}

// Example: Manual test function (not a unit test)
// Run this with: go run Util/emailService_test.go
func ExampleManualEmailTest() {
	// Load environment variables
	// Make sure to set RESEND_API_KEY and RESEND_FROM_EMAIL in .env
	
	emailService := NewEmailService()
	
	// Test order placed email
	err := emailService.SendOrderPlacedEmail(
		"buyer@example.com",
		"ORD-1234567890",
		2500.00,
	)
	
	if err != nil {
		fmt.Printf("Error sending order placed email: %v\n", err)
	} else {
		fmt.Println("✓ Order placed email sent successfully")
	}
	
	// Test order confirmed email
	err = emailService.SendOrderConfirmedEmail(
		"buyer@example.com",
		"ORD-1234567890",
		"processing",
	)
	
	if err != nil {
		fmt.Printf("Error sending order confirmed email: %v\n", err)
	} else {
		fmt.Println("✓ Order confirmed email sent successfully")
	}
}

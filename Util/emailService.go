package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type EmailService struct {
	apiKey string
	from   string
}

type EmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Html    string   `json:"html"`
}

func NewEmailService() *EmailService {
	return &EmailService{
		apiKey: os.Getenv("RESEND_API_KEY"),
		from:   os.Getenv("RESEND_FROM_EMAIL"),
	}
}

func (e *EmailService) SendEmail(to, subject, htmlContent string) error {
	if e.apiKey == "" {
		return fmt.Errorf("RESEND_API_KEY not configured")
	}

	emailReq := EmailRequest{
		From:    e.from,
		To:      []string{to},
		Subject: subject,
		Html:    htmlContent,
	}

	jsonData, err := json.Marshal(emailReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+e.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email: status %d", resp.StatusCode)
	}

	return nil
}

func (e *EmailService) SendOrderPlacedEmail(to, orderNumber string, total float64) error {
	subject := "Order Placed Successfully - " + orderNumber
	// Use direct Cloudinary image URL (not collection URL)
	logoURL := "https://res.cloudinary.com/ddylnmsou/image/upload/v1769519546/AIRetouch_20251205_124752325_mwsd4b.png"

	html := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
			<div style="text-align: center; margin-bottom: 30px;">
				<img src="%s" alt="Ecoleafo Logo" style="max-width: 150px; height: auto;" />
			</div>
			<h2 style="color: #4CAF50; text-align: center;">Order Placed Successfully!</h2>
			<p>Thank you for your order. Your order has been received and is being processed.</p>
			<div style="background-color: #f5f5f5; padding: 20px; border-radius: 5px; margin: 20px 0;">
				<p><strong>Order Number:</strong> %s</p>
				<p><strong>Total Amount:</strong> ৳%.2f</p>
			</div>
			<p>You will receive another email once your order is confirmed by the seller.</p>
			<p style="color: #666; font-size: 12px; margin-top: 30px; text-align: center;">This is an automated email. Please do not reply.</p>
		</div>
	`, logoURL, orderNumber, total)

	return e.SendEmail(to, subject, html)
}

func (e *EmailService) SendOrderConfirmedEmail(to, orderNumber, status string) error {
	subject := "Order Confirmed - " + orderNumber
	// Use direct Cloudinary image URL (not collection URL)
	logoURL := "https://res.cloudinary.com/ddylnmsou/image/upload/v1769519546/AIRetouch_20251205_124752325_mwsd4b.png"

	html := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
			<div style="text-align: center; margin-bottom: 30px;">
				<img src="%s" alt="Ecoleafo Logo" style="max-width: 150px; height: auto;" />
			</div>
			<h2 style="color: #2196F3; text-align: center;">Order Confirmed!</h2>
			<p>Great news! Your order has been confirmed and is being prepared for shipment.</p>
			<div style="background-color: #f5f5f5; padding: 20px; border-radius: 5px; margin: 20px 0;">
				<p><strong>Order Number:</strong> %s</p>
				<p><strong>Status:</strong> %s</p>
			</div>
			<p>We'll notify you once your order has been shipped.</p>
			<p style="color: #666; font-size: 12px; margin-top: 30px; text-align: center;">This is an automated email. Please do not reply.</p>
		</div>
	`, logoURL, orderNumber, status)

	return e.SendEmail(to, subject, html)
}

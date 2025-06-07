package utils

import (
	"context"
	"fmt"

	"github.com/resend/resend-go/v2"
)

func sendNotificationEmail(toEmail, subject, htmlBody string, ResendAPIKey string, FromEmail string) error {
	if ResendAPIKey == "" {
		return fmt.Errorf("la variable de entorno RESEND_API_KEY no está configurada")
	}
	if FromEmail == "" {
		return fmt.Errorf("la variable de entorno FROM_EMAIL no está configurada")
	}

	client := resend.NewClient(ResendAPIKey)
	ctx := context.Background()

	params := &resend.SendEmailRequest{
		From:    FromEmail,
		To:      []string{toEmail},
		Subject: subject,
		Html:    htmlBody,
	}

	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		return fmt.Errorf("error al enviar email con Resend: %w", err)
	}

	fmt.Printf("Email enviado con éxito. ID: %s\n", sent.Id)
	return nil
}

func SendIncidenceUpdateEmail(studentEmail, status string, data TemplateData, ResendAPIKey string, FromEmail string) error {
	if status != "borrador" && status != "resuelto" && status != "rechazado" {
		fmt.Printf("No se envía correo para el estado: %s\n", status)
		return nil
	}

	subject, body, err := GetEmailContent(status, data)
	if err != nil {
		return fmt.Errorf("error al obtener contenido del email: %w", err)
	}

	return sendNotificationEmail(studentEmail, subject, body, ResendAPIKey, FromEmail)
}

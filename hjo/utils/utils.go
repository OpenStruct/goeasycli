package utils

import (
	"time"
    "bytes"
	"html/template"
	"math/rand"
    "hjo/config"
	"hjo/logger"
	"github.com/resend/resend-go/v2"
	"go.uber.org/zap"
)

// Randomly Generate credentials
func CredentialsGenerator(numberOfCharacters int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321@#$%^&*()")
	rand.NewSource(time.Now().UnixNano())
	str := make([]rune, numberOfCharacters)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}


func SendEmailWithResend(fromEmail, toEmail, subject string, emailTemplate string, templateData any) {
	t, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		logger.Error("failed to parse email template", zap.Any("error", err))
		return
	}
	var tpl bytes.Buffer

	if err := t.Execute(&tpl, templateData); err != nil {
		logger.Error("failed to execute template", zap.Any("error", err))
		return
	}

	htmlContent := tpl.String()

	client := resend.NewClient(config.CFG.V.GetString("RESEND_KEY"))
	params := &resend.SendEmailRequest{
		From:    fromEmail,
		To:      []string{toEmail},
		Html:    htmlContent,
		Subject: subject,
	}
	sent, err := client.Emails.Send(params)
	if err != nil {
		logger.Error("failed to send e-mail", zap.Any("error", err))
		return
	}

	logger.Info("messege sent", zap.Any("message_id", sent.Id))

}
package service

import (
	"bytes"
	"doodocs_task/internal/config"
	"github.com/jordan-wright/email"
	"mime/multipart"
	"net/smtp"
)

type EmailService struct {
	smtpUsername string
	smtpPassword string
}

func NewEmailService(cfg config.SMTPConfig) *EmailService {
	return &EmailService{
		smtpUsername: cfg.Username,
		smtpPassword: cfg.Password,
	}
}

func (s *EmailService) CheckFile(file *multipart.FileHeader) bool {
	t := file.Header.Get("Content-Type")

	AllowedTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/pdf": true,
	}

	return AllowedTypes[t]
}

func (s *EmailService) FileToBytes(file *multipart.FileHeader) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var fileBytes []byte
	if file.Size > 0 {
		fileBytes = make([]byte, file.Size)
		_, err := f.Read(fileBytes)
		if err != nil {
			return nil, err
		}
	}

	return fileBytes, nil
}

func (s *EmailService) SendFileToEmails(fileBytes []byte, filename string, mimeType string, emails []string) error {
	e := email.NewEmail()
	e.Subject = "Doodocs voz'mite menya"
	e.Text = []byte("Please find the attached file.")

	e.From = s.smtpUsername
	p := s.smtpPassword

	_, err := e.Attach(bytes.NewReader(fileBytes), filename, mimeType)
	if err != nil {
		return err
	}

	// Отправляем сообщение на каждый из указанных email
	for _, recipientEmail := range emails {
		// check email
		e.To = []string{recipientEmail}

		// Отправляем письмо
		err = e.Send("smtp.mail.ru:587", smtp.PlainAuth("", e.From, p, "smtp.mail.ru"))
		if err != nil {
			return err
		}
	}

	return nil
}

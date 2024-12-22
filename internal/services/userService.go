package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/google/uuid"
	"github.com/wignn/library-api/internal/model"
	"github.com/wignn/library-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func GetUserById(db *repository.DB, id int) (*model.GetUserResponse, error) {
	return repository.GetUserById(db, id)
}

func UpdateUserProfile(db *repository.DB, id int, username, profilePicture, email string) (*model.User, error) {
	user := &model.User{
		ID:             id,
		Username:       username,
		Email:          email,
		ProfilePicture: &profilePicture,
	}

	return repository.UpdateUserProfile(db, user)
}
func SendEmailVerification(db *repository.DB, id int) error {
	user, err := repository.GetUserById(db, id)
	if err != nil {
		log.Printf("Error getting user by id: %v\n", err)
		return err
	}

	token, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating UUID: %v\n", err)
		return err
	}

    err = repository.UpdateUserToken(db, user.ID, token.String())

    if err != nil {
        log.Printf("Error updating user token: %v\n", err)
        return err
    }
	verificationURL := fmt.Sprintf("http://localhost/auth/reset/%s", token.String())

	htmlMessage := fmt.Sprintf(`
    <html>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <style>
            body {
                font-family: Arial, Helvetica, sans-serif;
                background-color: #f9f9f9;
                color: #333;
                margin: 0;
                padding: 20px;
                align-items: center;
                justify-content: center;
                display: flex;
            }
            .email-container {
                max-width: 600px;
                margin: 0 auto;
                background: #ffffff;
                border: 1px solid #ddd;
                border-radius: 8px;
                padding: 20px;
                box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            }
            .email-header {
                font-size: 20px;
                font-weight: bold;
                margin-bottom: 20px;
            }
            .email-body {
                font-size: 16px;
                line-height: 1.6;
            }
            .verify-button {
                display: inline-block;
                padding: 12px 25px;
                font-size: 16px;
                color: #fff;
                background-color: #007bff;
                text-decoration: none;
                border-radius: 5px;
                margin-top: 20px;
            }
            .verify-button:hover {
                background-color: #0056b3;
            }
        </style>
    </head>
    <body>
        <div class="email-container">
            <p class="email-header">Hello %s,</p>
            <div class="email-body">
                <p>Please click the button below to verify your email address:</p>
                <a href="%s" class="verify-button">Verify Email</a>
            </div>
        </div>
    </body>
    </html>
`, user.Username, verificationURL)

	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")
	to := []string{user.Email}

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: Email Verification\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		os.Getenv("EMAIL_USERNAME"), user.Email, htmlMessage,
	))

	err = smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL_USERNAME"), to, msg)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return err
	}

	log.Printf("Verification email sent to %s\n", user.Email)
	return nil
}


func ResetPassword(db *repository.DB, id int, newPassword, token string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

 return repository.ResetPassword(db, id, string(hashedPassword), token)
}
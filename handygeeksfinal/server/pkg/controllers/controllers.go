package controllers

import (
	"log"
	"os"
	"server/pkg/models"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"

	"server/pkg/utils"
)

func HandleHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is running",
	})
}

// Used for contact us page
func HandleContactUs(c *fiber.Ctx) error {
	// NOTE: USING GOMAIL PACKAGE (Preferred)
	log.Println("Contact-Us endpoint hit")

	var contactUsFormData models.ContactUsData
	if err := c.BodyParser(&contactUsFormData); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
		})
	}

	log.Println(contactUsFormData.Name)

	// // boiler plate code
	message := mail.NewMessage()
	// email will be sent using the email provided by the user
	message.SetHeader("From", contactUsFormData.Email)
	message.SetHeader("To", os.Getenv("MAIL_RECEIVER_EMAIL"))
	message.SetHeader("Subject", "New Contact Made On Website")

	// body will be sent from the frontend
	body := "Name: " + contactUsFormData.Name + "\nEmail: " + contactUsFormData.Email + "\nPhone: " + contactUsFormData.Phone + "\nDetails: " + contactUsFormData.Details + "\nQuestions: " + contactUsFormData.Questions
	message.SetBody("text/plain", body)

	dialer := mail.NewDialer(os.Getenv("HOST"), 587, os.Getenv("MAIL_SENDER_EMAIL"), os.Getenv("MAIL_SENDER_PASS"))

	log.Println("Sending email...")
	if err := dialer.DialAndSend(message); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error sending email",
		})
	}
	log.Println("Email sent!")

	return nil
}

func HandleLogin(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Cannot parse JSON",
			},
		)
	}

	var existingUser models.User
	existingUser.Email = user.Email
	if err := existingUser.GetUserByEmail(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"status":  "error",
				"message": "User not found",
			},
		)
	}

	if ok := utils.CompareMyPass(existingUser.Password, user.Password); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Invalid password",
			},
		)
	} else {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status":  "ok",
				"message": "Logged in",
			},
		)
	}

}
func HandleRegister(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Cannot parse JSON",
			},
		)
	}

	if err := newUser.CreateUser(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status": "error",
				// Most likely user already exists
				"message": err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "User created",
	})
}

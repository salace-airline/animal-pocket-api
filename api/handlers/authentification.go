package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/salace-airline/animalpocketresources/database"
	"github.com/salace-airline/animalpocketresources/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	err = database.DB.Db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User has created",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User

	database.DB.Db.Where("email=?", data["email"]).First(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "not found",
			"message": "user not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "bad request",
			"message": "incorrect password",
		})
	}

	// Create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*24).Unix(), 0)), // 1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "server error",
			"message": "could not login",
		})
	}

	// Create cookie
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "user connected",
		"data":    user,
	})
}

func Logout(c *fiber.Ctx) error {
	// Remove cookie
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "user has logged out",
	})
}

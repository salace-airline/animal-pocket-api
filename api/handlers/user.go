package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/salace-airline/animalpocketresources/database"
	"github.com/salace-airline/animalpocketresources/models"

	"github.com/gofiber/fiber/v2"
)

func GetActualUser(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	var user models.User
	if err != nil && user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	database.DB.Db.Where("id=?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	var actualUser models.User

	if err != nil && actualUser.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "unauthenticated",
		})
	} else {
		var updateActualUser models.User
		errActual := c.BodyParser(&updateActualUser)
		if errActual != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"status":  "error",
				"message": "Something's wrong with your input",
				"data":    errActual,
			})
		}

		// Check if updated fields aren't over the maximum range of resources
		maxBugAndFish := int32(80)
		maxSeaCreature := int32(40)

		for _, bug := range updateActualUser.CaughtBug {
			if bug > maxBugAndFish {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Bug is out of range (maximum 80 items)",
					"data":    errActual,
				})
			}
		}

		for _, fish := range updateActualUser.CaughtFish {
			if fish > maxBugAndFish {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Fish is out of range (maximum 80 items)",
					"data":    errActual,
				})
			}
		}

		for _, seaCreature := range updateActualUser.CaughtSeaCreatures {
			if seaCreature > maxSeaCreature {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Sea creature is out of range (maximum 40 items)",
					"data":    errActual,
				})
			}
		}

		// Fields to update
		actualUser.Name = updateActualUser.Name
		actualUser.Email = updateActualUser.Email
		actualUser.CaughtFish = updateActualUser.CaughtFish
		actualUser.CaughtBug = updateActualUser.CaughtBug
		actualUser.CaughtSeaCreatures = updateActualUser.CaughtSeaCreatures

		claims := token.Claims.(*jwt.RegisteredClaims)
		database.DB.Db.Where("id=?", claims.Issuer).Updates(&actualUser)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "user updated",
		})
	}
}

func UpdateUserFish(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	var actualUser models.User

	if err != nil && actualUser.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "unauthenticated",
		})
	} else {
		claims := token.Claims.(*jwt.RegisteredClaims)
		database.DB.Db.Where("id=?", claims.Issuer).First(&actualUser)

		type UpdateUserCaughtFish struct {
			FishNumber int32 `json:"caught_fish"`
		}

		updateActualUser := new(UpdateUserCaughtFish)
		errActual := c.BodyParser(&updateActualUser)

		if errActual == nil && updateActualUser.FishNumber == 0 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"status":  "error",
				"message": "Something's wrong with your input",
				"data":    errActual,
			})
		} else {
			// Check if updated fields aren't over the maximum range of resources
			maxBugAndFish := int32(80)

			if updateActualUser.FishNumber > maxBugAndFish {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Fish is out of range (maximum 80 items)",
					"data":    errActual,
				})
			}

			// Fields to update
			actualUser.CaughtFish = append(actualUser.CaughtFish, updateActualUser.FishNumber)

			database.DB.Db.Where("id=?", claims.Issuer).Updates(&actualUser)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "success",
				"message": "user updated",
			})
		}
	}
}

func UpdateUserBug(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	var actualUser models.User

	if err != nil && actualUser.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "unauthenticated",
		})
	} else {
		claims := token.Claims.(*jwt.RegisteredClaims)
		database.DB.Db.Where("id=?", claims.Issuer).First(&actualUser)

		type UpdateUserCaughtBug struct {
			BugNumber int32 `json:"caught_bug"`
		}

		updateActualUser := new(UpdateUserCaughtBug)
		errActual := c.BodyParser(&updateActualUser)

		if errActual == nil && updateActualUser.BugNumber == 0 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"status":  "error",
				"message": "Something's wrong with your input",
				"data":    errActual,
			})
		} else {
			// Check if updated fields aren't over the maximum range of resources
			maxBugAndFish := int32(80)

			if updateActualUser.BugNumber > maxBugAndFish {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Bug is out of range (maximum 80 items)",
					"data":    errActual,
				})
			}

			// Fields to update
			actualUser.CaughtBug = append(actualUser.CaughtBug, updateActualUser.BugNumber)

			database.DB.Db.Where("id=?", claims.Issuer).Updates(&actualUser)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "success",
				"message": "user updated",
			})
		}

	}
}

func UpdateUserSeaCreature(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	var actualUser models.User

	if err != nil && actualUser.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "unauthenticated",
		})
	} else {
		claims := token.Claims.(*jwt.RegisteredClaims)
		database.DB.Db.Where("id=?", claims.Issuer).First(&actualUser)

		type UpdateUserCaughtSeaCreature struct {
			SeaCreatureNumber int32 `json:"caught_sea_creature"`
		}

		updateActualUser := new(UpdateUserCaughtSeaCreature)
		errActual := c.BodyParser(&updateActualUser)

		if errActual == nil && updateActualUser.SeaCreatureNumber == 0 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"status":  "error",
				"message": "Something's wrong with your input",
				"data":    errActual,
			})
		} else {
			// Check if updated fields aren't over the maximum range of resources
			maxSeaCreature := int32(40)

			if updateActualUser.SeaCreatureNumber > maxSeaCreature {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Sea creature is out of range (maximum 40 items)",
					"data":    errActual,
				})
			}

			// Fields to update
			actualUser.CaughtSeaCreatures = append(actualUser.CaughtSeaCreatures, updateActualUser.SeaCreatureNumber)

			database.DB.Db.Where("id=?", claims.Issuer).Updates(&actualUser)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "success",
				"message": "user updated",
			})
		}
	}
}

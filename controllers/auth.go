package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/form3tech-oss/jwt-go"
)

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HPR(c *fiber.Ctx) error {
	inp := c.FormValue("inp")
	out, _ := HashPass(inp)
	return c.JSON(fiber.Map{
		"out": out,
	})
}

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	user_gt := os.Getenv("USER")
	pass_gt := os.Getenv("PASSHASH")

	fmt.Println(pass_gt)
	fmt.Println(pass)
	fmt.Println(CheckPasswordHash(pass, pass_gt))

	if user != user_gt || !CheckPasswordHash(pass, pass_gt) {
		return c.Status(401).JSON(fiber.Map{
			"ok":    false,
			"error": "Unauthorized",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user_gt
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"token": t})
}

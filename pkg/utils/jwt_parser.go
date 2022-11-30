package utils

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenMetadata struct {
	Expires int64
	UserId  int8
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		userId := int8(claims["id"].(float64))
		return &TokenMetadata{
			Expires: expires,
			UserId:  userId,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

func VerifyRoute(c *fiber.Ctx) (*int8, error) {
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(c)
	if err != nil {
		return nil, err
	}

	expires := claims.Expires

	if now > expires {
		return nil, err
	}

	return &claims.UserId, nil
}

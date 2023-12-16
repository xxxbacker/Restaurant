package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"golangRestaurantManagement/utils"
	"net/http"
	"strings"
	"time"
)

const (
	authorizationHeader = "Authorization"
	signingKey          = "1234567890123456789012345678"
	tokenTTL            = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int32  `json:"user_id"`
	UserPost string `json:"user_post"`
}

func (server *Server) GenerateToken(ctx *gin.Context, phone, password string) (string, error) {
	arg := db.GetAccountForPasswordParams{
		Phone:    phone,
		Password: utils.GeneratePasswordHash(password),
	}

	user, err := server.store.GetAccountForPassword(ctx, arg)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.AccountID,
		user.Post,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Server) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return int(claims.UserId), claims.UserPost, nil
}

func (server *Server) accountIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, userPost, err := server.ParseToken(headerParts[1])
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(accountIdCtx, userId)
	c.Set(accountPostCtx, userPost)
	c.Next()
}

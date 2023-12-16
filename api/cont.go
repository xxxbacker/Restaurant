package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	accountIdCtx   = "userId"
	accountPostCtx = "userPost"
	chequeIdCtx    = "chequeId"
	menuItemIdCtx  = "menuItemId"
)

func getAccountId(c *gin.Context) (int, error) {
	id, ok := c.Get(accountIdCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getOrdId(c *gin.Context) (int32, error) {
	id, ok := c.Get(ordIdCtx)
	if !ok {
		return 0, errors.New("ord id not found")
	}

	idInt, ok := id.(int32)
	if !ok {
		return 0, errors.New("ord id is of invalid type")
	}

	return idInt, nil
}

func getAccountPostId(c *gin.Context) (string, error) {
	post, ok := c.Get(accountPostCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	postStr, ok := post.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return postStr, nil
}

func getChequeId(c *gin.Context) (int32, error) {
	id, ok := c.Get(chequeIdCtx)
	if !ok {
		return 0, errors.New("ord id not found")
	}

	idInt, ok := id.(int32)
	if !ok {
		return 0, errors.New("ord id is of invalid type")
	}

	return idInt, nil
}

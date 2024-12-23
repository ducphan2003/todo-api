package middleware

import (
	"todo-api/common"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if appErr, ok := err.(*common.AppError); ok {
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
			appErr := common.ErrInternal(err.(error))
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			panic(err)
		}
	}()
	c.Next()
}

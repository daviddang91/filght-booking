package middleware

import (
	"net/http"
	"strings"

	"github.com/daviddang91/filght-booking/common/dto"
	"github.com/daviddang91/filght-booking/common/util"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstring := ctx.GetHeader("Authorization")
		if tokenstring == "" {
			response := dto.BuildErrorResponse("Unauthorized", "Request does not contain an access token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		parts := strings.SplitN(tokenstring, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response := dto.BuildErrorResponse("Unauthorized", "Request header auth Incorrect format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		_, err := util.ParseToken(parts[1])
		if err != nil {
			response := dto.BuildErrorResponse("Unauthorized", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Next()
	}
}

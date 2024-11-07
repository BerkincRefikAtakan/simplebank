package api

import (
	"errors"
	"fmt"
	"net/http"
	"simplebank/token"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizantionHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizantionheader := ctx.GetHeader(authorizantionHeaderKey)
		if len(authorizantionheader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizantionheader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		authorizantiontype := strings.ToLower(fields[0])
		if authorizantiontype != authorizationTypeBearer {
			err := fmt.Errorf("unsuppoted authorization type %s", authorizantiontype)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accesToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accesToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}

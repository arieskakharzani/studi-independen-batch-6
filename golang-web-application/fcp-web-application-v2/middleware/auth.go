package middleware

import (
	"a21hc3NpZ25tZW50/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here

		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "error unauthorized user id"})
			} else {
				ctx.Redirect(http.StatusSeeOther, "login")
			}
			return
		}

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			return
		}

		ctx.Set("email", claims.Email)

		ctx.Next()
	})
}

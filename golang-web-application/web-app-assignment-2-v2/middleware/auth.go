package middleware

import (
	"a21hc3NpZ25tZW50/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"net/http"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Mengambil cookie dengan nama session_token
		sessionToken, err := ctx.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// Jika cookie session_token tidak ada, kembalikan respon HTTP 401 jika request memiliki header Content-Type "application/json",
				// atau redirect ke halaman login jika tidak.
				if ctx.GetHeader("Content-Type") != "application/json" {
					ctx.Redirect(http.StatusFound, "/login")
					ctx.Abort()
					return
				}
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				ctx.Abort()
				return
			}
			// Jika terjadi error lain saat mengambil cookie, kembalikan respon HTTP 400.
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Parsing JWT token pada cookie
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(sessionToken, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			// Jika parsing token gagal, kembalikan respon HTTP 401.
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		if !token.Valid {
			// Jika token tidak valid, kembalikan respon HTTP 401.
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		// Menyimpan nilai UserID dari claims ke dalam context dengan key "id".
		ctx.Set("id", claims.UserID)

		// Lanjutkan request ke handler atau endpoint selanjutnya.
		ctx.Next()
	})
}

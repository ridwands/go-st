package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang/domain"
	"golang/pkg/responder"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		if len(tokenSplit) > 2 || len(tokenSplit) < 2 {
			responder.Error(w, 401, "Get Error Authorization")
			return
		}
		tokenString = tokenSplit[1]
		ctx := r.Context()

		//JWT Validate
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("JWT_SECRET_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//Create Context For User Login
			data := domain.UserContextStruct{
				NISN: claims["nisn"],
				Name: claims["name"],
			}
			ctx = domain.UserContext(ctx, data)
		} else {
			responder.Error(w, 401, "Error Authorization")
		}

		//Send Context to Request
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

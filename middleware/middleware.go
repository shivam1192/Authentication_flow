package middleware

import (
	"net/http"
	"../handler/userhandler"
	"../model/usermodel"
	"github.com/dgrijalva/jwt-go"
	"context"
)


func CookieMiddleware(handler http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		cookie,err := r.Cookie("token")
		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
			tokenStr := cookie.Value
			claim := &usermodel.Claims{}
			tkn,err := jwt.ParseWithClaims(tokenStr,claim,
				func(t *jwt.Token) (interface{},error){
					return userhandler.Jwtkey,nil
			})
			if err != nil || !tkn.Valid{
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		c2 := context.WithValue(r.Context(), "claim", claim.UsersID)
		handler.ServeHTTP(w, r.WithContext(c2))
	}
}

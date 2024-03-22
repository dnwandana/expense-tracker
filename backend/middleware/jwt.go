package middleware

import (
	"context"
	"net/http"

	"github.com/dnwandana/expense-tracker/model/web"
	"github.com/dnwandana/expense-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAccessToken(next func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the Authorization header
		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			response := web.ResponseMessage{
				Status:  false,
				Message: "missing x-access-token header",
			}

			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJsonResponse(w, response)
			return
		}

		// validate and decode the token
		token, err := utils.VerifyToken(accessToken)
		if err != nil {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid token",
			}

			w.WriteHeader(http.StatusUnauthorized)
			utils.WriteJsonResponse(w, response)
			return
		}

		// get the claims and check if the token is valid
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid token",
			}

			w.WriteHeader(http.StatusUnauthorized)
			utils.WriteJsonResponse(w, response)
			return
		}

		// if everything is ok
		// extract user_id from the token
		// call the next handler, and pass the user_id to the context
		type contextKey string
		var UserIDKey contextKey = "user_id"
		ctx := context.WithValue(r.Context(), UserIDKey, claims["user_id"])
		next(w, r.WithContext(ctx))
	}
}

package middleware

import (
	"context"
	"net/http"

	"github.com/dnwandana/expense-tracker/model/web"
	"github.com/dnwandana/expense-tracker/utils"
	"github.com/golang-jwt/jwt/v5"
)

type UserIDKey string

const UserID UserIDKey = "user_id"

func RequireAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the Authorization header
		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			response := web.ResponseMessage{
				Status:  false,
				Message: "missing x-access-token header",
			}

			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}

		// validate and decode the token
		token, err := utils.VerifyToken(accessToken)
		if err != nil {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid token",
			}

			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}

		// get the claims and check if the token is valid
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid token",
			}

			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}

		// if everything is ok
		// extract user_id from the token
		// call the next handler, and pass the user_id to the context
		ctx := context.WithValue(r.Context(), UserID, claims["user_id"])
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

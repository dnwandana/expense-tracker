package routes

import (
	"database/sql"
	"net/http"

	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/model/web"
	"github.com/dnwandana/expense-tracker/model/web/request"
	"github.com/dnwandana/expense-tracker/repository"
	"github.com/dnwandana/expense-tracker/utils"
)

func SetupAuthenticationRoutes(mux *http.ServeMux, db *sql.DB) {
	// initialize repository
	userRepo := repository.NewUserRepository(db)

	// signin endpoint
	mux.HandleFunc("POST /api/auth/signin", func(w http.ResponseWriter, r *http.Request) {
		// parse the request
		req := new(request.AuthenticationRequest)
		utils.ReadJsonrequest(r, req)

		// find the user by username
		user := userRepo.FindByUsername(req.Username)

		if user.Username == "" {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid username or password",
			}

			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJsonResponse(w, response)
			return
		}

		// compare the password
		isPasswordMatch := utils.ComparePassword(user.Password, req.Password)
		if !isPasswordMatch {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid username or password",
			}

			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJsonResponse(w, response)
			return
		}

		// sign the jwt token
		accessToken := utils.SignAccessToken(user)
		refreshToken := utils.SignRefreshToken(user)

		// create token struct
		token := struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		// success response
		response := web.ResponseData{
			Status: true,
			Data:   token,
		}

		utils.WriteJsonResponse(w, response)
	})

	// signup endpoint
	mux.HandleFunc("POST /api/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		// parse the request
		req := new(request.AuthenticationRequest)
		utils.ReadJsonrequest(r, req)

		// check if the username is already taken
		data := userRepo.FindByUsername(req.Username)

		// error response
		if data.Username != "" {
			response := web.ResponseMessage{
				Status:  false,
				Message: "username already taken",
			}

			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJsonResponse(w, response)
			return
		}

		// hash the password
		hashedPassword := utils.HashPassword(req.Password)

		// create user struct
		user := &entity.User{
			Username: req.Username,
			Password: hashedPassword,
		}

		// store the user to the database
		userRepo.Create(user)

		// success response
		response := web.ResponseMessage{
			Status:  true,
			Message: "your account has been created",
		}

		w.WriteHeader(http.StatusCreated)
		utils.WriteJsonResponse(w, response)
	})
}

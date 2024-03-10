package routes

import (
	"database/sql"
	"encoding/json"
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
		err := json.NewDecoder(r.Body).Decode(req)
		utils.PanicIfError(err)

		// find the user by username
		user, err := userRepo.FindByUsername(req.Username)
		utils.PanicIfError(err)

		if user.Username == "" {
			response := web.Response{
				Status:  false,
				Message: "invalid username or password",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(response)
			utils.PanicIfError(err)
			return
		}

		// compare the password
		isPasswordMatch := utils.ComparePassword(user.Password, req.Password)
		if !isPasswordMatch {
			response := web.Response{
				Status:  false,
				Message: "invalid username or password",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(response)
			utils.PanicIfError(err)
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
		response := web.Response{
			Status: true,
			Data:   token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		utils.PanicIfError(err)
	})

	// signup endpoint
	mux.HandleFunc("POST /api/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		// parse the request
		req := new(request.AuthenticationRequest)
		err := json.NewDecoder(r.Body).Decode(req)
		utils.PanicIfError(err)

		// check if the username is already taken
		data, err := userRepo.FindByUsername(req.Username)
		utils.PanicIfError(err)

		// error response
		if data.Username != "" {
			response := web.Response{
				Status:  false,
				Message: "username already taken",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(response)
			utils.PanicIfError(err)
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
		err = userRepo.Create(user)
		utils.PanicIfError(err)

		// success response
		response := web.Response{
			Status:  true,
			Message: "your account has been created",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(response)
		utils.PanicIfError(err)
	})
}

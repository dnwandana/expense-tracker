package routes

import (
	"database/sql"
	"net/http"

	"github.com/dnwandana/expense-tracker/middleware"
	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/model/web"
	"github.com/dnwandana/expense-tracker/model/web/request"
	"github.com/dnwandana/expense-tracker/repository"
	"github.com/dnwandana/expense-tracker/utils"
)

func SetupCategoryRoutes(mux *http.ServeMux, db *sql.DB) {
	// initialize repository
	categoryRepo := repository.NewCategoryRepository(db)

	// create category
	mux.HandleFunc("POST /api/categories", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// parse the request
		req := new(request.CategoryRequest)
		utils.ReadJsonrequest(r, req)

		// create category
		category := entity.Category{
			UserID: userID,
			Name:   req.Name,
		}
		categoryRepo.Create(&category)

		// response
		response := web.Response{
			Status:  true,
			Message: "category created",
		}
		w.WriteHeader(http.StatusCreated)
		utils.WriteJsonResponse(w, response)
	})))

	// get category endpoint
	mux.HandleFunc("GET /api/categories", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// get categories
		categories := categoryRepo.FindByUserID(userID)

		// response
		response := web.Response{
			Status: true,
			Data:   categories,
		}
		utils.WriteJsonResponse(w, response)
	})))

	// detail category endpoint
	mux.HandleFunc("GET /api/categories/{category_id}", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request params
		categoryID := r.PathValue("category_id")

		// get category
		category := categoryRepo.FindByID(categoryID)

		// response
		response := web.Response{
			Status: true,
			Data:   category,
		}
		utils.WriteJsonResponse(w, response)
	})))

	// update category endpoint
	mux.HandleFunc("PUT /api/categories/{category_id}", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// request params
		categoryID := r.PathValue("category_id")

		// parse the request
		req := new(request.CategoryRequest)
		utils.ReadJsonrequest(r, req)

		// update category
		category := entity.Category{
			ID:   categoryID,
			Name: req.Name,
		}
		categoryRepo.Update(userID, &category)

		// response
		response := web.Response{
			Status:  true,
			Message: "category updated",
		}
		utils.WriteJsonResponse(w, response)
	})))

	// delete category endpoint
	mux.HandleFunc("DELETE /api/categories/{category_id}", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// request params
		categoryID := r.PathValue("category_id")

		// delete category
		categoryRepo.Delete(userID, categoryID)

		// response
		response := web.Response{
			Status:  true,
			Message: "category deleted",
		}
		utils.WriteJsonResponse(w, response)
	})))
}

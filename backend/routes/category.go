package routes

import (
	"database/sql"
	"net/http"
	"strconv"

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
	mux.Handle("POST /api/categories", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userIDFloat, ok := r.Context().Value(middleware.UserID).(float64)
		if !ok {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid user id",
			}
			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}
		userID := int(userIDFloat)

		// parse the request
		req := new(request.Category)
		utils.ReadJsonrequest(r, req)

		// create category
		category := entity.Category{
			UserID: userID,
			Name:   req.Name,
		}
		categoryRepo.Create(&category)

		// response
		response := web.ResponseMessage{
			Status:  true,
			Message: "category created",
		}
		utils.WriteJsonResponse(w, response, http.StatusCreated)
	})))

	// get category endpoint
	mux.Handle("GET /api/categories", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userIDFloat, ok := r.Context().Value(middleware.UserID).(float64)
		if !ok {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid user id",
			}
			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}
		userID := int(userIDFloat)

		// get categories
		categories := categoryRepo.FindByUserID(userID)

		// response
		response := web.ResponseData{
			Status: true,
			Data:   categories,
		}
		utils.WriteJsonResponse(w, response)
	})))

	// detail category endpoint
	mux.Handle("GET /api/categories/{category_id}", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userIDFloat, ok := r.Context().Value(middleware.UserID).(float64)
		if !ok {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid user id",
			}
			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}
		userID := int(userIDFloat)

		// request params
		categoryIDStr := r.PathValue("category_id")
		categoryID, err := strconv.Atoi(categoryIDStr)
		utils.PanicIfError(err)

		// get category
		category := categoryRepo.FindOne(userID, categoryID)
		if category.ID == 0 {
			response := web.ResponseMessage{
				Status:  false,
				Message: "no category found",
			}
			utils.WriteJsonResponse(w, response, http.StatusNotFound)
			return
		}

		// response
		response := web.ResponseData{
			Status: true,
			Data:   category,
		}
		utils.WriteJsonResponse(w, response)
	})))

	// update category endpoint
	mux.Handle("PUT /api/categories/{category_id}", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userIDFloat, ok := r.Context().Value(middleware.UserID).(float64)
		if !ok {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid user id",
			}
			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}
		userID := int(userIDFloat)

		// request params
		categoryIDStr := r.PathValue("category_id")
		categoryID, err := strconv.Atoi(categoryIDStr)
		utils.PanicIfError(err)

		// get category
		category := categoryRepo.FindOne(userID, categoryID)
		if category.ID == 0 {
			response := web.ResponseMessage{
				Status:  false,
				Message: "no category found",
			}
			utils.WriteJsonResponse(w, response, http.StatusNotFound)
			return
		}

		// parse the request
		req := new(request.Category)
		utils.ReadJsonrequest(r, req)

		// update category
		newCategory := entity.Category{
			ID:   categoryID,
			Name: req.Name,
		}
		categoryRepo.Update(userID, &newCategory)

		// response
		response := web.ResponseMessage{
			Status:  true,
			Message: "category updated",
		}
		utils.WriteJsonResponse(w, response)
	})))

	// delete category endpoint
	mux.Handle("DELETE /api/categories/{category_id}", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userIDFloat, ok := r.Context().Value(middleware.UserID).(float64)
		if !ok {
			response := web.ResponseMessage{
				Status:  false,
				Message: "invalid user id",
			}
			utils.WriteJsonResponse(w, response, http.StatusBadRequest)
			return
		}
		userID := int(userIDFloat)

		// request params
		categoryIDStr := r.PathValue("category_id")
		categoryID, err := strconv.Atoi(categoryIDStr)
		utils.PanicIfError(err)

		// get category
		category := categoryRepo.FindOne(userID, categoryID)
		if category.ID == 0 {
			response := web.ResponseMessage{
				Status:  false,
				Message: "no category found",
			}
			utils.WriteJsonResponse(w, response, http.StatusNotFound)
			return
		}

		// delete category
		categoryRepo.Delete(userID, categoryID)

		// response
		response := web.ResponseMessage{
			Status:  true,
			Message: "category deleted",
		}
		utils.WriteJsonResponse(w, response)
	})))
}

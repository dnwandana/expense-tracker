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

func SetupExpenseRoutes(mux *http.ServeMux, db *sql.DB) {
	// initialize repository
	expenseRepo := repository.NewExpenseRepository(db)

	// create expense
	mux.HandleFunc("POST /api/expenses", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(int)

		// parse the request
		req := new(request.ExpenseRequest)
		utils.ReadJsonrequest(r, req)

		// create expense
		expense := entity.Expense{
			UserID:      userID,
			CategoryID:  req.CategoryID,
			Amount:      req.Amount,
			Description: req.Description,
		}
		expenseRepo.Create(&expense)

		// response
		response := web.ResponseMessage{
			Status:  true,
			Message: "expense created",
		}
		utils.WriteJsonResponse(w, response, http.StatusCreated)
	})))

	// get expense endpoint
	mux.HandleFunc("GET /api/expenses", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(int)

		// get expenses
		expenses := expenseRepo.FindByUserID(userID)

		// response
		response := web.ResponseData{
			Status: true,
			Data:   expenses,
		}
		utils.WriteJsonResponse(w, response)
	})))

	// detail expense endpoint
	mux.HandleFunc("GET /api/expenses/:expense_id", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request params
		expenseIDStr := r.PathValue("id")
		expenseID, err := strconv.Atoi(expenseIDStr)
		utils.PanicIfError(err)

		// get expense
		expense := expenseRepo.FindByID(expenseID)

		// response
		response := web.ResponseData{
			Status: true,
			Data:   expense,
		}

		utils.WriteJsonResponse(w, response)
	})))

	// update expense endpoint
	mux.HandleFunc("PUT /api/expenses/:expense_id", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(int)

		// request params
		expenseIDStr := r.PathValue("id")
		expenseID, err := strconv.Atoi(expenseIDStr)
		utils.PanicIfError(err)

		// parse the request
		req := new(request.ExpenseRequest)
		utils.ReadJsonrequest(r, req)

		// update expense
		expense := entity.Expense{
			ID:          expenseID,
			UserID:      userID,
			CategoryID:  req.CategoryID,
			Amount:      req.Amount,
			Description: req.Description,
		}
		expenseRepo.Update(userID, &expense)

		// response
		response := web.ResponseMessage{
			Status:  true,
			Message: "expense updated",
		}
		utils.WriteJsonResponse(w, response)
	})))

	// delete expense endpoint
	mux.HandleFunc("DELETE /api/expenses/:expense_id", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(int)

		// request params
		expenseIDStr := r.PathValue("id")
		expenseID, err := strconv.Atoi(expenseIDStr)
		utils.PanicIfError(err)

		// delete expense
		expenseRepo.Delete(userID, expenseID)

		// response
		response := web.ResponseMessage{
			Status:  true,
			Message: "expense deleted",
		}
		utils.WriteJsonResponse(w, response)
	})))
}

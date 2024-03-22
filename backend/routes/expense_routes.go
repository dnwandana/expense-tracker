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

func SetupExpenseRoutes(mux *http.ServeMux, db *sql.DB) {
	// initialize repository
	expenseRepo := repository.NewExpenseRepository(db)

	// create expense
	mux.HandleFunc("POST /api/expenses", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

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
		response := web.Response{
			Status:  true,
			Message: "expense created",
		}
		w.WriteHeader(http.StatusCreated)
		utils.WriteJsonResponse(w, response)
	})))

	// get expense endpoint
	mux.HandleFunc("GET /api/expenses", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// get expenses
		expenses := expenseRepo.FindByUserID(userID)

		// response
		response := web.Response{
			Status: true,
			Data:   expenses,
		}
		utils.WriteJsonResponse(w, response)
	})))

	// detail expense endpoint
	mux.HandleFunc("GET /api/expenses/:expense_id", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get expense id
		expenseID := r.PathValue("id")

		// get expense
		expense := expenseRepo.FindByID(expenseID)

		// response
		response := web.Response{
			Status: true,
			Data:   expense,
		}

		utils.WriteJsonResponse(w, response)
	})))

	// update expense endpoint
	mux.HandleFunc("PUT /api/expenses/:expense_id", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// parameter: expense id
		expenseID := r.PathValue("expense_id")

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
		response := web.Response{
			Status:  true,
			Message: "expense updated",
		}
		utils.WriteJsonResponse(w, response)
	})))

	// delete expense endpoint
	mux.HandleFunc("DELETE /api/expenses/:expense_id", middleware.RequireAccessToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id from jwt
		userID := r.Context().Value("user_id").(string)

		// parameter: expense id
		expenseID := r.PathValue("expense_id")

		// delete expense
		expenseRepo.Delete(userID, expenseID)

		// response
		response := web.Response{
			Status:  true,
			Message: "expense deleted",
		}
		utils.WriteJsonResponse(w, response)
	})))
}

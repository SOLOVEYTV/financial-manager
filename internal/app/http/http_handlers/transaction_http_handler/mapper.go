package transaction_http_handler

import (
	"strings"
	"time"

	"github.com/SOLOVEYTV/financial-manager/internal/domain/model"
)

type transactionResponse struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Type        string    `json:"type"`
	Amount      int64     `json:"amount"`
	Description *string   `json:"description,omitempty"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

func mapDomainTransactionToResponse(tr *model.Transaction) transactionResponse {
	return transactionResponse{
		ID:          tr.ID,
		CreatedAt:   tr.CreatedAt,
		UpdatedAt:   tr.UpdatedAt,
		DeletedAt:   tr.DeletedAt,
		Type:        tr.Type.String(),
		Amount:      tr.Amount,
		Description: tr.Description,
		Category:    tr.Category.String(),
		Date:        tr.Date,
	}
}

func parseTransactionType(s string) model.TransactionType {
	switch strings.ToUpper(s) {
	case "INCOME":
		return model.TransactionTypeIncoming
	case "EXPENSE":
		return model.TransactionTypeExpense
	default:
		return model.TransactionTypeUnspecified
	}
}

func parseTransactionCategory(s string) model.TransactionCategory {
	switch strings.ToUpper(s) {
	case "SALARY":
		return model.TransactionCategorySalary
	case "FREELANCE":
		return model.TransactionCategoryFreelance
	case "BUSINESS":
		return model.TransactionCategoryBusiness
	case "INVESTMENT":
		return model.TransactionCategoryInvestment
	case "RENTAL_INCOME":
		return model.TransactionCategoryRentalIncome
	case "REFUND":
		return model.TransactionCategoryRefund
	case "GIFT_INCOME":
		return model.TransactionCategoryGiftIncome
	case "OTHER_INCOME":
		return model.TransactionCategoryOtherIncome
	case "FOOD":
		return model.TransactionCategoryFood
	case "RESTAURANTS":
		return model.TransactionCategoryRestaurants
	case "TRANSPORT":
		return model.TransactionCategoryTransport
	case "HOUSING":
		return model.TransactionCategoryHousing
	case "UTILITIES":
		return model.TransactionCategoryUtilities
	case "HEALTHCARE":
		return model.TransactionCategoryHealthcare
	case "SHOPPING":
		return model.TransactionCategoryShopping
	case "ENTERTAINMENT":
		return model.TransactionCategoryEntertainment
	case "EDUCATION":
		return model.TransactionCategoryEducation
	case "SUBSCRIPTIONS":
		return model.TransactionCategorySubscriptions
	case "TRAVEL":
		return model.TransactionCategoryTravel
	case "PERSONAL_CARE":
		return model.TransactionCategoryPersonalCare
	case "INSURANCE":
		return model.TransactionCategoryInsurance
	case "TAXES":
		return model.TransactionCategoryTaxes
	case "GIFTS":
		return model.TransactionCategoryGifts
	case "PETS":
		return model.TransactionCategoryPets
	case "OTHER_EXPENSE":
		return model.TransactionCategoryOtherExpense
	default:
		return model.TransactionCategoryUnspecified
	}
}

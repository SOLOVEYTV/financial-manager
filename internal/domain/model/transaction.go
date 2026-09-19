package model

import "time"

type (
	TransactionType     int16
	TransactionCategory int16
)

const (
	TransactionTypeUnspecified TransactionType = 0
	TransactionTypeIncoming    TransactionType = 1
	TransactionTypeExpense     TransactionType = 2

	TransactionCategoryUnspecified TransactionCategory = 0

	// Income categories (TransactionTypeIncoming)

	TransactionCategorySalary       TransactionCategory = 1    // зарплата
	TransactionCategoryFreelance    TransactionCategory = 2    // фриланс, подработка
	TransactionCategoryBusiness     TransactionCategory = 3    // доход от бизнеса / ИП
	TransactionCategoryInvestment   TransactionCategory = 4    // дивиденды, проценты, продажа активов
	TransactionCategoryRentalIncome TransactionCategory = 5    // аренда (если сдаёте жильё)
	TransactionCategoryRefund       TransactionCategory = 6    // возвраты, cashback
	TransactionCategoryGiftIncome   TransactionCategory = 7    // полученные подарки / переводы
	TransactionCategoryOtherIncome  TransactionCategory = 1000 // прочий доход

	// Expense categories (TransactionTypeExpense)

	TransactionCategoryFood          TransactionCategory = 1001 // продукты
	TransactionCategoryRestaurants   TransactionCategory = 1002 // кафе, рестораны, доставка
	TransactionCategoryTransport     TransactionCategory = 1003 // метро, такси, бензин, авто
	TransactionCategoryHousing       TransactionCategory = 1004 // аренда, ипотека, ремонт
	TransactionCategoryUtilities     TransactionCategory = 1005 // ЖКХ, интернет, мобильная связь
	TransactionCategoryHealthcare    TransactionCategory = 1006 // врачи, аптека, анализы
	TransactionCategoryShopping      TransactionCategory = 1007 // одежда, техника, маркетплейсы
	TransactionCategoryEntertainment TransactionCategory = 1008 // кино, игры, хобби
	TransactionCategoryEducation     TransactionCategory = 1009 // курсы, книги, обучение
	TransactionCategorySubscriptions TransactionCategory = 1010 // Netflix, Spotify, SaaS
	TransactionCategoryTravel        TransactionCategory = 1011 // билеты, отели, поездки
	TransactionCategoryPersonalCare  TransactionCategory = 1012 // спортзал, салоны, косметика
	TransactionCategoryInsurance     TransactionCategory = 1013 // страховки
	TransactionCategoryTaxes         TransactionCategory = 1014 // налоги, штрафы
	TransactionCategoryGifts         TransactionCategory = 1015 // подарки другим
	TransactionCategoryPets          TransactionCategory = 1016 // корм, ветеринар
	TransactionCategoryOtherExpense  TransactionCategory = 3000 // прочие расходы
)

type Transaction struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Type        TransactionType
	Amount      int64
	Description *string
	Category    TransactionCategory
	Date        time.Time
}

func NewTransaction(
	id string,
	now time.Time,
	trType TransactionType,
	amount int64,
	description *string,
	category TransactionCategory,
	date time.Time,
) *Transaction {
	return &Transaction{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Type:        trType,
		Amount:      amount,
		Description: description,
		Category:    category,
		Date:        date,
	}
}

func (t TransactionType) String() string {
	switch t {
	case TransactionTypeIncoming:
		return "INCOME"
	case TransactionTypeExpense:
		return "EXPENSE"
	default:
		return "UNSPECIFIED"
	}
}

func (c TransactionCategory) String() string {
	switch c {
	case TransactionCategorySalary:
		return "SALARY"
	case TransactionCategoryFreelance:
		return "FREELANCE"
	case TransactionCategoryBusiness:
		return "BUSINESS"
	case TransactionCategoryInvestment:
		return "INVESTMENT"
	case TransactionCategoryRentalIncome:
		return "RENTAL_INCOME"
	case TransactionCategoryRefund:
		return "REFUND"
	case TransactionCategoryGiftIncome:
		return "GIFT_INCOME"
	case TransactionCategoryOtherIncome:
		return "OTHER_INCOME"
	case TransactionCategoryFood:
		return "FOOD"
	case TransactionCategoryRestaurants:
		return "RESTAURANTS"
	case TransactionCategoryTransport:
		return "TRANSPORT"
	case TransactionCategoryHousing:
		return "HOUSING"
	case TransactionCategoryUtilities:
		return "UTILITIES"
	case TransactionCategoryHealthcare:
		return "HEALTHCARE"
	case TransactionCategoryShopping:
		return "SHOPPING"
	case TransactionCategoryEntertainment:
		return "ENTERTAINMENT"
	case TransactionCategoryEducation:
		return "EDUCATION"
	case TransactionCategorySubscriptions:
		return "SUBSCRIPTIONS"
	case TransactionCategoryTravel:
		return "TRAVEL"
	case TransactionCategoryPersonalCare:
		return "PERSONAL_CARE"
	case TransactionCategoryInsurance:
		return "INSURANCE"
	case TransactionCategoryTaxes:
		return "TAXES"
	case TransactionCategoryGifts:
		return "GIFTS"
	case TransactionCategoryPets:
		return "PETS"
	case TransactionCategoryOtherExpense:
		return "OTHER_EXPENSE"
	default:
		return "UNSPECIFIED"
	}
}

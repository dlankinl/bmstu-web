package v2

import (
	"github.com/google/uuid"
	"ppo/web"
)

type EntrepreneursResponse struct {
	Pages         int        `json:"numPages"`
	Entrepreneurs []web.User `json:"entrepreneurs"`
}

type EntrepreneurResponse struct {
	Entrepreneur web.User `json:"entrepreneur"`
}

type ContactResponse struct {
	Contact web.Contact `json:"contact"`
}

type ContactsResponse struct {
	EntrepreneurId uuid.UUID     `json:"entrepreneurId"`
	Contacts       []web.Contact `json:"contacts"`
}

type ActFieldResponse struct {
	ActField web.ActivityField `json:"activityField"`
}

type ActFieldsResponse struct {
	Pages     int                 `json:"numPages"`
	ActFields []web.ActivityField `json:"activityFields"`
}

type CompanyResponse struct {
	Company web.Company `json:"company"`
}

type CompaniesResponse struct {
	Pages          int           `json:"numPages"`
	EntrepreneurId uuid.UUID     `json:"entrepreneurId"`
	Companies      []web.Company `json:"companies"`
}

type FinReportResponse struct {
	FinReport web.FinancialReport `json:"financialReport"`
}

type FinReportByPeriodResponse struct {
	CompanyId uuid.UUID             `json:"companyId"`
	Period    web.Period            `json:"period"`
	Revenue   float32               `json:"revenue"`
	Costs     float32               `json:"costs"`
	Profit    float32               `json:"profit"`
	Reports   []web.FinancialReport `json:"reports"`
}

type RatingResponse struct {
	Rating float32 `json:"rating"`
}

type EntrepreneurReportResponse struct {
	Revenue float32 `json:"revenue"`
	Costs   float32 `json:"costs"`
	Profit  float32 `json:"profit"`
	Taxes   float32 `json:"taxes"`
	TaxLoad float32 `json:"taxLoad"`
}
